---
page_title: "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san"
subcategory: ""
description: "Configuration of storage backend for NetApp ONTAP SAN."
xcsh_docs: {"aliases": ["backend servers", "custom storage config storage device list storage devices netapp trident netapp backend ontap san", "origin servers", "upstream servers"], "body_bytes": 25067, "body_sha256": "sha256:4480a66bed493f5c787481369f49dadcd6f63b939fde532248e1d002380a304d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:client_private_key", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:no_chap", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:password", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident", "path": "documentation/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-1203012223231023-2200220231023303-1122331312230001-0301322022031201-3022110023213232-3320210312013231-2131323210013121-0000231310220010", "registry_path": "docs/guides/resources--voltstack_site--reference--group-007.md", "relationships": [{"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--data_lif_dns_name", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san:ConflictingObjectAttributes:data_lif_dns_name,data_lif_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--data_lif_ip", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san:ConflictingObjectAttributes:data_lif_dns_name,data_lif_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--management_lif_dns_name", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san:ConflictingObjectAttributes:management_lif_dns_name,management_lif_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--management_lif_ip", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san:ConflictingObjectAttributes:management_lif_dns_name,management_lif_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san:ConflictingObjectAttributes:no_chap,use_chap", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:no_chap", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san:ConflictingObjectAttributes:no_chap,use_chap", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage_driver_name", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san:RequiredObjectAttributes:storage_driver_name,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "type": "requires"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--username", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san:RequiredObjectAttributes:storage_driver_name,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san"], "schema_version": 1, "sections": [{"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san client certificate"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_certificate", "description": "Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "client_certificate"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san client private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:client_private_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:client_private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:client_private_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "client_private_key"], "syntax": "block", "type": "object"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san data lif dns name"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--data_lif_dns_name", "description": "Exclusive with Backend Data LIF IP Address's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "data_lif_dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san data lif ip"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--data_lif_ip", "description": "Exclusive with Backend Data LIF IP Address is reachable at the given IP address.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "data_lif_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san igroup name"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--igroup_name", "description": "Name of the igroup for SAN volumes to use.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "igroup_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san labels"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--labels", "description": "List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san limit aggregate usage"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--limit_aggregate_usage", "description": "Fail provisioning if usage is above this percentage. Not enforced by default.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "limit_aggregate_usage"], "syntax": "attribute", "type": "number"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san limit volume size"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--limit_volume_size", "description": "Fail provisioning if requested volume size in GBi is above this value. Not enforced by default.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "limit_volume_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san management lif dns name"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--management_lif_dns_name", "description": "Exclusive with Backend Management LIF IP Address's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "management_lif_dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san management lif ip"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--management_lif_ip", "description": "Exclusive with Backend Management LIF IP Address is reachable at the given IP address.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "management_lif_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san no chap"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:no_chap", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "no_chap"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:password", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "password"], "syntax": "block", "type": "object"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san region"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--region", "description": "Virtual Pool Region.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "region"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san storage"], "anchor": "section", "description": "List of Virtual Storage Pool definitions which are referred back by Storage Class label match selection.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:storage", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage"], "syntax": "block", "type": "object"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san storage driver name"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage_driver_name", "description": "Configuration of Backend Name.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage_driver_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san storage prefix"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage_prefix", "description": "Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "storage_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san svm"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--svm", "description": "Storage virtual machine to use. Derived if an SVM managementLIF is specified.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "svm"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san trusted ca certificate"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--trusted_ca_certificate", "description": "Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based auth..", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "trusted_ca_certificate"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san use chap"], "anchor": "section", "description": "Device NetApp Backend ONTAP SAN CHAP configuration OPTIONS for enabled CHAP.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:use_chap", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "use_chap"], "syntax": "block", "type": "object"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san username"], "anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--username", "description": "Username to connect to the cluster/SVM.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "username"], "syntax": "attribute", "type": "string"}, {"aliases": ["custom storage config storage device list storage devices netapp trident netapp backend ontap san volume defaults"], "anchor": "section", "description": "It controls how each volume is provisioned by default using these OPTIONS in a special section of the configuration.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--adaptive_qos_policy", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,no_qos", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--adaptive_qos_policy", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--qos_policy", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "type": "conflicts"}, {"anchor": "schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--volume_defaults--qos_policy", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults:ConflictingObjectAttributes:no_qos,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,no_qos", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults:no_qos", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults:ConflictingObjectAttributes:no_qos,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_storage_config:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_san:volume_defaults:no_qos", "type": "conflicts"}], "schema_path": ["custom_storage_config", "storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_san", "volume_defaults"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Configuration of storage backend for NetApp ONTAP SAN.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_storage_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/)
- [custom_storage_config.storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/)
- [custom_storage_config.storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/)
- custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP SAN.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip"),
  validators.ConflictingObjectAttributes("no_chap",
    "use_chap")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-chap_choice": "[\"no_chap\",\"use_chap\"]",
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_san {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--client_certificate"></a>

### client_certificate property

Type: `"string"`. Optional.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [client_private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/client_private_key/): complete subsection reference.

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--data_lif_dns_name"></a>

### data_lif_dns_name property

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--data_lif_ip"></a>

### data_lif_ip property

Type: `"string"`. Optional.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--igroup_name"></a>

### igroup_name property

Type: `"string"`. Optional.

Name of the igroup for SAN volumes to use.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":20},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"20\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 20
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "20",
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "20",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--limit_aggregate_usage"></a>

### limit_aggregate_usage property

Type: `"number"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 100,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--limit_volume_size"></a>

### limit_volume_size property

Type: `"number"`. Optional.

Fail provisioning if requested volume size in GBi is above this value. Not enforced by default.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--management_lif_dns_name"></a>

### management_lif_dns_name property

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--management_lif_ip"></a>

### management_lif_ip property

Type: `"string"`. Optional.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

- [no_chap](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/no_chap/): complete subsection reference.

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/password/): complete subsection reference.

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--region"></a>

### region property

Type: `"string"`. Optional.

Backend Region. Virtual Pool Region.

Upstream description:

Virtual Pool Region.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/storage/): complete subsection reference.

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage_driver_name"></a>

### storage_driver_name property

Type: `"string"`. Optional.

\[Enum: ontap-san|ontap-san-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-san\`, \`ontap-san-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-san",
    "ontap-san-economy",
    "ontap-nas-flexgroup"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-san\\\",\\\"ontap-san-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--storage_prefix"></a>

### storage_prefix property

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 80),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 80,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 80,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "80",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--svm"></a>

### svm property

Type: `"string"`. Optional.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--trusted_ca_certificate"></a>

### trusted_ca_certificate property

Type: `"string"`. Optional.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(8192),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8192"
  }
}
```

- [use_chap](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/): complete subsection reference.

<a id="schema-custom_storage_config--storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_san--username"></a>

### username property

Type: `"string"`. Optional.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [volume_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/volume_defaults/): complete subsection reference.

## Next pages

- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.client_private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/client_private_key/)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.no_chap](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/no_chap/)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/password/)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/storage/)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.use_chap](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/use_chap/)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_san.volume_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_san/volume_defaults/)
- [custom_storage_config.storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_storage_config/storage_device_list/storage_devices/netapp_trident/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
