---
page_title: "Property reference"
subcategory: "Infrastructure"
description: "Property reference for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": ["azure vnet site"], "body_bytes": 266411, "body_sha256": "sha256:299f92cc7fe1cce72c95121146efb3ec8b2bf970c3f4fc02ffc5ae60db737de1", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:azure_vnet_site:properties:admin_password", "xcsh-docs:data-sources:azure_vnet_site:properties:azure_cred", "xcsh-docs:data-sources:azure_vnet_site:properties:block_all_services", "xcsh-docs:data-sources:azure_vnet_site:properties:blocked_services", "xcsh-docs:data-sources:azure_vnet_site:properties:coordinates", "xcsh-docs:data-sources:azure_vnet_site:properties:custom_dns", "xcsh-docs:data-sources:azure_vnet_site:properties:default_blocked_services", "xcsh-docs:data-sources:azure_vnet_site:properties:disable_encryption", "xcsh-docs:data-sources:azure_vnet_site:properties:enable_encryption", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw", "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar", "xcsh-docs:data-sources:azure_vnet_site:properties:kubernetes_upgrade_drain", "xcsh-docs:data-sources:azure_vnet_site:properties:log_receiver", "xcsh-docs:data-sources:azure_vnet_site:properties:logs_streaming_disabled", "xcsh-docs:data-sources:azure_vnet_site:properties:no_worker_nodes", "xcsh-docs:data-sources:azure_vnet_site:properties:offline_survivability_mode", "xcsh-docs:data-sources:azure_vnet_site:properties:os", "xcsh-docs:data-sources:azure_vnet_site:properties:sw", "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster", "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar", "xcsh-docs:data-sources:azure_vnet_site:properties:waf_signatures"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:azure_vnet_site:reference", "parent_id": "xcsh-docs:data-sources:azure_vnet_site:fundamentals", "path": "documentation/data-sources/azure_vnet_site/properties/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-2210100210332112-1233221121133301-0120203222303210-2221022223110220-0313230312203322-2232200300001103-2301220003133113-1333232230300113", "registry_path": "docs/guides/data-sources--azure_vnet_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["address"], "anchor": "schema-address", "description": "Site's geographical address that can be used to determine its latitude and longitude.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address"], "syntax": "attribute", "type": "string"}, {"aliases": ["admin password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:admin_password", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["admin_password"], "syntax": "attribute", "type": "object"}, {"aliases": ["alternate region"], "anchor": "schema-alternate_region", "description": "Exclusive with Name of the Azure region which does not support availability zones.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["alternate_region"], "syntax": "attribute", "type": "string"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["azure cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:azure_cred", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["azure_cred"], "syntax": "attribute", "type": "object"}, {"aliases": ["azure region"], "anchor": "schema-azure_region", "description": "Exclusive with Name of the Azure region which supports availability zones.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["azure_region"], "syntax": "attribute", "type": "string"}, {"aliases": ["block all services"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:block_all_services", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["block_all_services"], "syntax": "attribute", "type": "object"}, {"aliases": ["blocked services"], "anchor": "section", "description": "Disable node local services on this site. Note: The chosen services will GET disabled on all nodes in the site.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:blocked_services", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["blocked_services"], "syntax": "attribute", "type": "object"}, {"aliases": ["coordinates"], "anchor": "section", "description": "Coordinates of the site which provides the site physical location.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:coordinates", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["coordinates"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom dns"], "anchor": "section", "description": "Custom DNS is the configured for specify CE site.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:custom_dns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_dns"], "syntax": "attribute", "type": "object"}, {"aliases": ["default blocked services"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:default_blocked_services", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_blocked_services"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable encryption"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:disable_encryption", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable_encryption"], "syntax": "attribute", "type": "object"}, {"aliases": ["disk size"], "anchor": "schema-disk_size", "description": "Disk size to be used for this instance in GiB. 80 is 80 GiB.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disk_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["enable encryption"], "anchor": "section", "description": "Information related to disk encryption.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:enable_encryption", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["enable_encryption"], "syntax": "attribute", "type": "object"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["ingress egress gw"], "anchor": "section", "description": "Two interface Azure ingress/egress site.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress egress gw ar"], "anchor": "section", "description": "Two interface Azure ingress/egress site on Alternate Region with no support for zones.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_egress_gw_ar", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw_ar"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress gw"], "anchor": "section", "description": "Single interface Azure ingress site on on Recommended Region.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_gw"], "syntax": "attribute", "type": "object"}, {"aliases": ["ingress gw ar"], "anchor": "section", "description": "Single interface Azure ingress site.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:ingress_gw_ar", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_gw_ar"], "syntax": "attribute", "type": "object"}, {"aliases": ["kubernetes upgrade drain"], "anchor": "section", "description": "Specify how worker nodes within a site will be upgraded.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:kubernetes_upgrade_drain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["kubernetes_upgrade_drain"], "syntax": "attribute", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["log receiver"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:log_receiver", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["log_receiver"], "syntax": "attribute", "type": "object"}, {"aliases": ["logs streaming disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:logs_streaming_disabled", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["logs_streaming_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["machine type"], "anchor": "schema-machine_type", "description": "Select Instance size based on performance needed. The default setting for Accelerated Networking is enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable the setting under, Select Ingress Gateway or Ingress/Egress Gateway > advanced OPTIONS.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["machine_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["no worker nodes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:no_worker_nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_worker_nodes"], "syntax": "attribute", "type": "object"}, {"aliases": ["nodes per az"], "anchor": "schema-nodes_per_az", "description": "Exclusive with Desired Worker Nodes Per AZ. Max limit is up to 21.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["nodes_per_az"], "syntax": "attribute", "type": "number"}, {"aliases": ["offline survivability mode"], "anchor": "section", "description": "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:offline_survivability_mode", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["offline_survivability_mode"], "syntax": "attribute", "type": "object"}, {"aliases": ["os"], "anchor": "section", "description": "Select the F5XC Operating System Version for the site. By default, latest available OS Version will be used. Refer to release notes to find required released OS versions.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:os", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["os"], "syntax": "attribute", "type": "object"}, {"aliases": ["resource group"], "anchor": "schema-resource_group", "description": "Azure resource group for resources that will be created.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["resource_group"], "syntax": "attribute", "type": "string"}, {"aliases": ["ssh key"], "anchor": "schema-ssh_key", "description": "Public SSH key for accessing the site.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ssh_key"], "syntax": "attribute", "type": "string"}, {"aliases": ["sw"], "anchor": "section", "description": "Select the F5XC Software Version for the site. By default, latest available F5XC Software Version will be used. Refer to release notes to find required released SW versions.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:sw", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["sw"], "syntax": "attribute", "type": "object"}, {"aliases": ["tags"], "anchor": "schema-tags", "description": "Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify, organize, search for, and filter resources in Azure console.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["tags"], "syntax": "attribute", "type": "map"}, {"aliases": ["total nodes"], "anchor": "schema-total_nodes", "description": "Exclusive with Total number of worker nodes to be deployed across all AZ's used in the Site.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["total_nodes"], "syntax": "attribute", "type": "number"}, {"aliases": ["vnet"], "anchor": "section", "description": "This defines choice about Azure VNet for a view.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:vnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["vnet"], "syntax": "attribute", "type": "object"}, {"aliases": ["voltstack cluster"], "anchor": "section", "description": "App Stack Cluster of single interface Azure nodes.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["voltstack cluster ar"], "anchor": "section", "description": "App Stack Cluster of single interface Azure nodes.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:voltstack_cluster_ar", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster_ar"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf signatures"], "anchor": "section", "description": "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes.", "document_id": "xcsh-docs:data-sources:azure_vnet_site:properties:waf_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_signatures"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/azure_vnet_site/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_azure_vnet_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Computed.

Site's geographical address that can be used to determine its latitude and longitude.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/): complete subsection reference.

<a id="schema-alternate_region"></a>

### alternate_region property

Type: `"string"`. Computed.

\[OneOf: alternate\_region, azure\_region\] Exclusive with \[azure\_region\] Name of the Azure
region which does not support availability zones.

Upstream description:

Exclusive with \[azure\_region\] Name of the Azure region which does not support availability zones.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

OneOf alternatives in this subsection:

- [alternate_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-alternate_region)
- [azure_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-azure_region)

Select alternatives according to the provider validators above.

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Upstream description:

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

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
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [azure_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/azure_cred/): complete subsection reference.

<a id="schema-azure_region"></a>

### azure_region property

Type: `"string"`. Computed.

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

Upstream description:

Exclusive with \[alternate\_region\] Name of the Azure region which supports availability zones.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [block_all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/block_all_services/): complete subsection reference.

- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/blocked_services/): complete subsection reference.

