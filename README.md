# gsteps-go


- 查看服务的运行状态
  ```bash
  sudo supervisorctl status
  ```
- 启动服务
  ```bash
  sudo supervisorctl start $[program-name]
  ```
- 重启服务
  ```bash
  sudo supervisorctl signal usr2 $[program-name]
  ```
- 停止服务
  ```bash
  sudo supervisorctl stop $[program-name]
  ```
