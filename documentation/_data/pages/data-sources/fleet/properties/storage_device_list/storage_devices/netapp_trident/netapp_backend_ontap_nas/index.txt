---
page_title: "storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas"
subcategory: ""
description: "Configuration of storage backend for NetApp ONTAP NAS."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "storage device list storage devices netapp trident netapp backend ontap nas", "upstream servers"], "body_bytes": 20504, "body_sha256": "sha256:569cd65a1906a55b6acceeca6db67b602797fb618fa2432ed4583c3022689c47", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:auto_export_cidrs", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:password", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage", "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "parent_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident", "path": "documentation/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2233200200220103-0221020102010211-2022202313031323-0313221223321323-0113130210031322-2012300313300312-0022110121232021-2131310032303133", "registry_path": "docs/guides/data-sources--fleet--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas"], "schema_version": 1, "sections": [{"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas auto export cidrs"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:auto_export_cidrs", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "auto_export_cidrs"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas auto export policy"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--auto_export_policy", "description": "Enable automatic export policy creation and updating.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "auto_export_policy"], "syntax": "attribute", "type": "bool"}, {"aliases": ["backend servers", "origin servers", "storage device list storage devices netapp trident netapp backend ontap nas backend name", "upstream servers"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--backend_name", "description": "Configuration of Backend Name. Driver is name + \"_\" + dataLIF.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "backend_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas client certificate"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--client_certificate", "description": "Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "client_certificate"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas client private key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:client_private_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "client_private_key"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas data lif dns name"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_dns_name", "description": "Exclusive with Backend Data LIF IP Address's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "data_lif_dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas data lif ip"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_ip", "description": "Exclusive with Backend Data LIF IP Address is reachable at the given IP address.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "data_lif_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas labels"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--labels", "description": "List of labels for Storage Device used in NetApp ONTAP. It is used for storage class selection.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas limit aggregate usage"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--limit_aggregate_usage", "description": "Fail provisioning if usage is above this percentage. Not enforced by default.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "limit_aggregate_usage"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas limit volume size"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--limit_volume_size", "description": "Fail provisioning if requested volume size is above this value. Not enforced by default.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "limit_volume_size"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas management lif dns name"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_dns_name", "description": "Exclusive with Backend Management LIF IP Address's IP address is discovered using DNS name resolution. The name given here is fully qualified domain name.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "management_lif_dns_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas management lif ip"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--management_lif_ip", "description": "Exclusive with Backend Management LIF IP Address is reachable at the given IP address.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "management_lif_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas nfs mount options"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--nfs_mount_options", "description": "Comma-separated list of NFS mount OPTIONS. Not enforced by default.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "nfs_mount_options"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:password", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "password"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas region"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--region", "description": "Virtual Pool Region.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "region"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas storage"], "anchor": "section", "description": "List of Virtual Storage Pool definitions which are referred back by Storage Class label match selection.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas storage driver name"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_driver_name", "description": "Configuration of Backend Name.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage_driver_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas storage prefix"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_prefix", "description": "Prefix used when provisioning new volumes in the SVM. Once set this cannot be updated.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "storage_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas svm"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--svm", "description": "Storage virtual machine to use. Derived if an SVM managementLIF is specified.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "svm"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas trusted ca certificate"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--trusted_ca_certificate", "description": "Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based auth..", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "trusted_ca_certificate"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas username"], "anchor": "schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--username", "description": "Username to connect to the cluster/SVM.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "username"], "syntax": "attribute", "type": "string"}, {"aliases": ["storage device list storage devices netapp trident netapp backend ontap nas volume defaults"], "anchor": "section", "description": "It controls how each volume is provisioned by default using these OPTIONS in a special section of the configuration.", "document_id": "xcsh-docs:data-sources:fleet:properties:storage_device_list:storage_devices:netapp_trident:netapp_backend_ontap_nas:volume_defaults", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["storage_device_list", "storage_devices", "netapp_trident", "netapp_backend_ontap_nas", "volume_defaults"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration of storage backend for NetApp ONTAP NAS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/)
- [storage_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/)
- [storage_device_list.storage_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas

<a id="section"></a>

Type: `"single"`. Computed.

Configuration of storage backend for NetApp ONTAP NAS.

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

## Direct properties

- [auto_export_cidrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/auto_export_cidrs/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--auto_export_policy"></a>

### auto_export_policy property

Type: `"bool"`. Computed.

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

Type: `"string"`. Computed.

Configuration of Backend Name. Driver is name + '\_' + dataLIF.

Upstream description:

Configuration of Backend Name. Driver is name + "\_" + dataLIF.

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

Type: `"string"`. Computed.

Please Enter Base64-encoded value of client certificate. Used for certificate-based auth.

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

- [client_private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--data_lif_dns_name"></a>

### data_lif_dns_name property

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[data\_lif\_ip\] Backend Data LIF IP Address's IP address is discovered using DNS
name resolution. The name given here is fully qualified domain name.

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

Type: `"string"`. Computed.

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

Upstream description:

Exclusive with \[data\_lif\_dns\_name\] Backend Data LIF IP Address is reachable at the given IP
address.

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

Type: `["map", "string"]`. Computed.

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

Type: `"string"`. Computed.

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

Type: `"string"`. Computed.

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

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

Upstream description:

Exclusive with \[management\_lif\_ip\] Backend Management LIF IP Address's IP address is discovered
using DNS name resolution. The name given here is fully qualified domain name.

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

Type: `"string"`. Computed.

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

Upstream description:

Exclusive with \[management\_lif\_dns\_name\] Backend Management LIF IP Address is reachable at the
given IP address.

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

Type: `"string"`. Computed.

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

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/password/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--region"></a>

### region property

Type: `"string"`. Computed.

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

- [storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/storage/): complete subsection reference.

<a id="schema-storage_device_list--storage_devices--netapp_trident--netapp_backend_ontap_nas--storage_driver_name"></a>

### storage_driver_name property

Type: `"string"`. Computed.

\[Enum: ontap-nas|ontap-nas-economy|ontap-nas-flexgroup\] Storage Backend Driver. Configuration of
Backend Name. Possible values are \`ontap-nas\`, \`ontap-nas-economy\`, \`ontap-nas-flexgroup\`.

Upstream description:

Configuration of Backend Name.

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

Type: `"string"`. Computed.

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

Type: `"string"`. Computed.

Storage virtual machine to use. Derived if an SVM managementLIF is specified.

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

Type: `"string"`. Computed.

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth.

Upstream description:

Please Enter Base64-encoded value of trusted CA certificate. Optional. Used for certificate-based
auth..

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

Type: `"string"`. Computed.

Username. Username to connect to the cluster/SVM.

Upstream description:

Username to connect to the cluster/SVM.

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

- [volume_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/volume_defaults/): complete subsection reference.

## Next pages

- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.auto_export_cidrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/auto_export_cidrs/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.client_private_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/client_private_key/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/password/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/storage/)
- [storage_device_list.storage_devices.netapp_trident.netapp_backend_ontap_nas.volume_defaults](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/netapp_backend_ontap_nas/volume_defaults/)
- [storage_device_list.storage_devices.netapp_trident](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/properties/storage_device_list/storage_devices/netapp_trident/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
