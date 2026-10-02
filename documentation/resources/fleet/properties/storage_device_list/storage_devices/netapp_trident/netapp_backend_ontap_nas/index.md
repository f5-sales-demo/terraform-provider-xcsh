---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas"
subcategory: ""
description: "Configuration of storage backend for NetApp ONTAP NAS."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "storage device list storage devices netapp trident netapp backend ontap nas", "upstream servers"], "body_bytes": 22440, "body_sha256": "sha256:42320c84663f337ab5330b59872b0ae988fc5a0a96319494ce604ccf038ac980", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:auto_export_cidrs", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:password", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage", "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "parent_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "path": "documentation/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0212303102231111-0123213211300022-2033323003020332-1102233001310032-3231013001133100-1303333132211323-0201123013213012-0223113021221131", "registry_path": "docs/guides/resources--fleet--reference--group-003.md", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_dns_name", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas:ConflictingObjectAttributes:data_lif_dns_name,data_lif_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_ip", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas:ConflictingObjectAttributes:data_lif_dns_name,data_lif_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_dns_name", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas:ConflictingObjectAttributes:management_lif_dns_name,management_lif_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_ip", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas:ConflictingObjectAttributes:management_lif_dns_name,management_lif_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_driver_name", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas:RequiredObjectAttributes:storage_driver_name,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "type": "requires"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--username", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas:RequiredObjectAttributes:storage_driver_name,username", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas"], "schema_version": 1, "sections": [{"aliases": ["auto export cidrs"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:auto_export_cidrs", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "auto_export_cidrs"], "syntax": "block", "type": "object"}, {"aliases": ["auto export policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--auto_export_policy", "description": "Enable automatic export policy creation and updating.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "auto_export_policy"], "syntax": "attribute", "type": "bool"}, {"aliases": ["backend name", "backend servers", "origin servers", "upstream servers"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--backend_name", "description": "Configuration of Backend Name. Driver is name + \"_\" + dataLIF.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "backend_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["cert", "certificate", "client certificate", "existing certificates", "tls certificates"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_certificate", "description": "Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "client_certificate"], "syntax": "attribute", "type": "string"}, {"aliases": ["client private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key:clear_secret_info", "type": "conflicts"}], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "client_private_key"], "syntax": "block", "type": "object"}, {"aliases": ["backend servers", "data lif dns name", "origin servers", "upstream servers"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_dns_name", "description": "Exclusive with Backend Data LIF IP Address's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "data_lif_dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["backend servers", "data lif ip", "origin servers", "upstream servers"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_ip", "description": "Exclusive with Backend Data LIF IP Address is reachable at the given IP address.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "data_lif_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["labels"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--labels", "description": "List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["limit aggregate usage"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--limit_aggregate_usage", "description": "Fail provisioning if usage is above this percentage. Not enforced by default.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "limit_aggregate_usage"], "syntax": "attribute", "type": "string"}, {"aliases": ["limit volume size"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--limit_volume_size", "description": "Fail provisioning if requested volume size is above this value. Not enforced by default.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "limit_volume_size"], "syntax": "attribute", "type": "string"}, {"aliases": ["backend servers", "management lif dns name", "origin servers", "upstream servers"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_dns_name", "description": "Exclusive with Backend Management LIF IP Address's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "management_lif_dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["backend servers", "management lif ip", "origin servers", "upstream servers"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_ip", "description": "Exclusive with Backend Management LIF IP Address is reachable at the given IP address.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "management_lif_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["nfs mount options"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--nfs_mount_options", "description": "Comma-separated list of NFS mount OPTIONS. Not enforced by default.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "nfs_mount_options"], "syntax": "attribute", "type": "string"}, {"aliases": ["password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:password", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "password"], "syntax": "block", "type": "object"}, {"aliases": ["region"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--region", "description": "Virtual Pool Region.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "region"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage"], "anchor": "section", "description": "List of Virtual Storage Pool definitions which are referred back by Storage Class label match selection.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage"], "syntax": "block", "type": "object"}, {"aliases": ["backend servers", "origin servers", "storage driver name", "upstream servers"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_driver_name", "description": "Configuration of Backend Name.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage_driver_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage prefix"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_prefix", "description": "Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["svm"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--svm", "description": "Storage virtual machine to use. Derived if an SVM managementLIF is specified.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "svm"], "syntax": "attribute", "type": "string"}, {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "trusted ca certificate"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--trusted_ca_certificate", "description": "Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based auth..", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "trusted_ca_certificate"], "syntax": "attribute", "type": "string"}, {"aliases": ["username"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--username", "description": "Username to connect to the cluster/SVM.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "username"], "syntax": "attribute", "type": "string"}, {"aliases": ["volume defaults"], "anchor": "section", "description": "It controls how each volume is provisioned by default using these OPTIONS in a special section of the configuration.", "document_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--adaptive_qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,no_qos", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--adaptive_qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults", "type": "conflicts"}, {"anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--volume_defaults--qos_policy", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults:ConflictingObjectAttributes:no_qos,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults:ConflictingObjectAttributes:adaptive_qos_policy,no_qos", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults:no_qos", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults:ConflictingObjectAttributes:no_qos,qos_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults:no_qos", "type": "conflicts"}], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "volume_defaults"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Configuration of storage backend for NetApp ONTAP NAS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration of storage backend for NetApp ONTAP NAS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_driver_name",
    "username"),
  validators.ConflictingObjectAttributes("data_lif_dns_name",
    "data_lif_ip"),
  validators.ConflictingObjectAttributes("management_lif_dns_name",
    "management_lif_ip")}
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
  "x-ves-oneof-field-data_lif": "[\"data_lif_dns_name\",\"data_lif_ip\"]",
  "x-ves-oneof-field-management_lif": "[\"management_lif_dns_name\",\"management_lif_ip\"]"
}
```

Terraform syntax:

```terraform
netapp_backend_ontap_nas {
  # Configure direct properties listed below.
}
```

## Direct properties

- [auto_export_cidrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/auto_export_cidrs/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--auto_export_policy"></a>

### auto_export_policy property

Type: `"bool"`. Optional.

Policy configuration for this feature.

Upstream description:

Enable automatic export policy creation and updating.

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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--backend_name"></a>

### backend_name property

Type: `"string"`. Optional.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Upstream description:

Configuration of Backend Name. Driver is name + "\_" + dataLIF.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 50),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 50,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 50,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "50",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_certificate"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [client_private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_dns_name"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_ip"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.

Receipt-pinned upstream constraints:

```json
{
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--limit_aggregate_usage"></a>

### limit_aggregate_usage property

Type: `"string"`. Optional.

Fail provisioning if usage is above this percentage. Not enforced by default.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--limit_volume_size"></a>

### limit_volume_size property

Type: `"string"`. Optional.

Fail provisioning if requested volume size is above this value. Not enforced by default.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_dns_name"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_ip"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--nfs_mount_options"></a>

### nfs_mount_options property

Type: `"string"`. Optional.

Comma-separated list of NFS mount OPTIONS. Not enforced by default.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/password/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--region"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/storage/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_driver_name"></a>

### storage_driver_name property

Type: `"string"`. Optional.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ontap-nas",
    "ontap-nas-economy",
    "ontap-nas-flexgroup"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ontap-nas\\\",\\\"ontap-nas-economy\\\",\\\"ontap-nas-flexgroup\\\"]"
  }
}
```

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_prefix"></a>

### storage_prefix property

Type: `"string"`. Optional.

Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--svm"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--trusted_ca_certificate"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--username"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [volume_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/volume_defaults/): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/auto_export_cidrs/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/password/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/storage/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/volume_defaults/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