- [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/coordinates/): complete subsection reference.

- [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/custom_dns/): complete subsection reference.

- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/default_blocked_services/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Computed.

Description of the AzureVNETSite.

Upstream description:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/disable_encryption/): complete subsection reference.

<a id="schema-disk_size"></a>

### disk_size property

Type: `"number"`. Computed.

Disk size to be used for this instance in GiB. 80 is 80 GiB. Server applies default when omitted.

Upstream description:

Disk size to be used for this instance in GiB. 80 is 80 GiB.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

- [enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/enable_encryption/): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/): complete subsection reference.

- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/): complete subsection reference.

- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/): complete subsection reference.

- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/): complete subsection reference.

- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Upstream description:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

- [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/log_receiver/): complete subsection reference.

- [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/logs_streaming_disabled/): complete subsection reference.

<a id="schema-machine_type"></a>

### machine_type property

Type: `"string"`. Computed.

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

Upstream description:

Select Instance size based on performance needed. The default setting for Accelerated Networking is
enabled, thus make sure you select a Virtual Machine that supports accelerated networking or disable
the setting under, Select Ingress Gateway or Ingress/Egress Gateway &gt; advanced OPTIONS.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the AzureVNETSite.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace where the AzureVNETSite exists.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

- [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/no_worker_nodes/): complete subsection reference.

<a id="schema-nodes_per_az"></a>

### nodes_per_az property

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Upstream description:

Exclusive with \[no\_worker\_nodes total\_nodes\] Desired Worker Nodes Per AZ. Max limit is up to
21.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 21,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "21"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "21"
  }
}
```

- [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/offline_survivability_mode/): complete subsection reference.

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/os/): complete subsection reference.

<a id="schema-resource_group"></a>

### resource_group property

Type: `"string"`. Computed.

Azure resource group for resources that will be created.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-ssh_key"></a>

### ssh_key property

Type: `"string"`. Computed.

Public SSH key. Public SSH key for accessing the site.

Upstream description:

Public SSH key for accessing the site.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "8192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/sw/): complete subsection reference.

<a id="schema-tags"></a>

### tags property

Type: `["map", "string"]`. Computed.

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

Azure Tags is a label consisting of a user-defined key and value. It helps to manage, identify,
organize, search for, and filter resources in Azure console.

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
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "127",
    "ves.io.schema.rules.map.max_pairs": "40",
    "ves.io.schema.rules.map.values.string.max_len": "255"
  }
}
```

<a id="schema-total_nodes"></a>

### total_nodes property

Type: `"number"`. Computed.

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Upstream description:

