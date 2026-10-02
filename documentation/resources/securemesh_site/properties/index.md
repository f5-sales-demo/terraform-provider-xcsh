---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_securemesh_site."
xcsh_docs: {"aliases": ["securemesh site"], "body_bytes": 116466, "body_sha256": "sha256:50045307860e73be1c8e86c5d6a527c673dc01e22f1c34ca7e05a01066a8f442", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:blocked_services", "xcsh-docs:resources:securemesh_site:properties:bond_device_list", "xcsh-docs:resources:securemesh_site:properties:coordinates", "xcsh-docs:resources:securemesh_site:properties:custom_network_config", "xcsh-docs:resources:securemesh_site:properties:default_blocked_services", "xcsh-docs:resources:securemesh_site:properties:default_network_config", "xcsh-docs:resources:securemesh_site:properties:kubernetes_upgrade_drain", "xcsh-docs:resources:securemesh_site:properties:log_receiver", "xcsh-docs:resources:securemesh_site:properties:logs_streaming_disabled", "xcsh-docs:resources:securemesh_site:properties:master_node_configuration", "xcsh-docs:resources:securemesh_site:properties:no_bond_devices", "xcsh-docs:resources:securemesh_site:properties:offline_survivability_mode", "xcsh-docs:resources:securemesh_site:properties:os", "xcsh-docs:resources:securemesh_site:properties:performance_enhancement_mode", "xcsh-docs:resources:securemesh_site:properties:sw", "xcsh-docs:resources:securemesh_site:properties:timeouts", "xcsh-docs:resources:securemesh_site:properties:waf_signatures"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:reference", "parent_id": "xcsh-docs:resources:securemesh_site:fundamentals", "path": "documentation/resources/securemesh_site/properties/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0230321021201222-1111001300121312-3000300333211231-1100002333330020-1322112003322030-0300113022223332-3302330011112221-3302300032302122", "registry_path": "docs/guides/resources--securemesh_site--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["address"], "anchor": "schema-address", "description": "Site's geographical address that can be used to determine its latitude and longitude.", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["address"], "syntax": "attribute", "type": "string"}, {"aliases": ["annotations"], "anchor": "schema-annotations", "description": "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects.", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["annotations"], "syntax": "attribute", "type": "map"}, {"aliases": ["blocked services"], "anchor": "section", "description": "Disable node local services on this site. Note: The chosen services will GET disabled on all nodes in the site.", "document_id": "xcsh-docs:resources:securemesh_site:properties:blocked_services", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["blocked_services"], "syntax": "block", "type": "object"}, {"aliases": ["bond device list"], "anchor": "section", "description": "List of bond devices for this fleet.", "document_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bond_device_list:RequiredObjectAttributes:bond_devices", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:bond_device_list:bond_devices", "type": "requires"}], "schema_path": ["bond_device_list"], "syntax": "block", "type": "object"}, {"aliases": ["coordinates"], "anchor": "section", "description": "Coordinates of the site which provides the site physical location.", "document_id": "xcsh-docs:resources:securemesh_site:properties:coordinates", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["coordinates"], "syntax": "block", "type": "object"}, {"aliases": ["custom network config"], "anchor": "section", "description": "SmsNetworkConfiguration.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_enhanced_firewall_policies,active_network_policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:active_enhanced_firewall_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_enhanced_firewall_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:active_enhanced_firewall_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_forward_proxy_policies,forward_proxy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:active_forward_proxy_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_forward_proxy_policies,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:active_forward_proxy_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_enhanced_firewall_policies,active_network_policies", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:active_network_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_network_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:active_network_policies", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:default_config,slo_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:default_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:default_interface_config,interface_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:default_interface_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:default_sli_config,sli_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:default_sli_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_forward_proxy_policies,forward_proxy_allow_all", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:forward_proxy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:forward_proxy_allow_all,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:forward_proxy_allow_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:global_network_list,no_global_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:global_network_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:default_interface_config,interface_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:interface_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_forward_proxy_policies,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:no_forward_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:forward_proxy_allow_all,no_forward_proxy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:no_forward_proxy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:global_network_list,no_global_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:no_global_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_enhanced_firewall_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:no_network_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:active_network_policies,no_network_policy", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:no_network_policy", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:default_sli_config,sli_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sli_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:default_config,slo_config", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sm_connection_public_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config:ConflictingObjectAttributes:sm_connection_public_ip,sm_connection_pvt_ip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:sm_connection_pvt_ip", "type": "conflicts"}], "schema_path": ["custom_network_config"], "syntax": "block", "type": "object"}, {"aliases": ["default blocked services"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site:properties:default_blocked_services", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_blocked_services"], "syntax": "attribute", "type": "object"}, {"aliases": ["default network config"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site:properties:default_network_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["default_network_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["description"], "anchor": "schema-description", "description": "Human readable description for the object.", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["description"], "syntax": "attribute", "type": "string"}, {"aliases": ["disable"], "anchor": "schema-disable", "description": "A value of true will administratively disable the object.", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["disable"], "syntax": "attribute", "type": "bool"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Unique identifier for the resource.", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["kubernetes upgrade drain"], "anchor": "section", "description": "Specify how worker nodes within a site will be upgraded.", "document_id": "xcsh-docs:resources:securemesh_site:properties:kubernetes_upgrade_drain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain:ConflictingObjectAttributes:disable_upgrade_drain,enable_upgrade_drain", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain:ConflictingObjectAttributes:disable_upgrade_drain,enable_upgrade_drain", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "type": "conflicts"}], "schema_path": ["kubernetes_upgrade_drain"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "schema-labels", "description": "Map of string keys and values that can be used to organize and categorize (scope and select) objects as chosen by the user. Values specified here will be used by selector expression.", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["labels"], "syntax": "attribute", "type": "map"}, {"aliases": ["log receiver"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:securemesh_site:properties:log_receiver", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-log_receiver--name", "enforcement": "provider-schema", "group": "log_receiver:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:log_receiver", "type": "requires"}], "schema_path": ["log_receiver"], "syntax": "block", "type": "object"}, {"aliases": ["logs streaming disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site:properties:logs_streaming_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["logs_streaming_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["master node configuration"], "anchor": "section", "description": "Configuration of master nodes.", "document_id": "xcsh-docs:resources:securemesh_site:properties:master_node_configuration", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-master_node_configuration--name", "enforcement": "provider-schema", "group": "master_node_configuration:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:master_node_configuration", "type": "requires"}], "schema_path": ["master_node_configuration"], "syntax": "block", "type": "object"}, {"aliases": ["name"], "anchor": "schema-name", "description": "This is the name of configuration object. It has to be unique within the namespace. It can only be specified during create API and cannot be changed during replace API. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "This defines the workspace within which each the configuration object is to be created. Must be a DNS_LABEL format. For a namespace object itself, namespace value will be \"\"", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["no bond devices"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site:properties:no_bond_devices", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["no_bond_devices"], "syntax": "attribute", "type": "object"}, {"aliases": ["offline survivability mode"], "anchor": "section", "description": "Offline Survivability allows the Site to continue functioning normally without traffic loss during periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this feature is enabled, a site can continue to function as is with existing configuration for upto 7 days, even when the site is", "document_id": "xcsh-docs:resources:securemesh_site:properties:offline_survivability_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "offline_survivability_mode:ConflictingObjectAttributes:enable_offline_survivability_mode,no_offline_survivability_mode", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:offline_survivability_mode:enable_offline_survivability_mode", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "offline_survivability_mode:ConflictingObjectAttributes:enable_offline_survivability_mode,no_offline_survivability_mode", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:offline_survivability_mode:no_offline_survivability_mode", "type": "conflicts"}], "schema_path": ["offline_survivability_mode"], "syntax": "block", "type": "object"}, {"aliases": ["os"], "anchor": "section", "description": "Select the F5XC Operating System Version for the site. By default, latest available OS Version will be used. Refer to release notes to find required released OS versions.", "document_id": "xcsh-docs:resources:securemesh_site:properties:os", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-os--operating_system_version", "enforcement": "provider-schema", "group": "os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:os", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "os:ConflictingObjectAttributes:default_os_version,operating_system_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:os:default_os_version", "type": "conflicts"}], "schema_path": ["os"], "syntax": "block", "type": "object"}, {"aliases": ["performance enhancement mode"], "anchor": "section", "description": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "document_id": "xcsh-docs:resources:securemesh_site:properties:performance_enhancement_mode", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "type": "conflicts"}], "schema_path": ["performance_enhancement_mode"], "syntax": "block", "type": "object"}, {"aliases": ["sw"], "anchor": "section", "description": "Select the F5XC Software Version for the site. By default, latest available F5XC Software Version will be used. Refer to release notes to find required released SW versions.", "document_id": "xcsh-docs:resources:securemesh_site:properties:sw", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-sw--volterra_software_version", "enforcement": "provider-schema", "group": "sw:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:sw", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "sw:ConflictingObjectAttributes:default_sw_version,volterra_software_version", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:sw:default_sw_version", "type": "conflicts"}], "schema_path": ["sw"], "syntax": "block", "type": "object"}, {"aliases": ["duration", "operation timeout", "timeouts"], "anchor": "section", "description": "timeouts", "document_id": "xcsh-docs:resources:securemesh_site:properties:timeouts", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["timeouts"], "syntax": "block", "type": "object"}, {"aliases": ["volterra certified hw"], "anchor": "schema-volterra_certified_hw", "description": "Name for generic server certified hardware to form this Secure Mesh site.", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["volterra_certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf signatures"], "anchor": "section", "description": "Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied manually. Refer to release notes for details about available Signatures update modes.", "document_id": "xcsh-docs:resources:securemesh_site:properties:waf_signatures", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:waf_signatures:automatic", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "waf_signatures:ConflictingObjectAttributes:automatic,manual", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:waf_signatures:manual", "type": "conflicts"}], "schema_path": ["waf_signatures"], "syntax": "block", "type": "object"}, {"aliases": ["worker nodes"], "anchor": "schema-worker_nodes", "description": "Names of worker nodes.", "document_id": "xcsh-docs:resources:securemesh_site:reference", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["worker_nodes"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_securemesh_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- Property reference

## Direct properties

<a id="schema-address"></a>

### address property

Type: `"string"`. Optional, Computed.

Site's geographical address that can be used to determine its latitude and longitude.

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

<a id="schema-annotations"></a>

### annotations property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata.

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

- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/): complete subsection reference.

- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/): complete subsection reference.

- [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/coordinates/): complete subsection reference.

- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/): complete subsection reference.

- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/default_blocked_services/): complete subsection reference.

- [default_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/default_network_config/): complete subsection reference.

<a id="schema-description"></a>

### description property

Type: `"string"`. Optional.

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

<a id="schema-disable"></a>

### disable property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Upstream description:

A value of true will administratively disable the object.

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

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/): complete subsection reference.

<a id="schema-labels"></a>

### labels property

Type: `["map", "string"]`. Optional.

Labels is a user defined key value map that can be attached to resources for organization and
filtering.

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

- [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/log_receiver/): complete subsection reference.

- [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/logs_streaming_disabled/): complete subsection reference.

- [master_node_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/master_node_configuration/): complete subsection reference.

<a id="schema-name"></a>

### name property

Type: `"string"`. Required.

Name of the Securemesh Site. Must be unique within the namespace.

Upstream description:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NameValidator(),
}
```

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

Type: `"string"`. Required.

Namespace where the Securemesh Site is created.

Upstream description:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  validators.NamespaceValidator(),
}
```

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

- [no_bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/no_bond_devices/): complete subsection reference.

- [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/offline_survivability_mode/): complete subsection reference.

- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/os/): complete subsection reference.

- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/performance_enhancement_mode/): complete subsection reference.

