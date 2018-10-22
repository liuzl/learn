* [SSIS – Copy Amazon S3 files from AWS to Azure](https://zappysys.com/blog/ssis-copy-move-amazon-s3-files-from-aws-to-azure/)
* https://rclone.org/
  * config aws s3
  * config azure storeage blobs
  * `rclone copy aws:bucket azure:account`
  * 常用命令
    * `rclone config`
    * `rclone lsd aws:bucket`
  * https://github.com/ncw/rclone/issues/1718
    * `--ignore-size` option
