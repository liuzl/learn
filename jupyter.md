* [远程jupyter notebook](https://www.blopig.com/blog/2018/03/running-jupyter-notebook-on-a-remote-server-via-ssh/)
  * `jupyter lab --port=9000 --no-browser`
  * `ssh -N -f -L 8888:localhost:9000 user@remoteserver`
