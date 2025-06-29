OPENVPN_DIR=openvpn/config
EASYRSA_DIR=$(OPENVPN_DIR)/easy-rsa

.PHONY: openvpn-init-certs openvpn-server

openvpn-init-certs:
	git clone https://github.com/OpenVPN/easy-rsa.git $(EASYRSA_DIR) || true
	cd $(EASYRSA_DIR) && cp vars.example vars && ./easyrsa init-pki
	cd $(EASYRSA_DIR) && ./easyrsa build-ca nopass
	cd $(EASYRSA_DIR) && ./easyrsa gen-dh
	cd $(EASYRSA_DIR) && ./easyrsa build-server-full server nopass
	cd $(EASYRSA_DIR) && ./easyrsa build-client-full client nopass
	cp $(EASYRSA_DIR)/pki/ca.crt $(OPENVPN_DIR)/
	cp $(EASYRSA_DIR)/pki/dh.pem $(OPENVPN_DIR)/
	cp $(EASYRSA_DIR)/pki/issued/server.crt $(OPENVPN_DIR)/
	cp $(EASYRSA_DIR)/pki/private/server.key $(OPENVPN_DIR)/
	cp $(EASYRSA_DIR)/pki/issued/client.crt $(OPENVPN_DIR)/
	cp $(EASYRSA_DIR)/pki/private/client.key $(OPENVPN_DIR)/

openvpn-server:
	sudo openvpn --config $(OPENVPN_DIR)/server.conf 