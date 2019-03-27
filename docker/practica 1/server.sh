apt-get update
apt-get upgrade
apt-get install -y php libapache2-mod-php
/usr/sbin/apache2ctl -D FOREGROUND
