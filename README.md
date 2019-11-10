# Golang Mux Webapp Project Skeleton
## Deployment Notes
1. Build for linux (Ubuntu) architecture amd64:
```shell
cd /path/to/<PROJECT>/
GOOS=linux GOARCH=amd64 go build -o <PROJECT> -v
```
2. Copy your project excluding the Go source `*.go` to your server in `/var/www/html/<PROJECT>`.
```shell
rsync -azv -e 'ssh -i /path/to/.pem' --exclude '*.go' --exclude 'debug' --exclude '.gitignore' --exclude '.git' --exclude '.vscode' ../<PROJECT> user@example.com:/var/www/html/
```
3. Configure supervisor to run the app.
```shell
sudo vim /etc/supervisor/conf.d/<PROJECT>.conf

# add the following
[program:<PROJECT>]
environment =
   ENV=prod,
command=/var/www/html/<PROJECT>/<PROJECT>
autostart=true
autorestart=true
startretries=10
user=www-data
directory=/var/www/html/<PROJECT>/
redirect_stderr=true
stdout_logfile=/var/log/supervisor/<PROJECT>.log
stdout_logfile_maxbytes=50MB
stdout_logfile_backups=10


# reload supervisor
sudo supervisorctl reload
```
4. Configure nginx as the frontend.
```shell
sudo vim /etc/nginx/sites-available/<PROJECT>.conf

# add the following
upstream <PROJECT> {
   server unix:/tmp/<PROJECT>.sock;
}

server {
   listen 80;
   server_name <PROJECT>;
   access_log /var/log/nginx/<PROJECT>-access.log;
   error_log /var/log/nginx/<PROJECT>-error.log error;
   location /static/ { alias /var/www/html/<PROJECT>/res/assets/; }
   location / {
      proxy_pass http://<PROJECT>;
   }
}


# enable the website configuration
sudo ln -s /etc/nginx/sites-available/<PROJECT>.conf /etc/nginx/sites-enabled
sudo service nginx reload
```
