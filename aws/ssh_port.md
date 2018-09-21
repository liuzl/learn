# To Change the SSH Port for Your Linux Server
1. Connect to your server via SSH (more info).
2. Switch to the root user (more info).
3. Run the following command:`vi /etc/ssh/sshd_config`.
4. Locate the following line: `# Port 22`
5. Remove # and change 22 to your desired port number.
6. Restart the sshd service
