```sh
git clone https://github.com/awslabs/amazon-kinesis-agent.git

cd amazon-kinesis-agent/

sudo ./setup --install

sudo vim /etc/aws-kinesis/agent.json
sudo vim /etc/init.d/aws-kinesis-agent

sudo service aws-kinesis-agent start
sudo systemctl daemon-reload
sudo ls /var/log/aws-kinesis-agent/
sudo service aws-kinesis-agent status
```
