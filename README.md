## 專案結構
/pkg: 共用模組  
/global: 轉案內全域變數  
/configs: 設定.yaml  
/deployment: sql 文件  
/internal: 內部主要程序  
/internal/infra: 專案初始化所需的dependencies  
/internal/router(/internal/handler): api 路由 (handler 層)  
/internal/service: 邏輯層  
/internal/model: table schema 層  
/internal/dao: 資料存取層  
/internal/crawler: 調用 BSC 服務層 和 分配任務組件  
/internal/subscriber: 訂閱組件  


## Indexer-Service (Dependencies: postgres, redis)
1. 透過 make postgres, make redis, make create_db, migrate_up 指令，建立地端環境
2. 透過 make run_server 啟動 golang 程序
3. 子程序:
   1. RunEthereumHeaderSyncing: 每30秒同步最新區塊號，並依序將尚未同步過的區塊號寫入 redis (最多1000筆)  
   為何是30秒，主要是想模擬鏈上原提供websocket服務 (ethclient.SubscribeNewHead方法)，可能是因為測試鏈關係，所以沒有提供此服務  
   2. RunEthereumSyncing: 每30秒從 redis 裡取出 250 筆區塊號，並透過 goroutine 方式爬回對應的區塊、交易和 event log 資料  
   若過程中 error 發生，會執行 db rollback 程序，並將這批區塊號寫回 redis，待下次再執行
   3. RunEthereumBlockStable: 每60秒去更新區塊號狀態，是否為穩定(is_stable = true when 當前區塊號小於最新區塊號減20)  

## API-Service (Dependencies: postgres)
1. 透過 make run_server 指令啟動 HTTP 服務，並依指定 API 路由取得對應資料  
2. api-service 取得資料不再是直接打 BSC RPC 服務，可以避免被外部使用者請求過多

## 可拓展方向
1. 透過 kubectl 工具實現 master-worker pattern  
把 RunEthereumHeaderSyncing 和 RunEthereumBlockStable 功能放在 master-service ，而 RunEthereumSyncing 功能放在 worker-service 裡  
並提供接口讓後台人員打入 master-service ，進而動態地擴容/縮容 worker-service  
2. 將 postgres 拆分成 cmd-database 和 read-database  
   透過 Nats 或是 kafka 等工具實現 CQRS 概念



## References
1. eth: https://ethereum.org/en/developers/docs/transactions/
2. json-rpc: https://eth.wiki/json-rpc/API
3. geth: https://goethereumbook.org/zh/transaction-query/ 
   
## Some Tips:
1. BSC rpc rate limit 10k/5min