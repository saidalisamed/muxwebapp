# MuxWebApp
## Deployment Notes
1. Build for linux (Ubuntu) architecture amd64:
```shell
cd /path/to/muxwebapp/
GOOS=linux GOARCH=amd64 go build -o muxwebapp -v
```
2. Copy your project excluding the Go source `*.go` to your server in `/home/app.muxwebapp/`.
```shell
rsync -azv -e 'ssh -i ~.ssh/yourkey.pem' --exclude '*.go' --exclude 'debug' --exclude '__debug_bin' --exclude '.gitignore' --exclude '.git' --exclude 'prod.yml' ../muxwebapp user@server.example.com:/home/app.muxwebapp/
```
3. Configure supervisor to run the app.
```shell
sudo vim /etc/supervisor/conf.d/muxwebapp.conf

# add the following
[program:muxwebapp]
environment =
   ENV=prod,
command=/home/app.muxwebapp/muxwebapp
autostart=true
autorestart=true
startretries=10
user=app.muxwebapp
directory=/home/app.muxwebapp/muxwebapp/
redirect_stderr=true
stdout_logfile=/var/log/supervisor/muxwebapp.log
stdout_logfile_maxbytes=50MB
stdout_logfile_backups=10

# reload supervisor
sudo supervisorctl reload
```
4. Configure nginx as the frontend.
```shell
sudo vim /etc/nginx/sites-available/muxwebapp.conf

# add the following
upstream muxwebapp {
   server unix:/tmp/muxwebapp.sock;
}

server {
   listen 80;
   server_name muxwebapp.example.com;
   access_log /var/log/nginx/muxwebapp-access.log;
   error_log /var/log/nginx/muxwebapp-error.log error;
   location /static/ { alias /home/app.muxwebapp/muxwebapp/res/assets/; }
   location / {
      proxy_pass http://muxwebapp;
   }
}

# enable the website configuration
sudo ln -s /etc/nginx/sites-available/muxwebapp.conf /etc/nginx/sites-enabled
sudo service nginx reload
```
