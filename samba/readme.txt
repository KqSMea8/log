vi /etc/samba/smb.conf

[global]
        workgroup = WORKGROUP
        server string = Ted Samba Server %v
        netbios name = TedSamba
        security = user
        map to guest = Bad User
        passdb backend = tdbsam

[FileShare]
        comment = share some files
        path = /home
        public = yes
        writeable = yes
        create mask = 0644
        directory mask = 0755

[WebDev]
        comment = project development directory
        path = /home
        browseable = yes
        writable = yes
        valid users = ted
        write list = ted
        printable = no
        create mask = 0644
        directory mask = 0755



service smb restart
service nmb restart
chown -R nobody:nobody  /home
