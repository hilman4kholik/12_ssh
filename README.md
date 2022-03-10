SIMPLE SSH AUTO LOGIN & SEND SHELLSCRIPT

file config.ini adalah untuk parameternya

```
[settings]
address = 10.8.0.17 ;ip client yang akan di ssh
user = root ;Username client yang akan di ssh
password = 123456 ;password client yang akan di ssh

[commands]
;Contoh multiple command
;shellscript = cd /opt/applications, ls -sh, killall app, killall app1, reboot
shellScript = df -h
```

pastikan file config.ini berada satu folder dengan aplikasinya.

cukup lakukan chmod +x auto_ssh, lalu jalankan dengan cara:
```
./auto_ssh
```

ini bisa berfungsi untuk semua distro linux.

====================================================================
Jangan lupa berikan star dan sematkan sumber jika ini membantu anda.