USE sflowdb;

CREATE TABLE devices (
	    device_id INT AUTO_INCREMENT PRIMARY KEY,
        ip_address VARCHAR(15) NOT NULL,
        hostname VARCHAR(255) NULL,
        snmp_location VARCHAR(255) NULL,
        snmp_contact VARCHAR(255) NULL,
        snmp_version VARCHAR(10) NULL,
        snmp_community VARCHAR(255) NULL,
        snmp_port INT NULL,
);

CREATE TABLE sflow_interval(
        data_insercao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        device_id INT NOT NULL,
        sequence_number INT NOT NULL,
        source_id_index INT NOT NULL,
        if_operational_status BOOLEAN NOT NULL,
        if_admin_status BOOLEAN NOT NULL,
        if_speed INT NOT NULL,
        if_in_octets INT NOT NULL,
        if_in_packets INT NOT NULL,
        if_out_octets INT NOT NULL,
        if_out_packets INT NOT NULL,
)

CREATE TABLE sflow_probe (
        data_insercao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        probe_id INT AUTO_INCREMENT PRIMARY KEY,
        device_id INT NOT NULL,
        sequence_number INT NOT NULL,
        source_id_index INT NOT NULL,
        dropped_packets INT NOT NULL,
        vlan_src INT NOT NULL,
        vlan_dst INT NOT NULL,
        sample_rate INT NOT NULL,
        sample_pool INT NOT NULL,
)

CREATE TABLE probed_l2_header (
        probe_id INT NOT NULL,
        frame_length INT NOT NULL,
        frame_type INT NOT NULL,
        mac_src VARCHAR(17) NOT NULL,
        mac_dst VARCHAR(17) NOT NULL,
        eth_type INT NOT NULL,
)

CREATE TABLE ether_types (
        protocol_id INT PRIMARY KEY,
        protocol_name VARCHAR(255) NOT NULL,
)

CREATE TABLE probed_l3_header (
        probe_id INT NOT NULL,
        ip_src VARCHAR(15) NOT NULL,
        ip_dst VARCHAR(15) NOT NULL,
        l4_protocol INT NOT NULL,
)

CREATE TABLE l4_protocols (
        protocol_id INT PRIMARY KEY,
        protocol_name VARCHAR(255) NOT NULL,
)

CREATE TABLE probed_l4_header (
        probe_id INT NOT NULL,
        src_port INT NOT NULL,
        dst_port INT NOT NULL,
)

CREATE TABLE l5_protocols (
        port_id INT PRIMARY KEY,
        default_protocol VARCHAR(255) NOT NULL,
)


CREATE USER 'sflowadmin'@'%' IDENTIFIED BY 'SflowAdmin123';
GRANT ALL PRIVILEGES ON syslogdb.* TO 'sflowadmin'@'%';
FLUSH PRIVILEGES;