Exclusive with \[no\_worker\_nodes nodes\_per\_az\] Total number of worker nodes to be deployed
across all AZ's used in the Site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 61,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "61"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "61"
  }
}
```

- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/): complete subsection reference.

- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/): complete subsection reference.

- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/): complete subsection reference.

- [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/waf_signatures/): complete subsection reference.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-address) |
| `admin_password` | [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/#section) |
| `admin_password.blindfold_secret_info` | [admin_password.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/blindfold_secret_info/#section) |
| `admin_password.blindfold_secret_info.decryption_provider` | [admin_password.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/blindfold_secret_info/#schema-admin_password--blindfold_secret_info--decryption_provider) |
| `admin_password.blindfold_secret_info.location` | [admin_password.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/blindfold_secret_info/#schema-admin_password--blindfold_secret_info--location) |
| `admin_password.blindfold_secret_info.store_provider` | [admin_password.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/blindfold_secret_info/#schema-admin_password--blindfold_secret_info--store_provider) |
| `admin_password.clear_secret_info` | [admin_password.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/clear_secret_info/#section) |
| `admin_password.clear_secret_info.provider_ref` | [admin_password.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/clear_secret_info/#schema-admin_password--clear_secret_info--provider_ref) |
| `admin_password.clear_secret_info.url` | [admin_password.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/clear_secret_info/#schema-admin_password--clear_secret_info--url) |
| `alternate_region` | [alternate_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-alternate_region) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-annotations) |
| `azure_cred` | [azure_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/azure_cred/#section) |
| `azure_cred.name` | [azure_cred.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/azure_cred/#schema-azure_cred--name) |
| `azure_cred.namespace` | [azure_cred.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/azure_cred/#schema-azure_cred--namespace) |
| `azure_cred.tenant` | [azure_cred.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/azure_cred/#schema-azure_cred--tenant) |
| `azure_region` | [azure_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-azure_region) |
| `block_all_services` | [block_all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/block_all_services/#section) |
| `blocked_services` | [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/blocked_services/#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/blocked_services/blocked_service/#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/blocked_services/blocked_service/dns/#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/blocked_services/blocked_service/#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/blocked_services/blocked_service/ssh/#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/blocked_services/blocked_service/web_user_interface/#section) |
| `coordinates` | [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/coordinates/#section) |
| `coordinates.latitude` | [coordinates.latitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/coordinates/#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/coordinates/#schema-coordinates--longitude) |
| `custom_dns` | [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/custom_dns/#section) |
| `custom_dns.inside_nameserver` | [custom_dns.inside_nameserver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/custom_dns/#schema-custom_dns--inside_nameserver) |
| `custom_dns.outside_nameserver` | [custom_dns.outside_nameserver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/custom_dns/#schema-custom_dns--outside_nameserver) |
| `default_blocked_services` | [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/default_blocked_services/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-description) |
| `disable_encryption` | [disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/disable_encryption/#section) |
| `disk_size` | [disk_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-disk_size) |
| `enable_encryption` | [enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/enable_encryption/#section) |
| `enable_encryption.disk_encryption_set_id` | [enable_encryption.disk_encryption_set_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/enable_encryption/#schema-enable_encryption--disk_encryption_set_id) |
| `enable_encryption.resource_group` | [enable_encryption.resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/enable_encryption/#schema-enable_encryption--resource_group) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-id) |
| `ingress_egress_gw` | [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/#section) |
| `ingress_egress_gw.accelerated_networking` | [ingress_egress_gw.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/#section) |
| `ingress_egress_gw.accelerated_networking.disable_spec` | [ingress_egress_gw.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/disable_spec/#section) |
| `ingress_egress_gw.accelerated_networking.enable` | [ingress_egress_gw.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/accelerated_networking/enable/#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/enhanced_firewall_policies/#section) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-ingress_egress_gw--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `ingress_egress_gw.active_forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/#section) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--name) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw.active_forward_proxy_policies.forward_proxy_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_forward_proxy_policies/forward_proxy_policies/#schema-ingress_egress_gw--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `ingress_egress_gw.active_network_policies` | [ingress_egress_gw.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/#section) |
| `ingress_egress_gw.active_network_policies.network_policies` | [ingress_egress_gw.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/network_policies/#section) |
| `ingress_egress_gw.active_network_policies.network_policies.name` | [ingress_egress_gw.active_network_policies.network_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/network_policies/#schema-ingress_egress_gw--active_network_policies--network_policies--name) |
| `ingress_egress_gw.active_network_policies.network_policies.namespace` | [ingress_egress_gw.active_network_policies.network_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/network_policies/#schema-ingress_egress_gw--active_network_policies--network_policies--namespace) |
| `ingress_egress_gw.active_network_policies.network_policies.tenant` | [ingress_egress_gw.active_network_policies.network_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/active_network_policies/network_policies/#schema-ingress_egress_gw--active_network_policies--network_policies--tenant) |
| `ingress_egress_gw.az_nodes` | [ingress_egress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/#section) |
| `ingress_egress_gw.az_nodes.azure_az` | [ingress_egress_gw.az_nodes.azure_az](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/#schema-ingress_egress_gw--az_nodes--azure_az) |
| `ingress_egress_gw.az_nodes.inside_subnet` | [ingress_egress_gw.az_nodes.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/inside_subnet/#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet` | [ingress_egress_gw.az_nodes.inside_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet/#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet.subnet_name` | [ingress_egress_gw.az_nodes.inside_subnet.subnet.subnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet/#schema-ingress_egress_gw--az_nodes--inside_subnet--subnet--subnet_name) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw.az_nodes.inside_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet/#schema-ingress_egress_gw--az_nodes--inside_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet.vnet_resource_group` | [ingress_egress_gw.az_nodes.inside_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet/vnet_resource_group/#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/#section) |
| `ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.inside_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/inside_subnet/subnet_param/#schema-ingress_egress_gw--az_nodes--inside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.az_nodes.outside_subnet` | [ingress_egress_gw.az_nodes.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet` | [ingress_egress_gw.az_nodes.outside_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet.subnet_name` | [ingress_egress_gw.az_nodes.outside_subnet.subnet.subnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/#schema-ingress_egress_gw--az_nodes--outside_subnet--subnet--subnet_name) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw.az_nodes.outside_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/#schema-ingress_egress_gw--az_nodes--outside_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet.vnet_resource_group` | [ingress_egress_gw.az_nodes.outside_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet/vnet_resource_group/#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet_param/#section) |
| `ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw.az_nodes.outside_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/az_nodes/outside_subnet/subnet_param/#schema-ingress_egress_gw--az_nodes--outside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.azure_certified_hw` | [ingress_egress_gw.azure_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/#schema-ingress_egress_gw--azure_certified_hw) |
| `ingress_egress_gw.dc_cluster_group_inside_vn` | [ingress_egress_gw.dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/#section) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.name` | [ingress_egress_gw.dc_cluster_group_inside_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/#schema-ingress_egress_gw--dc_cluster_group_inside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_inside_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/#schema-ingress_egress_gw--dc_cluster_group_inside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_inside_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_inside_vn/#schema-ingress_egress_gw--dc_cluster_group_inside_vn--tenant) |
| `ingress_egress_gw.dc_cluster_group_outside_vn` | [ingress_egress_gw.dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/#section) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.name` | [ingress_egress_gw.dc_cluster_group_outside_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/#schema-ingress_egress_gw--dc_cluster_group_outside_vn--name) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw.dc_cluster_group_outside_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/#schema-ingress_egress_gw--dc_cluster_group_outside_vn--namespace) |
| `ingress_egress_gw.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw.dc_cluster_group_outside_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/dc_cluster_group_outside_vn/#schema-ingress_egress_gw--dc_cluster_group_outside_vn--tenant) |
| `ingress_egress_gw.forward_proxy_allow_all` | [ingress_egress_gw.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/forward_proxy_allow_all/#section) |
| `ingress_egress_gw.global_network_list` | [ingress_egress_gw.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections` | [ingress_egress_gw.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#section) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-ingress_egress_gw--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw.hub` | [ingress_egress_gw.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/#section) |
| `ingress_egress_gw.hub.express_route_disabled` | [ingress_egress_gw.hub.express_route_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_disabled/#section) |
| `ingress_egress_gw.hub.express_route_enabled` | [ingress_egress_gw.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/#section) |
| `ingress_egress_gw.hub.express_route_enabled.advertise_to_route_server` | [ingress_egress_gw.hub.express_route_enabled.advertise_to_route_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/advertise_to_route_server/#section) |
| `ingress_egress_gw.hub.express_route_enabled.auto_asn` | [ingress_egress_gw.hub.express_route_enabled.auto_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/auto_asn/#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections` | [ingress_egress_gw.hub.express_route_enabled.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.circuit_id` | [ingress_egress_gw.hub.express_route_enabled.connections.circuit_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/#schema-ingress_egress_gw--hub--express_route_enabled--connections--circuit_id) |
| `ingress_egress_gw.hub.express_route_enabled.connections.metadata` | [ingress_egress_gw.hub.express_route_enabled.connections.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/metadata/#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.metadata.description_spec` | [ingress_egress_gw.hub.express_route_enabled.connections.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/metadata/#schema-ingress_egress_gw--hub--express_route_enabled--connections--metadata--description_spec) |
| `ingress_egress_gw.hub.express_route_enabled.connections.metadata.name` | [ingress_egress_gw.hub.express_route_enabled.connections.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/metadata/#schema-ingress_egress_gw--hub--express_route_enabled--connections--metadata--name) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.decryption_provider` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--decryption_provider) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.location` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--location) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.store_provider` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--store_provider) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/clear_secret_info/#section) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.provider_ref` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/clear_secret_info/#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info--provider_ref) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.url` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/authorized_key/clear_secret_info/#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info--url) |
| `ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.circuit_id` | [ingress_egress_gw.hub.express_route_enabled.connections.other_subscription.circuit_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/other_subscription/#schema-ingress_egress_gw--hub--express_route_enabled--connections--other_subscription--circuit_id) |
| `ingress_egress_gw.hub.express_route_enabled.connections.weight` | [ingress_egress_gw.hub.express_route_enabled.connections.weight](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/connections/#schema-ingress_egress_gw--hub--express_route_enabled--connections--weight) |
| `ingress_egress_gw.hub.express_route_enabled.custom_asn` | [ingress_egress_gw.hub.express_route_enabled.custom_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/#schema-ingress_egress_gw--hub--express_route_enabled--custom_asn) |
| `ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server` | [ingress_egress_gw.hub.express_route_enabled.do_not_advertise_to_route_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/do_not_advertise_to_route_server/#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/gateway_subnet/#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/gateway_subnet/auto/#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/gateway_subnet/subnet/#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/gateway_subnet/subnet/#schema-ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/gateway_subnet/subnet/vnet_resource_group/#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/gateway_subnet/subnet_param/#section) |
| `ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param.ipv4` | [ingress_egress_gw.hub.express_route_enabled.gateway_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/gateway_subnet/subnet_param/#schema-ingress_egress_gw--hub--express_route_enabled--gateway_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/auto/#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/subnet/#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/subnet/#schema-ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/subnet/vnet_resource_group/#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/subnet_param/#section) |
| `ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param.ipv4` | [ingress_egress_gw.hub.express_route_enabled.route_server_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/route_server_subnet/subnet_param/#schema-ingress_egress_gw--hub--express_route_enabled--route_server_subnet--subnet_param--ipv4) |
| `ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route` | [ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/site_registration_over_express_route/#section) |
| `ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route.cloudlink_network_name` | [ingress_egress_gw.hub.express_route_enabled.site_registration_over_express_route.cloudlink_network_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/site_registration_over_express_route/#schema-ingress_egress_gw--hub--express_route_enabled--site_registration_over_express_route--cloudlink_network_name) |
| `ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet` | [ingress_egress_gw.hub.express_route_enabled.site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/site_registration_over_internet/#section) |
| `ingress_egress_gw.hub.express_route_enabled.sku_ergw1az` | [ingress_egress_gw.hub.express_route_enabled.sku_ergw1az](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/sku_ergw1az/#section) |
| `ingress_egress_gw.hub.express_route_enabled.sku_ergw2az` | [ingress_egress_gw.hub.express_route_enabled.sku_ergw2az](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/sku_ergw2az/#section) |
| `ingress_egress_gw.hub.express_route_enabled.sku_high_perf` | [ingress_egress_gw.hub.express_route_enabled.sku_high_perf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/sku_high_perf/#section) |
| `ingress_egress_gw.hub.express_route_enabled.sku_standard` | [ingress_egress_gw.hub.express_route_enabled.sku_standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/express_route_enabled/sku_standard/#section) |
| `ingress_egress_gw.hub.spoke_vnets` | [ingress_egress_gw.hub.spoke_vnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/#section) |
| `ingress_egress_gw.hub.spoke_vnets.auto` | [ingress_egress_gw.hub.spoke_vnets.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/auto/#section) |
| `ingress_egress_gw.hub.spoke_vnets.labels` | [ingress_egress_gw.hub.spoke_vnets.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/labels/#section) |
| `ingress_egress_gw.hub.spoke_vnets.manual` | [ingress_egress_gw.hub.spoke_vnets.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/manual/#section) |
| `ingress_egress_gw.hub.spoke_vnets.vnet` | [ingress_egress_gw.hub.spoke_vnets.vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/#section) |
| `ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing` | [ingress_egress_gw.hub.spoke_vnets.vnet.f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/f5_orchestrated_routing/#section) |
| `ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing` | [ingress_egress_gw.hub.spoke_vnets.vnet.manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/manual_routing/#section) |
| `ingress_egress_gw.hub.spoke_vnets.vnet.resource_group` | [ingress_egress_gw.hub.spoke_vnets.vnet.resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/#schema-ingress_egress_gw--hub--spoke_vnets--vnet--resource_group) |
| `ingress_egress_gw.hub.spoke_vnets.vnet.vnet_name` | [ingress_egress_gw.hub.spoke_vnets.vnet.vnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/hub/spoke_vnets/vnet/#schema-ingress_egress_gw--hub--spoke_vnets--vnet--vnet_name) |
| `ingress_egress_gw.inside_static_routes` | [ingress_egress_gw.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list` | [ingress_egress_gw.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/nexthop/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.inside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/inside_static_routes/static_route_list/#schema-ingress_egress_gw--inside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.no_dc_cluster_group` | [ingress_egress_gw.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_dc_cluster_group/#section) |
| `ingress_egress_gw.no_forward_proxy` | [ingress_egress_gw.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_forward_proxy/#section) |
| `ingress_egress_gw.no_global_network` | [ingress_egress_gw.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_global_network/#section) |
| `ingress_egress_gw.no_inside_static_routes` | [ingress_egress_gw.no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_inside_static_routes/#section) |
| `ingress_egress_gw.no_network_policy` | [ingress_egress_gw.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_network_policy/#section) |
| `ingress_egress_gw.no_outside_static_routes` | [ingress_egress_gw.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/no_outside_static_routes/#section) |
| `ingress_egress_gw.not_hub` | [ingress_egress_gw.not_hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/not_hub/#section) |
| `ingress_egress_gw.outside_static_routes` | [ingress_egress_gw.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list` | [ingress_egress_gw.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/nexthop/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw.outside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/outside_static_routes/static_route_list/#schema-ingress_egress_gw--outside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw.performance_enhancement_mode` | [ingress_egress_gw.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/#section) |
| `ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/#section) |
| `ingress_egress_gw.sm_connection_public_ip` | [ingress_egress_gw.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/sm_connection_public_ip/#section) |
| `ingress_egress_gw.sm_connection_pvt_ip` | [ingress_egress_gw.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/sm_connection_pvt_ip/#section) |
| `ingress_egress_gw_ar` | [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/#section) |
| `ingress_egress_gw_ar.accelerated_networking` | [ingress_egress_gw_ar.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/accelerated_networking/#section) |
| `ingress_egress_gw_ar.accelerated_networking.disable_spec` | [ingress_egress_gw_ar.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/accelerated_networking/disable_spec/#section) |
| `ingress_egress_gw_ar.accelerated_networking.enable` | [ingress_egress_gw_ar.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/accelerated_networking/enable/#section) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies` | [ingress_egress_gw_ar.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_enhanced_firewall_policies/#section) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies` | [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/#section) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [ingress_egress_gw_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-ingress_egress_gw_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `ingress_egress_gw_ar.active_forward_proxy_policies` | [ingress_egress_gw_ar.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/#section) |
| `ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies` | [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/forward_proxy_policies/#section) |
| `ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.name` | [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/forward_proxy_policies/#schema-ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies--name) |
| `ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.namespace` | [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/forward_proxy_policies/#schema-ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.tenant` | [ingress_egress_gw_ar.active_forward_proxy_policies.forward_proxy_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_forward_proxy_policies/forward_proxy_policies/#schema-ingress_egress_gw_ar--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `ingress_egress_gw_ar.active_network_policies` | [ingress_egress_gw_ar.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_network_policies/#section) |
| `ingress_egress_gw_ar.active_network_policies.network_policies` | [ingress_egress_gw_ar.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_network_policies/network_policies/#section) |
| `ingress_egress_gw_ar.active_network_policies.network_policies.name` | [ingress_egress_gw_ar.active_network_policies.network_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_network_policies/network_policies/#schema-ingress_egress_gw_ar--active_network_policies--network_policies--name) |
| `ingress_egress_gw_ar.active_network_policies.network_policies.namespace` | [ingress_egress_gw_ar.active_network_policies.network_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_network_policies/network_policies/#schema-ingress_egress_gw_ar--active_network_policies--network_policies--namespace) |
| `ingress_egress_gw_ar.active_network_policies.network_policies.tenant` | [ingress_egress_gw_ar.active_network_policies.network_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/active_network_policies/network_policies/#schema-ingress_egress_gw_ar--active_network_policies--network_policies--tenant) |
| `ingress_egress_gw_ar.azure_certified_hw` | [ingress_egress_gw_ar.azure_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/#schema-ingress_egress_gw_ar--azure_certified_hw) |
| `ingress_egress_gw_ar.dc_cluster_group_inside_vn` | [ingress_egress_gw_ar.dc_cluster_group_inside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_inside_vn/#section) |
| `ingress_egress_gw_ar.dc_cluster_group_inside_vn.name` | [ingress_egress_gw_ar.dc_cluster_group_inside_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_inside_vn/#schema-ingress_egress_gw_ar--dc_cluster_group_inside_vn--name) |
| `ingress_egress_gw_ar.dc_cluster_group_inside_vn.namespace` | [ingress_egress_gw_ar.dc_cluster_group_inside_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_inside_vn/#schema-ingress_egress_gw_ar--dc_cluster_group_inside_vn--namespace) |
| `ingress_egress_gw_ar.dc_cluster_group_inside_vn.tenant` | [ingress_egress_gw_ar.dc_cluster_group_inside_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_inside_vn/#schema-ingress_egress_gw_ar--dc_cluster_group_inside_vn--tenant) |
| `ingress_egress_gw_ar.dc_cluster_group_outside_vn` | [ingress_egress_gw_ar.dc_cluster_group_outside_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_outside_vn/#section) |
| `ingress_egress_gw_ar.dc_cluster_group_outside_vn.name` | [ingress_egress_gw_ar.dc_cluster_group_outside_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_outside_vn/#schema-ingress_egress_gw_ar--dc_cluster_group_outside_vn--name) |
| `ingress_egress_gw_ar.dc_cluster_group_outside_vn.namespace` | [ingress_egress_gw_ar.dc_cluster_group_outside_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_outside_vn/#schema-ingress_egress_gw_ar--dc_cluster_group_outside_vn--namespace) |
| `ingress_egress_gw_ar.dc_cluster_group_outside_vn.tenant` | [ingress_egress_gw_ar.dc_cluster_group_outside_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/dc_cluster_group_outside_vn/#schema-ingress_egress_gw_ar--dc_cluster_group_outside_vn--tenant) |
| `ingress_egress_gw_ar.forward_proxy_allow_all` | [ingress_egress_gw_ar.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/forward_proxy_allow_all/#section) |
| `ingress_egress_gw_ar.global_network_list` | [ingress_egress_gw_ar.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections` | [ingress_egress_gw_ar.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/sli_to_global_dr/#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [ingress_egress_gw_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/slo_to_global_dr/#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#section) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [ingress_egress_gw_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-ingress_egress_gw_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `ingress_egress_gw_ar.hub` | [ingress_egress_gw_ar.hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/#section) |
| `ingress_egress_gw_ar.hub.express_route_disabled` | [ingress_egress_gw_ar.hub.express_route_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_disabled/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled` | [ingress_egress_gw_ar.hub.express_route_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server` | [ingress_egress_gw_ar.hub.express_route_enabled.advertise_to_route_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/advertise_to_route_server/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.auto_asn` | [ingress_egress_gw_ar.hub.express_route_enabled.auto_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/auto_asn/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections` | [ingress_egress_gw_ar.hub.express_route_enabled.connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.circuit_id` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.circuit_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--circuit_id) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/metadata/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata.description_spec` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/metadata/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata--description_spec) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata.name` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.metadata.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/metadata/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--metadata--name) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.decryption_provider` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.decryption_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--decryption_provider) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.location` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.location](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--location) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.store_provider` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.blindfold_secret_info.store_provider](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/blindfold_secret_info/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--blindfold_secret_info--store_provider) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/clear_secret_info/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.provider_ref` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.provider_ref](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/clear_secret_info/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info--provider_ref) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.url` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.authorized_key.clear_secret_info.url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/authorized_key/clear_secret_info/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--authorized_key--clear_secret_info--url) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.circuit_id` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.other_subscription.circuit_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/other_subscription/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--other_subscription--circuit_id) |
| `ingress_egress_gw_ar.hub.express_route_enabled.connections.weight` | [ingress_egress_gw_ar.hub.express_route_enabled.connections.weight](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/connections/#schema-ingress_egress_gw_ar--hub--express_route_enabled--connections--weight) |
| `ingress_egress_gw_ar.hub.express_route_enabled.custom_asn` | [ingress_egress_gw_ar.hub.express_route_enabled.custom_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/#schema-ingress_egress_gw_ar--hub--express_route_enabled--custom_asn) |
| `ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server` | [ingress_egress_gw_ar.hub.express_route_enabled.do_not_advertise_to_route_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/do_not_advertise_to_route_server/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/auto/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet/#schema-ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet/vnet_resource_group/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet_param/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param.ipv4` | [ingress_egress_gw_ar.hub.express_route_enabled.gateway_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/gateway_subnet/subnet_param/#schema-ingress_egress_gw_ar--hub--express_route_enabled--gateway_subnet--subnet_param--ipv4) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/route_server_subnet/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/route_server_subnet/auto/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/route_server_subnet/subnet/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/route_server_subnet/subnet/#schema-ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/route_server_subnet/subnet/vnet_resource_group/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/route_server_subnet/subnet_param/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param.ipv4` | [ingress_egress_gw_ar.hub.express_route_enabled.route_server_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/route_server_subnet/subnet_param/#schema-ingress_egress_gw_ar--hub--express_route_enabled--route_server_subnet--subnet_param--ipv4) |
| `ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route` | [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/site_registration_over_express_route/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route.cloudlink_network_name` | [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_express_route.cloudlink_network_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/site_registration_over_express_route/#schema-ingress_egress_gw_ar--hub--express_route_enabled--site_registration_over_express_route--cloudlink_network_name) |
| `ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet` | [ingress_egress_gw_ar.hub.express_route_enabled.site_registration_over_internet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/site_registration_over_internet/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az` | [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw1az](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/sku_ergw1az/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az` | [ingress_egress_gw_ar.hub.express_route_enabled.sku_ergw2az](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/sku_ergw2az/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf` | [ingress_egress_gw_ar.hub.express_route_enabled.sku_high_perf](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/sku_high_perf/#section) |
| `ingress_egress_gw_ar.hub.express_route_enabled.sku_standard` | [ingress_egress_gw_ar.hub.express_route_enabled.sku_standard](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/express_route_enabled/sku_standard/#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets` | [ingress_egress_gw_ar.hub.spoke_vnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.auto` | [ingress_egress_gw_ar.hub.spoke_vnets.auto](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/auto/#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.labels` | [ingress_egress_gw_ar.hub.spoke_vnets.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/labels/#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.manual` | [ingress_egress_gw_ar.hub.spoke_vnets.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/manual/#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/vnet/#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet.f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/vnet/f5_orchestrated_routing/#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet.manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/vnet/manual_routing/#section) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet.resource_group` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet.resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/vnet/#schema-ingress_egress_gw_ar--hub--spoke_vnets--vnet--resource_group) |
| `ingress_egress_gw_ar.hub.spoke_vnets.vnet.vnet_name` | [ingress_egress_gw_ar.hub.spoke_vnets.vnet.vnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/hub/spoke_vnets/vnet/#schema-ingress_egress_gw_ar--hub--spoke_vnets--vnet--vnet_name) |
| `ingress_egress_gw_ar.inside_static_routes` | [ingress_egress_gw_ar.inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list` | [ingress_egress_gw_ar.inside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/nexthop/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw_ar.inside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw_ar.inside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/inside_static_routes/static_route_list/#schema-ingress_egress_gw_ar--inside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw_ar.no_dc_cluster_group` | [ingress_egress_gw_ar.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/no_dc_cluster_group/#section) |
| `ingress_egress_gw_ar.no_forward_proxy` | [ingress_egress_gw_ar.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/no_forward_proxy/#section) |
| `ingress_egress_gw_ar.no_global_network` | [ingress_egress_gw_ar.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/no_global_network/#section) |
| `ingress_egress_gw_ar.no_inside_static_routes` | [ingress_egress_gw_ar.no_inside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/no_inside_static_routes/#section) |
| `ingress_egress_gw_ar.no_network_policy` | [ingress_egress_gw_ar.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/no_network_policy/#section) |
| `ingress_egress_gw_ar.no_outside_static_routes` | [ingress_egress_gw_ar.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/no_outside_static_routes/#section) |
| `ingress_egress_gw_ar.node` | [ingress_egress_gw_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/#section) |
| `ingress_egress_gw_ar.node.fault_domain` | [ingress_egress_gw_ar.node.fault_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/#schema-ingress_egress_gw_ar--node--fault_domain) |
| `ingress_egress_gw_ar.node.inside_subnet` | [ingress_egress_gw_ar.node.inside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/#section) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet` | [ingress_egress_gw_ar.node.inside_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet/#section) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet.subnet_name` | [ingress_egress_gw_ar.node.inside_subnet.subnet.subnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet/#schema-ingress_egress_gw_ar--node--inside_subnet--subnet--subnet_name) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw_ar.node.inside_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet/#schema-ingress_egress_gw_ar--node--inside_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group` | [ingress_egress_gw_ar.node.inside_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet/vnet_resource_group/#section) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet_param` | [ingress_egress_gw_ar.node.inside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet_param/#section) |
| `ingress_egress_gw_ar.node.inside_subnet.subnet_param.ipv4` | [ingress_egress_gw_ar.node.inside_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/inside_subnet/subnet_param/#schema-ingress_egress_gw_ar--node--inside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw_ar.node.node_number` | [ingress_egress_gw_ar.node.node_number](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/#schema-ingress_egress_gw_ar--node--node_number) |
| `ingress_egress_gw_ar.node.outside_subnet` | [ingress_egress_gw_ar.node.outside_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/#section) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet` | [ingress_egress_gw_ar.node.outside_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/subnet/#section) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet.subnet_name` | [ingress_egress_gw_ar.node.outside_subnet.subnet.subnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/subnet/#schema-ingress_egress_gw_ar--node--outside_subnet--subnet--subnet_name) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet.subnet_resource_grp` | [ingress_egress_gw_ar.node.outside_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/subnet/#schema-ingress_egress_gw_ar--node--outside_subnet--subnet--subnet_resource_grp) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group` | [ingress_egress_gw_ar.node.outside_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/subnet/vnet_resource_group/#section) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet_param` | [ingress_egress_gw_ar.node.outside_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/subnet_param/#section) |
| `ingress_egress_gw_ar.node.outside_subnet.subnet_param.ipv4` | [ingress_egress_gw_ar.node.outside_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/outside_subnet/subnet_param/#schema-ingress_egress_gw_ar--node--outside_subnet--subnet_param--ipv4) |
| `ingress_egress_gw_ar.node.update_domain` | [ingress_egress_gw_ar.node.update_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/node/#schema-ingress_egress_gw_ar--node--update_domain) |
| `ingress_egress_gw_ar.not_hub` | [ingress_egress_gw_ar.not_hub](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/not_hub/#section) |
| `ingress_egress_gw_ar.outside_static_routes` | [ingress_egress_gw_ar.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list` | [ingress_egress_gw_ar.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.attrs` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `ingress_egress_gw_ar.outside_static_routes.static_route_list.simple_static_route` | [ingress_egress_gw_ar.outside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/outside_static_routes/static_route_list/#schema-ingress_egress_gw_ar--outside_static_routes--static_route_list--simple_static_route) |
| `ingress_egress_gw_ar.performance_enhancement_mode` | [ingress_egress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/#section) |
| `ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_egress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/#section) |
| `ingress_egress_gw_ar.sm_connection_public_ip` | [ingress_egress_gw_ar.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/sm_connection_public_ip/#section) |
| `ingress_egress_gw_ar.sm_connection_pvt_ip` | [ingress_egress_gw_ar.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/sm_connection_pvt_ip/#section) |
| `ingress_gw` | [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/#section) |
| `ingress_gw.accelerated_networking` | [ingress_gw.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/accelerated_networking/#section) |
| `ingress_gw.accelerated_networking.disable_spec` | [ingress_gw.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/accelerated_networking/disable_spec/#section) |
| `ingress_gw.accelerated_networking.enable` | [ingress_gw.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/accelerated_networking/enable/#section) |
| `ingress_gw.az_nodes` | [ingress_gw.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/#section) |
| `ingress_gw.az_nodes.azure_az` | [ingress_gw.az_nodes.azure_az](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/#schema-ingress_gw--az_nodes--azure_az) |
| `ingress_gw.az_nodes.local_subnet` | [ingress_gw.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/local_subnet/#section) |
| `ingress_gw.az_nodes.local_subnet.subnet` | [ingress_gw.az_nodes.local_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/local_subnet/subnet/#section) |
| `ingress_gw.az_nodes.local_subnet.subnet.subnet_name` | [ingress_gw.az_nodes.local_subnet.subnet.subnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/local_subnet/subnet/#schema-ingress_gw--az_nodes--local_subnet--subnet--subnet_name) |
| `ingress_gw.az_nodes.local_subnet.subnet.subnet_resource_grp` | [ingress_gw.az_nodes.local_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/local_subnet/subnet/#schema-ingress_gw--az_nodes--local_subnet--subnet--subnet_resource_grp) |
| `ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group` | [ingress_gw.az_nodes.local_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/local_subnet/subnet/vnet_resource_group/#section) |
| `ingress_gw.az_nodes.local_subnet.subnet_param` | [ingress_gw.az_nodes.local_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/local_subnet/subnet_param/#section) |
| `ingress_gw.az_nodes.local_subnet.subnet_param.ipv4` | [ingress_gw.az_nodes.local_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/az_nodes/local_subnet/subnet_param/#schema-ingress_gw--az_nodes--local_subnet--subnet_param--ipv4) |
| `ingress_gw.azure_certified_hw` | [ingress_gw.azure_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/#schema-ingress_gw--azure_certified_hw) |
| `ingress_gw.performance_enhancement_mode` | [ingress_gw.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/#section) |
| `ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/#section) |
| `ingress_gw_ar` | [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/#section) |
| `ingress_gw_ar.accelerated_networking` | [ingress_gw_ar.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/#section) |
| `ingress_gw_ar.accelerated_networking.disable_spec` | [ingress_gw_ar.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/disable_spec/#section) |
| `ingress_gw_ar.accelerated_networking.enable` | [ingress_gw_ar.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/accelerated_networking/enable/#section) |
| `ingress_gw_ar.azure_certified_hw` | [ingress_gw_ar.azure_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/#schema-ingress_gw_ar--azure_certified_hw) |
| `ingress_gw_ar.node` | [ingress_gw_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/#section) |
| `ingress_gw_ar.node.fault_domain` | [ingress_gw_ar.node.fault_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/#schema-ingress_gw_ar--node--fault_domain) |
| `ingress_gw_ar.node.local_subnet` | [ingress_gw_ar.node.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/#section) |
| `ingress_gw_ar.node.local_subnet.subnet` | [ingress_gw_ar.node.local_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet/#section) |
| `ingress_gw_ar.node.local_subnet.subnet.subnet_name` | [ingress_gw_ar.node.local_subnet.subnet.subnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet/#schema-ingress_gw_ar--node--local_subnet--subnet--subnet_name) |
| `ingress_gw_ar.node.local_subnet.subnet.subnet_resource_grp` | [ingress_gw_ar.node.local_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet/#schema-ingress_gw_ar--node--local_subnet--subnet--subnet_resource_grp) |
| `ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group` | [ingress_gw_ar.node.local_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet/vnet_resource_group/#section) |
| `ingress_gw_ar.node.local_subnet.subnet_param` | [ingress_gw_ar.node.local_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet_param/#section) |
| `ingress_gw_ar.node.local_subnet.subnet_param.ipv4` | [ingress_gw_ar.node.local_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/local_subnet/subnet_param/#schema-ingress_gw_ar--node--local_subnet--subnet_param--ipv4) |
| `ingress_gw_ar.node.node_number` | [ingress_gw_ar.node.node_number](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/#schema-ingress_gw_ar--node--node_number) |
| `ingress_gw_ar.node.update_domain` | [ingress_gw_ar.node.update_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/node/#schema-ingress_gw_ar--node--update_domain) |
| `ingress_gw_ar.performance_enhancement_mode` | [ingress_gw_ar.performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/#section) |
| `ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/#section) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/#section) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-labels) |
| `log_receiver` | [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/log_receiver/#section) |
| `log_receiver.name` | [log_receiver.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/log_receiver/#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/log_receiver/#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/log_receiver/#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/logs_streaming_disabled/#section) |
| `machine_type` | [machine_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-machine_type) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-namespace) |
| `no_worker_nodes` | [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/no_worker_nodes/#section) |
| `nodes_per_az` | [nodes_per_az](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-nodes_per_az) |
| `offline_survivability_mode` | [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/offline_survivability_mode/#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/offline_survivability_mode/enable_offline_survivability_mode/#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/offline_survivability_mode/no_offline_survivability_mode/#section) |
| `os` | [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/os/#section) |
| `os.default_os_version` | [os.default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/os/default_os_version/#section) |
| `os.operating_system_version` | [os.operating_system_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/os/#schema-os--operating_system_version) |
| `resource_group` | [resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-resource_group) |
| `ssh_key` | [ssh_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-ssh_key) |
| `sw` | [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/sw/#section) |
| `sw.default_sw_version` | [sw.default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/sw/default_sw_version/#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/sw/#schema-sw--volterra_software_version) |
| `tags` | [tags](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-tags) |
| `total_nodes` | [total_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/#schema-total_nodes) |
| `vnet` | [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/#section) |
| `vnet.existing_vnet` | [vnet.existing_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/#section) |
| `vnet.existing_vnet.f5_orchestrated_routing` | [vnet.existing_vnet.f5_orchestrated_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/f5_orchestrated_routing/#section) |
| `vnet.existing_vnet.manual_routing` | [vnet.existing_vnet.manual_routing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/manual_routing/#section) |
| `vnet.existing_vnet.resource_group` | [vnet.existing_vnet.resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/#schema-vnet--existing_vnet--resource_group) |
| `vnet.existing_vnet.vnet_name` | [vnet.existing_vnet.vnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/existing_vnet/#schema-vnet--existing_vnet--vnet_name) |
| `vnet.new_vnet` | [vnet.new_vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/new_vnet/#section) |
| `vnet.new_vnet.autogenerate` | [vnet.new_vnet.autogenerate](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/new_vnet/autogenerate/#section) |
| `vnet.new_vnet.name` | [vnet.new_vnet.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/new_vnet/#schema-vnet--new_vnet--name) |
| `vnet.new_vnet.primary_ipv4` | [vnet.new_vnet.primary_ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/new_vnet/#schema-vnet--new_vnet--primary_ipv4) |
| `voltstack_cluster` | [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/#section) |
| `voltstack_cluster.accelerated_networking` | [voltstack_cluster.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/accelerated_networking/#section) |
| `voltstack_cluster.accelerated_networking.disable_spec` | [voltstack_cluster.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/accelerated_networking/disable_spec/#section) |
| `voltstack_cluster.accelerated_networking.enable` | [voltstack_cluster.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/accelerated_networking/enable/#section) |
| `voltstack_cluster.active_enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/#section) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-voltstack_cluster--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `voltstack_cluster.active_forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_forward_proxy_policies/#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/#section) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--name) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_forward_proxy_policies/forward_proxy_policies/#schema-voltstack_cluster--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `voltstack_cluster.active_network_policies` | [voltstack_cluster.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/#section) |
| `voltstack_cluster.active_network_policies.network_policies` | [voltstack_cluster.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/network_policies/#section) |
| `voltstack_cluster.active_network_policies.network_policies.name` | [voltstack_cluster.active_network_policies.network_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/network_policies/#schema-voltstack_cluster--active_network_policies--network_policies--name) |
| `voltstack_cluster.active_network_policies.network_policies.namespace` | [voltstack_cluster.active_network_policies.network_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/network_policies/#schema-voltstack_cluster--active_network_policies--network_policies--namespace) |
| `voltstack_cluster.active_network_policies.network_policies.tenant` | [voltstack_cluster.active_network_policies.network_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/active_network_policies/network_policies/#schema-voltstack_cluster--active_network_policies--network_policies--tenant) |
| `voltstack_cluster.az_nodes` | [voltstack_cluster.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/#section) |
| `voltstack_cluster.az_nodes.azure_az` | [voltstack_cluster.az_nodes.azure_az](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/#schema-voltstack_cluster--az_nodes--azure_az) |
| `voltstack_cluster.az_nodes.local_subnet` | [voltstack_cluster.az_nodes.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet` | [voltstack_cluster.az_nodes.local_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet.subnet_name` | [voltstack_cluster.az_nodes.local_subnet.subnet.subnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/#schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_name) |
| `voltstack_cluster.az_nodes.local_subnet.subnet.subnet_resource_grp` | [voltstack_cluster.az_nodes.local_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/#schema-voltstack_cluster--az_nodes--local_subnet--subnet--subnet_resource_grp) |
| `voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group` | [voltstack_cluster.az_nodes.local_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet/vnet_resource_group/#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param` | [voltstack_cluster.az_nodes.local_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/#section) |
| `voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4` | [voltstack_cluster.az_nodes.local_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/az_nodes/local_subnet/subnet_param/#schema-voltstack_cluster--az_nodes--local_subnet--subnet_param--ipv4) |
| `voltstack_cluster.azure_certified_hw` | [voltstack_cluster.azure_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/#schema-voltstack_cluster--azure_certified_hw) |
| `voltstack_cluster.dc_cluster_group` | [voltstack_cluster.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/dc_cluster_group/#section) |
| `voltstack_cluster.dc_cluster_group.name` | [voltstack_cluster.dc_cluster_group.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/dc_cluster_group/#schema-voltstack_cluster--dc_cluster_group--name) |
| `voltstack_cluster.dc_cluster_group.namespace` | [voltstack_cluster.dc_cluster_group.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/dc_cluster_group/#schema-voltstack_cluster--dc_cluster_group--namespace) |
| `voltstack_cluster.dc_cluster_group.tenant` | [voltstack_cluster.dc_cluster_group.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/dc_cluster_group/#schema-voltstack_cluster--dc_cluster_group--tenant) |
| `voltstack_cluster.default_storage` | [voltstack_cluster.default_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/default_storage/#section) |
| `voltstack_cluster.forward_proxy_allow_all` | [voltstack_cluster.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/forward_proxy_allow_all/#section) |
| `voltstack_cluster.global_network_list` | [voltstack_cluster.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/#section) |
| `voltstack_cluster.global_network_list.global_network_connections` | [voltstack_cluster.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#section) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-voltstack_cluster--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `voltstack_cluster.k8s_cluster` | [voltstack_cluster.k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/k8s_cluster/#section) |
| `voltstack_cluster.k8s_cluster.name` | [voltstack_cluster.k8s_cluster.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/k8s_cluster/#schema-voltstack_cluster--k8s_cluster--name) |
| `voltstack_cluster.k8s_cluster.namespace` | [voltstack_cluster.k8s_cluster.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/k8s_cluster/#schema-voltstack_cluster--k8s_cluster--namespace) |
| `voltstack_cluster.k8s_cluster.tenant` | [voltstack_cluster.k8s_cluster.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/k8s_cluster/#schema-voltstack_cluster--k8s_cluster--tenant) |
| `voltstack_cluster.no_dc_cluster_group` | [voltstack_cluster.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/no_dc_cluster_group/#section) |
| `voltstack_cluster.no_forward_proxy` | [voltstack_cluster.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/no_forward_proxy/#section) |
| `voltstack_cluster.no_global_network` | [voltstack_cluster.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/no_global_network/#section) |
| `voltstack_cluster.no_k8s_cluster` | [voltstack_cluster.no_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/no_k8s_cluster/#section) |
| `voltstack_cluster.no_network_policy` | [voltstack_cluster.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/no_network_policy/#section) |
| `voltstack_cluster.no_outside_static_routes` | [voltstack_cluster.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/no_outside_static_routes/#section) |
| `voltstack_cluster.outside_static_routes` | [voltstack_cluster.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list` | [voltstack_cluster.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/nexthop/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-voltstack_cluster--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `voltstack_cluster.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster.outside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/outside_static_routes/static_route_list/#schema-voltstack_cluster--outside_static_routes--static_route_list--simple_static_route) |
| `voltstack_cluster.sm_connection_public_ip` | [voltstack_cluster.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/sm_connection_public_ip/#section) |
| `voltstack_cluster.sm_connection_pvt_ip` | [voltstack_cluster.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/sm_connection_pvt_ip/#section) |
| `voltstack_cluster.storage_class_list` | [voltstack_cluster.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/storage_class_list/#section) |
| `voltstack_cluster.storage_class_list.storage_classes` | [voltstack_cluster.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/storage_class_list/storage_classes/#section) |
| `voltstack_cluster.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster.storage_class_list.storage_classes.default_storage_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/storage_class_list/storage_classes/#schema-voltstack_cluster--storage_class_list--storage_classes--default_storage_class) |
| `voltstack_cluster.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster.storage_class_list.storage_classes.storage_class_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/storage_class_list/storage_classes/#schema-voltstack_cluster--storage_class_list--storage_classes--storage_class_name) |
| `voltstack_cluster_ar` | [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/#section) |
| `voltstack_cluster_ar.accelerated_networking` | [voltstack_cluster_ar.accelerated_networking](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/#section) |
| `voltstack_cluster_ar.accelerated_networking.disable_spec` | [voltstack_cluster_ar.accelerated_networking.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/disable_spec/#section) |
| `voltstack_cluster_ar.accelerated_networking.enable` | [voltstack_cluster_ar.accelerated_networking.enable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/accelerated_networking/enable/#section) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies` | [voltstack_cluster_ar.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/#section) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies` | [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/#section) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [voltstack_cluster_ar.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-voltstack_cluster_ar--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `voltstack_cluster_ar.active_forward_proxy_policies` | [voltstack_cluster_ar.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_forward_proxy_policies/#section) |
| `voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies` | [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_forward_proxy_policies/forward_proxy_policies/#section) |
| `voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.name` | [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_forward_proxy_policies/forward_proxy_policies/#schema-voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies--name) |
| `voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.namespace` | [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_forward_proxy_policies/forward_proxy_policies/#schema-voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.tenant` | [voltstack_cluster_ar.active_forward_proxy_policies.forward_proxy_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_forward_proxy_policies/forward_proxy_policies/#schema-voltstack_cluster_ar--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `voltstack_cluster_ar.active_network_policies` | [voltstack_cluster_ar.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_network_policies/#section) |
| `voltstack_cluster_ar.active_network_policies.network_policies` | [voltstack_cluster_ar.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_network_policies/network_policies/#section) |
| `voltstack_cluster_ar.active_network_policies.network_policies.name` | [voltstack_cluster_ar.active_network_policies.network_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_network_policies/network_policies/#schema-voltstack_cluster_ar--active_network_policies--network_policies--name) |
| `voltstack_cluster_ar.active_network_policies.network_policies.namespace` | [voltstack_cluster_ar.active_network_policies.network_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_network_policies/network_policies/#schema-voltstack_cluster_ar--active_network_policies--network_policies--namespace) |
| `voltstack_cluster_ar.active_network_policies.network_policies.tenant` | [voltstack_cluster_ar.active_network_policies.network_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/active_network_policies/network_policies/#schema-voltstack_cluster_ar--active_network_policies--network_policies--tenant) |
| `voltstack_cluster_ar.azure_certified_hw` | [voltstack_cluster_ar.azure_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/#schema-voltstack_cluster_ar--azure_certified_hw) |
| `voltstack_cluster_ar.dc_cluster_group` | [voltstack_cluster_ar.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/dc_cluster_group/#section) |
| `voltstack_cluster_ar.dc_cluster_group.name` | [voltstack_cluster_ar.dc_cluster_group.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/dc_cluster_group/#schema-voltstack_cluster_ar--dc_cluster_group--name) |
| `voltstack_cluster_ar.dc_cluster_group.namespace` | [voltstack_cluster_ar.dc_cluster_group.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/dc_cluster_group/#schema-voltstack_cluster_ar--dc_cluster_group--namespace) |
| `voltstack_cluster_ar.dc_cluster_group.tenant` | [voltstack_cluster_ar.dc_cluster_group.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/dc_cluster_group/#schema-voltstack_cluster_ar--dc_cluster_group--tenant) |
| `voltstack_cluster_ar.default_storage` | [voltstack_cluster_ar.default_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/default_storage/#section) |
| `voltstack_cluster_ar.forward_proxy_allow_all` | [voltstack_cluster_ar.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/forward_proxy_allow_all/#section) |
| `voltstack_cluster_ar.global_network_list` | [voltstack_cluster_ar.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections` | [voltstack_cluster_ar.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/sli_to_global_dr/#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [voltstack_cluster_ar.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-voltstack_cluster_ar--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#section) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [voltstack_cluster_ar.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-voltstack_cluster_ar--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `voltstack_cluster_ar.k8s_cluster` | [voltstack_cluster_ar.k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/k8s_cluster/#section) |
| `voltstack_cluster_ar.k8s_cluster.name` | [voltstack_cluster_ar.k8s_cluster.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/k8s_cluster/#schema-voltstack_cluster_ar--k8s_cluster--name) |
| `voltstack_cluster_ar.k8s_cluster.namespace` | [voltstack_cluster_ar.k8s_cluster.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/k8s_cluster/#schema-voltstack_cluster_ar--k8s_cluster--namespace) |
| `voltstack_cluster_ar.k8s_cluster.tenant` | [voltstack_cluster_ar.k8s_cluster.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/k8s_cluster/#schema-voltstack_cluster_ar--k8s_cluster--tenant) |
| `voltstack_cluster_ar.no_dc_cluster_group` | [voltstack_cluster_ar.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_dc_cluster_group/#section) |
| `voltstack_cluster_ar.no_forward_proxy` | [voltstack_cluster_ar.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_forward_proxy/#section) |
| `voltstack_cluster_ar.no_global_network` | [voltstack_cluster_ar.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_global_network/#section) |
| `voltstack_cluster_ar.no_k8s_cluster` | [voltstack_cluster_ar.no_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_k8s_cluster/#section) |
| `voltstack_cluster_ar.no_network_policy` | [voltstack_cluster_ar.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_network_policy/#section) |
| `voltstack_cluster_ar.no_outside_static_routes` | [voltstack_cluster_ar.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/no_outside_static_routes/#section) |
| `voltstack_cluster_ar.node` | [voltstack_cluster_ar.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/#section) |
| `voltstack_cluster_ar.node.fault_domain` | [voltstack_cluster_ar.node.fault_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/#schema-voltstack_cluster_ar--node--fault_domain) |
| `voltstack_cluster_ar.node.local_subnet` | [voltstack_cluster_ar.node.local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/#section) |
| `voltstack_cluster_ar.node.local_subnet.subnet` | [voltstack_cluster_ar.node.local_subnet.subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet/#section) |
| `voltstack_cluster_ar.node.local_subnet.subnet.subnet_name` | [voltstack_cluster_ar.node.local_subnet.subnet.subnet_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet/#schema-voltstack_cluster_ar--node--local_subnet--subnet--subnet_name) |
| `voltstack_cluster_ar.node.local_subnet.subnet.subnet_resource_grp` | [voltstack_cluster_ar.node.local_subnet.subnet.subnet_resource_grp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet/#schema-voltstack_cluster_ar--node--local_subnet--subnet--subnet_resource_grp) |
| `voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group` | [voltstack_cluster_ar.node.local_subnet.subnet.vnet_resource_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet/vnet_resource_group/#section) |
| `voltstack_cluster_ar.node.local_subnet.subnet_param` | [voltstack_cluster_ar.node.local_subnet.subnet_param](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet_param/#section) |
| `voltstack_cluster_ar.node.local_subnet.subnet_param.ipv4` | [voltstack_cluster_ar.node.local_subnet.subnet_param.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/local_subnet/subnet_param/#schema-voltstack_cluster_ar--node--local_subnet--subnet_param--ipv4) |
| `voltstack_cluster_ar.node.node_number` | [voltstack_cluster_ar.node.node_number](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/#schema-voltstack_cluster_ar--node--node_number) |
| `voltstack_cluster_ar.node.update_domain` | [voltstack_cluster_ar.node.update_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/node/#schema-voltstack_cluster_ar--node--update_domain) |
| `voltstack_cluster_ar.outside_static_routes` | [voltstack_cluster_ar.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list` | [voltstack_cluster_ar.outside_static_routes.static_route_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.attrs` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--attrs) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/labels/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--kind) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--name) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--namespace) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--tenant) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/interface/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--interface--uid) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv4/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv4--addr) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/dual_stack/ipv6/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--dual_stack--ipv6--addr) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv4/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv4--addr) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/nexthop_address/ipv6/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--nexthop_address--ipv6--addr) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.type` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.nexthop.type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/nexthop/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--nexthop--type) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--plen) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv4/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv4--prefix) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#section) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--plen) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` | [voltstack_cluster_ar.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/custom_static_route/subnets/ipv6/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--custom_static_route--subnets--ipv6--prefix) |
| `voltstack_cluster_ar.outside_static_routes.static_route_list.simple_static_route` | [voltstack_cluster_ar.outside_static_routes.static_route_list.simple_static_route](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/outside_static_routes/static_route_list/#schema-voltstack_cluster_ar--outside_static_routes--static_route_list--simple_static_route) |
| `voltstack_cluster_ar.sm_connection_public_ip` | [voltstack_cluster_ar.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/sm_connection_public_ip/#section) |
| `voltstack_cluster_ar.sm_connection_pvt_ip` | [voltstack_cluster_ar.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/sm_connection_pvt_ip/#section) |
| `voltstack_cluster_ar.storage_class_list` | [voltstack_cluster_ar.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/#section) |
| `voltstack_cluster_ar.storage_class_list.storage_classes` | [voltstack_cluster_ar.storage_class_list.storage_classes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/storage_classes/#section) |
| `voltstack_cluster_ar.storage_class_list.storage_classes.default_storage_class` | [voltstack_cluster_ar.storage_class_list.storage_classes.default_storage_class](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/storage_classes/#schema-voltstack_cluster_ar--storage_class_list--storage_classes--default_storage_class) |
| `voltstack_cluster_ar.storage_class_list.storage_classes.storage_class_name` | [voltstack_cluster_ar.storage_class_list.storage_classes.storage_class_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/storage_class_list/storage_classes/#schema-voltstack_cluster_ar--storage_class_list--storage_classes--storage_class_name) |
| `waf_signatures` | [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/waf_signatures/#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/waf_signatures/automatic/#section) |
| `waf_signatures.manual` | [waf_signatures.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/waf_signatures/manual/#section) |

## Next pages

- [admin_password](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/admin_password/)
- [azure_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/azure_cred/)
- [block_all_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/block_all_services/)
- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/blocked_services/)
- [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/coordinates/)
- [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/custom_dns/)
- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/default_blocked_services/)
- [disable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/disable_encryption/)
- [enable_encryption](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/enable_encryption/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw/)
- [ingress_egress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_egress_gw_ar/)
- [ingress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/ingress_gw_ar/)
- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/kubernetes_upgrade_drain/)
- [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/log_receiver/)
- [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/logs_streaming_disabled/)
- [no_worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/no_worker_nodes/)
- [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/offline_survivability_mode/)
- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/os/)
- [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/sw/)
- [vnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/vnet/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster/)
- [voltstack_cluster_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/voltstack_cluster_ar/)
- [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/properties/waf_signatures/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/azure_vnet_site/)
