table deleted
Table created successfully!
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /records                  --> assignment/handlers.ImportData (3 handlers)
[GIN-debug] GET    /records                  --> assignment/handlers.GetAllRecords (3 handlers)
[GIN-debug] PUT    /records                  --> assignment/handlers.UpdateRecord (3 handlers)
[GIN-debug] DELETE /records/:id              --> assignment/handlers.DeleteRecord (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2024/10/30 - 14:39:15 | 200 |    6.066987ms |             ::1 | PUT      "/records"
HIi: <nil>
Key: record:1
Error: <nil>
[GIN] 2024/10/30 - 14:39:20 | 200 |    2.910118ms |             ::1 | DELETE   "/records/1"
