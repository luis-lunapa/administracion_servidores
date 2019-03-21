#!/bin/bash
iptables -F
iptables -X
iptables -A INPUT -j DROP # Se bloquea todo filter
iptables -I INPUT -i lo -j ACCEPT # Aceptar todo lo

iptables -A INPUT -m state --state INVALID -j DROP # Tirar invalido

iptables -A INPUT -p tcp --dport 22 -j ACCEPT # Acepta SSH
iptables -A INPUT -p tcp --dport 80 -j ACCEPT # Acepta http
iptables -A INPUT -p tcp --dport 443 -j ACCEPT # Acepta https
iptables -A INPUT -p tcp --dport 53 -j ACCEPT # Acepta DNS


iptables -A OUTPUT -m state --state INVALID -j DROP # Tirar invalido
iptables -A OUTPUT -p tcp --dport 22 -j ACCEPT # Acepta salidas ssh
iptables -A OUTPUT -p tcp --dport 53 -j ACCEPT # Acepta respuestas dns
iptables -A OUTPUT -p tcp --dport 80 -j ACCEPT # Acepta salidas http
iptables -A OUTPUT -p tcp --dport 443 -j ACCEPT # Acepta salidas https

# Configurar nat
iptables -t nat -A POSTROUTING -o enp0s8 -j MASQUERADE
iptables -A FORWARD -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
iptables -t nat -A FORWARD  -o enp0s8 -j ACCEPT