- [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/sw/): complete subsection reference.

- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/timeouts/): complete subsection reference.

<a id="schema-volterra_certified_hw"></a>

### volterra_certified_hw property

Type: `"string"`. Required.

Name for generic server certified hardware to form this Secure Mesh site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

- [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/waf_signatures/): complete subsection reference.

<a id="schema-worker_nodes"></a>

### worker_nodes property

Type: `["list", "string"]`. Optional.

Worker Nodes. Names of worker nodes.

Upstream description:

Names of worker nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `address` | [address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-address) |
| `annotations` | [annotations](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-annotations) |
| `blocked_services` | [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/#section) |
| `blocked_services.blocked_service` | [blocked_services.blocked_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/blocked_service/#section) |
| `blocked_services.blocked_service.dns` | [blocked_services.blocked_service.dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/blocked_service/dns/#section) |
| `blocked_services.blocked_service.network_type` | [blocked_services.blocked_service.network_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/blocked_service/#schema-blocked_services--blocked_service--network_type) |
| `blocked_services.blocked_service.ssh` | [blocked_services.blocked_service.ssh](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/blocked_service/ssh/#section) |
| `blocked_services.blocked_service.web_user_interface` | [blocked_services.blocked_service.web_user_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/blocked_service/web_user_interface/#section) |
| `bond_device_list` | [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/#section) |
| `bond_device_list.bond_devices` | [bond_device_list.bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/#section) |
| `bond_device_list.bond_devices.active_backup` | [bond_device_list.bond_devices.active_backup](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/active_backup/#section) |
| `bond_device_list.bond_devices.devices` | [bond_device_list.bond_devices.devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/#schema-bond_device_list--bond_devices--devices) |
| `bond_device_list.bond_devices.lacp` | [bond_device_list.bond_devices.lacp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/lacp/#section) |
| `bond_device_list.bond_devices.lacp.rate` | [bond_device_list.bond_devices.lacp.rate](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/lacp/#schema-bond_device_list--bond_devices--lacp--rate) |
| `bond_device_list.bond_devices.link_polling_interval` | [bond_device_list.bond_devices.link_polling_interval](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/#schema-bond_device_list--bond_devices--link_polling_interval) |
| `bond_device_list.bond_devices.link_up_delay` | [bond_device_list.bond_devices.link_up_delay](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/#schema-bond_device_list--bond_devices--link_up_delay) |
| `bond_device_list.bond_devices.name` | [bond_device_list.bond_devices.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/bond_devices/#schema-bond_device_list--bond_devices--name) |
| `coordinates` | [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/coordinates/#section) |
| `coordinates.latitude` | [coordinates.latitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/coordinates/#schema-coordinates--latitude) |
| `coordinates.longitude` | [coordinates.longitude](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/coordinates/#schema-coordinates--longitude) |
| `custom_network_config` | [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/#section) |
| `custom_network_config.active_enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_enhanced_firewall_policies/#section) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_enhanced_firewall_policies/enhanced_firewall_policies/#section) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--name) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--namespace) |
| `custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` | [custom_network_config.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_enhanced_firewall_policies/enhanced_firewall_policies/#schema-custom_network_config--active_enhanced_firewall_policies--enhanced_firewall_policies--tenant) |
| `custom_network_config.active_forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_forward_proxy_policies/#section) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_forward_proxy_policies/forward_proxy_policies/#section) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_forward_proxy_policies/forward_proxy_policies/#schema-custom_network_config--active_forward_proxy_policies--forward_proxy_policies--name) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_forward_proxy_policies/forward_proxy_policies/#schema-custom_network_config--active_forward_proxy_policies--forward_proxy_policies--namespace) |
| `custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant` | [custom_network_config.active_forward_proxy_policies.forward_proxy_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_forward_proxy_policies/forward_proxy_policies/#schema-custom_network_config--active_forward_proxy_policies--forward_proxy_policies--tenant) |
| `custom_network_config.active_network_policies` | [custom_network_config.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_network_policies/#section) |
| `custom_network_config.active_network_policies.network_policies` | [custom_network_config.active_network_policies.network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_network_policies/network_policies/#section) |
| `custom_network_config.active_network_policies.network_policies.name` | [custom_network_config.active_network_policies.network_policies.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_network_policies/network_policies/#schema-custom_network_config--active_network_policies--network_policies--name) |
| `custom_network_config.active_network_policies.network_policies.namespace` | [custom_network_config.active_network_policies.network_policies.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_network_policies/network_policies/#schema-custom_network_config--active_network_policies--network_policies--namespace) |
| `custom_network_config.active_network_policies.network_policies.tenant` | [custom_network_config.active_network_policies.network_policies.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/active_network_policies/network_policies/#schema-custom_network_config--active_network_policies--network_policies--tenant) |
| `custom_network_config.default_config` | [custom_network_config.default_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/default_config/#section) |
| `custom_network_config.default_interface_config` | [custom_network_config.default_interface_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/default_interface_config/#section) |
| `custom_network_config.default_sli_config` | [custom_network_config.default_sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/default_sli_config/#section) |
| `custom_network_config.forward_proxy_allow_all` | [custom_network_config.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/forward_proxy_allow_all/#section) |
| `custom_network_config.global_network_list` | [custom_network_config.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/#section) |
| `custom_network_config.global_network_list.global_network_connections` | [custom_network_config.global_network_list.global_network_connections](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/#section) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/#section) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#section) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--name) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--namespace) |
| `custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/sli_to_global_dr/global_vn/#schema-custom_network_config--global_network_list--global_network_connections--sli_to_global_dr--global_vn--tenant) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/#section) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#section) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--name) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--namespace) |
| `custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` | [custom_network_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/global_network_list/global_network_connections/slo_to_global_dr/global_vn/#schema-custom_network_config--global_network_list--global_network_connections--slo_to_global_dr--global_vn--tenant) |
| `custom_network_config.interface_list` | [custom_network_config.interface_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/#section) |
| `custom_network_config.interface_list.interfaces` | [custom_network_config.interface_list.interfaces](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/#section) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_disabled/#section) |
| `custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled` | [custom_network_config.interface_list.interfaces.dc_cluster_group_connectivity_interface_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dc_cluster_group_connectivity_interface_enabled/#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface` | [custom_network_config.interface_list.interfaces.dedicated_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/cluster/#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_interface.device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/#schema-custom_network_config--interface_list--interfaces--dedicated_interface--device) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.is_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/is_primary/#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/monitor/#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.dedicated_interface.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/monitor_disabled/#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_interface.mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/#schema-custom_network_config--interface_list--interfaces--dedicated_interface--mtu) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_interface.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/#schema-custom_network_config--interface_list--interfaces--dedicated_interface--node) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.not_primary` | [custom_network_config.interface_list.interfaces.dedicated_interface.not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/not_primary/#section) |
| `custom_network_config.interface_list.interfaces.dedicated_interface.priority` | [custom_network_config.interface_list.interfaces.dedicated_interface.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_interface/#schema-custom_network_config--interface_list--interfaces--dedicated_interface--priority) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface` | [custom_network_config.interface_list.interfaces.dedicated_management_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/#section) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/cluster/#section) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.device` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/#schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--device) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/#schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--mtu) |
| `custom_network_config.interface_list.interfaces.dedicated_management_interface.node` | [custom_network_config.interface_list.interfaces.dedicated_management_interface.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/dedicated_management_interface/#schema-custom_network_config--interface_list--interfaces--dedicated_management_interface--node) |
| `custom_network_config.interface_list.interfaces.description_spec` | [custom_network_config.interface_list.interfaces.description_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/#schema-custom_network_config--interface_list--interfaces--description_spec) |
| `custom_network_config.interface_list.interfaces.ethernet_interface` | [custom_network_config.interface_list.interfaces.ethernet_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.cluster` | [custom_network_config.interface_list.interfaces.ethernet_interface.cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/cluster/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.device` | [custom_network_config.interface_list.interfaces.ethernet_interface.device](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--device) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_client](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_client/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/automatic_from_end/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/automatic_from_start/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dgw_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dgw_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.dns_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--dns_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/first_address/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/last_address/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.network_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--network_prefix) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pool_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pool_settings) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/pools/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.end_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/pools/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools--end_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.exclude](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/pools/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools--exclude) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.pools.start_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/pools/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_networks--pools--start_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_networks.same_as_dgw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/dhcp_networks/same_as_dgw/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.dhcp_option82_tag](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--dhcp_option82_tag) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.fixed_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--fixed_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/interface_ip_map/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.dhcp_server.interface_ip_map.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/dhcp_server/interface_ip_map/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--dhcp_server--interface_ip_map--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/host/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/configured_list/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.configured_list.dns_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/configured_list/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--configured_list--dns_list) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.configured_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--dns_config--local_dns--configured_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/first_address/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/dns_config/local_dns/last_address/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.network_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--network_prefix) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_end](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/automatic_from_end/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.automatic_from_start](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/automatic_from_start/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--network_prefix) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pool_settings) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/pools/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/pools/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--end_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/dhcp_networks/pools/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--dhcp_networks--pools--start_ip) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.fixed_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--fixed_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/interface_ip_map/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/ipv6_auto_config/router/stateful/interface_ip_map/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--ipv6_auto_config--router--stateful--interface_ip_map--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.is_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.is_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/is_primary/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/monitor/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled` | [custom_network_config.interface_list.interfaces.ethernet_interface.monitor_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/monitor_disabled/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.mtu` | [custom_network_config.interface_list.interfaces.ethernet_interface.mtu](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--mtu) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.no_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/no_ipv6_address/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.node` | [custom_network_config.interface_list.interfaces.ethernet_interface.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--node) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.not_primary` | [custom_network_config.interface_list.interfaces.ethernet_interface.not_primary](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/not_primary/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.priority` | [custom_network_config.interface_list.interfaces.ethernet_interface.priority](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--priority) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_inside_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/site_local_inside_network/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.site_local_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/site_local_network/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/cluster_static_ip/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.cluster_static_ip.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/cluster_static_ip/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--cluster_static_ip--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/node_static_ip/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.default_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/node_static_ip/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--default_gw) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.dns_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/node_static_ip/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--dns_server) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ip.node_static_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ip/node_static_ip/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ip--node_static_ip--ip_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/cluster_static_ip/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.cluster_static_ip.interface_ip_map](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/cluster_static_ip/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--cluster_static_ip--interface_ip_map) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/node_static_ip/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.default_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/node_static_ip/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip--default_gw) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.dns_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/node_static_ip/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip--dns_server) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address` | [custom_network_config.interface_list.interfaces.ethernet_interface.static_ipv6_address.node_static_ip.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/static_ipv6_address/node_static_ip/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--static_ipv6_address--node_static_ip--ip_address) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.storage_network` | [custom_network_config.interface_list.interfaces.ethernet_interface.storage_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/storage_network/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.untagged` | [custom_network_config.interface_list.interfaces.ethernet_interface.untagged](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/untagged/#section) |
| `custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id` | [custom_network_config.interface_list.interfaces.ethernet_interface.vlan_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/ethernet_interface/#schema-custom_network_config--interface_list--interfaces--ethernet_interface--vlan_id) |
| `custom_network_config.interface_list.interfaces.labels` | [custom_network_config.interface_list.interfaces.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/interface_list/interfaces/#schema-custom_network_config--interface_list--interfaces--labels) |
| `custom_network_config.no_forward_proxy` | [custom_network_config.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/no_forward_proxy/#section) |
| `custom_network_config.no_global_network` | [custom_network_config.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/no_global_network/#section) |
| `custom_network_config.no_network_policy` | [custom_network_config.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/no_network_policy/#section) |
| `custom_network_config.sli_config` | [custom_network_config.sli_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/#section) |
| `custom_network_config.sli_config.dc_cluster_group` | [custom_network_config.sli_config.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/dc_cluster_group/#section) |
| `custom_network_config.sli_config.dc_cluster_group.name` | [custom_network_config.sli_config.dc_cluster_group.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/dc_cluster_group/#schema-custom_network_config--sli_config--dc_cluster_group--name) |
| `custom_network_config.sli_config.dc_cluster_group.namespace` | [custom_network_config.sli_config.dc_cluster_group.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/dc_cluster_group/#schema-custom_network_config--sli_config--dc_cluster_group--namespace) |
| `custom_network_config.sli_config.dc_cluster_group.tenant` | [custom_network_config.sli_config.dc_cluster_group.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/dc_cluster_group/#schema-custom_network_config--sli_config--dc_cluster_group--tenant) |
| `custom_network_config.sli_config.labels` | [custom_network_config.sli_config.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/#schema-custom_network_config--sli_config--labels) |
| `custom_network_config.sli_config.nameserver` | [custom_network_config.sli_config.nameserver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/#schema-custom_network_config--sli_config--nameserver) |
| `custom_network_config.sli_config.no_dc_cluster_group` | [custom_network_config.sli_config.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/no_dc_cluster_group/#section) |
| `custom_network_config.sli_config.no_static_routes` | [custom_network_config.sli_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/no_static_routes/#section) |
| `custom_network_config.sli_config.no_v6_static_routes` | [custom_network_config.sli_config.no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/no_v6_static_routes/#section) |
| `custom_network_config.sli_config.static_routes` | [custom_network_config.sli_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/#section) |
| `custom_network_config.sli_config.static_routes.static_routes` | [custom_network_config.sli_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/#section) |
| `custom_network_config.sli_config.static_routes.static_routes.attrs` | [custom_network_config.sli_config.static_routes.static_routes.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/#schema-custom_network_config--sli_config--static_routes--static_routes--attrs) |
| `custom_network_config.sli_config.static_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_routes.static_routes.default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/default_gateway/#section) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_routes.static_routes.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/#schema-custom_network_config--sli_config--static_routes--static_routes--ip_address) |
| `custom_network_config.sli_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_routes.static_routes.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/#schema-custom_network_config--sli_config--static_routes--static_routes--ip_prefixes) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/#section) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/#section) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/interface/#section) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_routes/static_routes/node_interface/list/#schema-custom_network_config--sli_config--static_routes--static_routes--node_interface--list--node) |
| `custom_network_config.sli_config.static_v6_routes` | [custom_network_config.sli_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes` | [custom_network_config.sli_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.attrs` | [custom_network_config.sli_config.static_v6_routes.static_routes.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/#schema-custom_network_config--sli_config--static_v6_routes--static_routes--attrs) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/default_gateway/#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/#schema-custom_network_config--sli_config--static_v6_routes--static_routes--ip_address) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.sli_config.static_v6_routes.static_routes.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/#schema-custom_network_config--sli_config--static_v6_routes--static_routes--ip_prefixes) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/interface/#section) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/static_v6_routes/static_routes/node_interface/list/#schema-custom_network_config--sli_config--static_v6_routes--static_routes--node_interface--list--node) |
| `custom_network_config.sli_config.vip` | [custom_network_config.sli_config.vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sli_config/#schema-custom_network_config--sli_config--vip) |
| `custom_network_config.slo_config` | [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/#section) |
| `custom_network_config.slo_config.dc_cluster_group` | [custom_network_config.slo_config.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/dc_cluster_group/#section) |
| `custom_network_config.slo_config.dc_cluster_group.name` | [custom_network_config.slo_config.dc_cluster_group.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/dc_cluster_group/#schema-custom_network_config--slo_config--dc_cluster_group--name) |
| `custom_network_config.slo_config.dc_cluster_group.namespace` | [custom_network_config.slo_config.dc_cluster_group.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/dc_cluster_group/#schema-custom_network_config--slo_config--dc_cluster_group--namespace) |
| `custom_network_config.slo_config.dc_cluster_group.tenant` | [custom_network_config.slo_config.dc_cluster_group.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/dc_cluster_group/#schema-custom_network_config--slo_config--dc_cluster_group--tenant) |
| `custom_network_config.slo_config.labels` | [custom_network_config.slo_config.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/#schema-custom_network_config--slo_config--labels) |
| `custom_network_config.slo_config.nameserver` | [custom_network_config.slo_config.nameserver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/#schema-custom_network_config--slo_config--nameserver) |
| `custom_network_config.slo_config.no_dc_cluster_group` | [custom_network_config.slo_config.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/no_dc_cluster_group/#section) |
| `custom_network_config.slo_config.no_static_routes` | [custom_network_config.slo_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/no_static_routes/#section) |
| `custom_network_config.slo_config.no_v6_static_routes` | [custom_network_config.slo_config.no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/no_v6_static_routes/#section) |
| `custom_network_config.slo_config.static_routes` | [custom_network_config.slo_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/#section) |
| `custom_network_config.slo_config.static_routes.static_routes` | [custom_network_config.slo_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/#section) |
| `custom_network_config.slo_config.static_routes.static_routes.attrs` | [custom_network_config.slo_config.static_routes.static_routes.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/#schema-custom_network_config--slo_config--static_routes--static_routes--attrs) |
| `custom_network_config.slo_config.static_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_routes.static_routes.default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/default_gateway/#section) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_routes.static_routes.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/#schema-custom_network_config--slo_config--static_routes--static_routes--ip_address) |
| `custom_network_config.slo_config.static_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_routes.static_routes.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/#schema-custom_network_config--slo_config--static_routes--static_routes--ip_prefixes) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/#section) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/#section) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/interface/#section) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/node_interface/list/#schema-custom_network_config--slo_config--static_routes--static_routes--node_interface--list--node) |
| `custom_network_config.slo_config.static_v6_routes` | [custom_network_config.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes` | [custom_network_config.slo_config.static_v6_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.attrs` | [custom_network_config.slo_config.static_v6_routes.static_routes.attrs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/#schema-custom_network_config--slo_config--static_v6_routes--static_routes--attrs) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway` | [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/default_gateway/#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_address` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_address](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/#schema-custom_network_config--slo_config--static_v6_routes--static_routes--ip_address) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes` | [custom_network_config.slo_config.static_v6_routes.static_routes.ip_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/#schema-custom_network_config--slo_config--static_v6_routes--static_routes--ip_prefixes) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/interface/#section) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.kind](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--kind) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--name) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--namespace) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--tenant) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface.uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/interface/#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--interface--uid) |
| `custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node` | [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.node](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_v6_routes/static_routes/node_interface/list/#schema-custom_network_config--slo_config--static_v6_routes--static_routes--node_interface--list--node) |
| `custom_network_config.slo_config.vip` | [custom_network_config.slo_config.vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/#schema-custom_network_config--slo_config--vip) |
| `custom_network_config.sm_connection_public_ip` | [custom_network_config.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sm_connection_public_ip/#section) |
| `custom_network_config.sm_connection_pvt_ip` | [custom_network_config.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/sm_connection_pvt_ip/#section) |
| `custom_network_config.tunnel_dead_timeout` | [custom_network_config.tunnel_dead_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/#schema-custom_network_config--tunnel_dead_timeout) |
| `custom_network_config.vip_vrrp_mode` | [custom_network_config.vip_vrrp_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/#schema-custom_network_config--vip_vrrp_mode) |
| `default_blocked_services` | [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/default_blocked_services/#section) |
| `default_network_config` | [default_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/default_network_config/#section) |
| `description` | [description](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-description) |
| `disable` | [disable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-disable) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-id) |
| `kubernetes_upgrade_drain` | [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.disable_upgrade_drain` | [kubernetes_upgrade_drain.disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain` | [kubernetes_upgrade_drain.enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/#section) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` | [kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/#schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout) |
| `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` | [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/#section) |
| `labels` | [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-labels) |
| `log_receiver` | [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/log_receiver/#section) |
| `log_receiver.name` | [log_receiver.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/log_receiver/#schema-log_receiver--name) |
| `log_receiver.namespace` | [log_receiver.namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/log_receiver/#schema-log_receiver--namespace) |
| `log_receiver.tenant` | [log_receiver.tenant](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/log_receiver/#schema-log_receiver--tenant) |
| `logs_streaming_disabled` | [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/logs_streaming_disabled/#section) |
| `master_node_configuration` | [master_node_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/master_node_configuration/#section) |
| `master_node_configuration.name` | [master_node_configuration.name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/master_node_configuration/#schema-master_node_configuration--name) |
| `master_node_configuration.public_ip` | [master_node_configuration.public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/master_node_configuration/#schema-master_node_configuration--public_ip) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-namespace) |
| `no_bond_devices` | [no_bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/no_bond_devices/#section) |
| `offline_survivability_mode` | [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/offline_survivability_mode/#section) |
| `offline_survivability_mode.enable_offline_survivability_mode` | [offline_survivability_mode.enable_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/offline_survivability_mode/enable_offline_survivability_mode/#section) |
| `offline_survivability_mode.no_offline_survivability_mode` | [offline_survivability_mode.no_offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/offline_survivability_mode/no_offline_survivability_mode/#section) |
| `os` | [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/os/#section) |
| `os.default_os_version` | [os.default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/os/default_os_version/#section) |
| `os.operating_system_version` | [os.operating_system_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/os/#schema-os--operating_system_version) |
| `performance_enhancement_mode` | [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/performance_enhancement_mode/#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced` | [performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/jumbo/#section) |
| `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` | [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced` | [performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/#section) |
| `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` | [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/#section) |
| `sw` | [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/sw/#section) |
| `sw.default_sw_version` | [sw.default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/sw/default_sw_version/#section) |
| `sw.volterra_software_version` | [sw.volterra_software_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/sw/#schema-sw--volterra_software_version) |
| `timeouts` | [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/timeouts/#section) |
| `timeouts.create` | [timeouts.create](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/timeouts/#schema-timeouts--create) |
| `timeouts.delete` | [timeouts.delete](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/timeouts/#schema-timeouts--delete) |
| `timeouts.read` | [timeouts.read](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/timeouts/#schema-timeouts--read) |
| `timeouts.update` | [timeouts.update](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/timeouts/#schema-timeouts--update) |
| `volterra_certified_hw` | [volterra_certified_hw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-volterra_certified_hw) |
| `waf_signatures` | [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/waf_signatures/#section) |
| `waf_signatures.automatic` | [waf_signatures.automatic](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/waf_signatures/automatic/#section) |
| `waf_signatures.manual` | [waf_signatures.manual](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/waf_signatures/manual/#section) |
| `worker_nodes` | [worker_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/#schema-worker_nodes) |

## Next pages

- [blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/blocked_services/)
- [bond_device_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/bond_device_list/)
- [coordinates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/coordinates/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [default_blocked_services](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/default_blocked_services/)
- [default_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/default_network_config/)
- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/kubernetes_upgrade_drain/)
- [log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/log_receiver/)
- [logs_streaming_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/logs_streaming_disabled/)
- [master_node_configuration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/master_node_configuration/)
- [no_bond_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/no_bond_devices/)
- [offline_survivability_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/offline_survivability_mode/)
- [os](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/os/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/performance_enhancement_mode/)
- [sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/sw/)
- [timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/timeouts/)
- [waf_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/waf_signatures/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
