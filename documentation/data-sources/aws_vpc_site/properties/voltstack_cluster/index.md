---
page_title: "voltstack_cluster"
subcategory: "Infrastructure"
description: "App Stack cluster of single interface AWS nodes."
xcsh_docs: {"aliases": ["voltstack cluster"], "body_bytes": 10322, "body_sha256": "sha256:ac447bd3b7443c6cb5501a1962672b418c82f156813c57940990f0f44fb71068", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_network_policies", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:az_nodes", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:dc_cluster_group", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:default_storage", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:forward_proxy_allow_all", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:global_network_list", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:k8s_cluster", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_dc_cluster_group", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_forward_proxy", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_global_network", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_k8s_cluster", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_network_policy", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_outside_static_routes", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:sm_connection_public_ip", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:sm_connection_pvt_ip", "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:storage_class_list"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:reference", "path": "documentation/data-sources/aws_vpc_site/properties/voltstack_cluster/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3032000321201331-3010332213001100-2033203210030223-0033302023003012-1131213122303130-0100122111311301-3321301033013012-0131000331232321", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster"], "schema_version": 1, "sections": [{"aliases": ["active enhanced firewall policies"], "anchor": "section", "description": "List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS available under firewall policies with an additional option for service insertion.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_enhanced_firewall_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "active_enhanced_firewall_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["active forward proxy policies"], "anchor": "section", "description": "Ordered List of Forward Proxy Policies active.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_forward_proxy_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "active_forward_proxy_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["active network policies"], "anchor": "section", "description": "List of firewall policy views.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:active_network_policies", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "active_network_policies"], "syntax": "attribute", "type": "object"}, {"aliases": ["allowed vip port"], "anchor": "section", "description": "This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:allowed_vip_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "allowed_vip_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws certified hw"], "anchor": "schema-voltstack_cluster--aws_certified_hw", "description": "Name for AWS certified hardware.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "aws_certified_hw"], "syntax": "attribute", "type": "string"}, {"aliases": ["az nodes"], "anchor": "section", "description": "Only Single AZ or Three AZ(s) nodes are supported currently.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:az_nodes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["voltstack_cluster", "az_nodes"], "syntax": "attribute", "type": "object"}, {"aliases": ["dc cluster group"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:dc_cluster_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["default storage"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:default_storage", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "default_storage"], "syntax": "attribute", "type": "object"}, {"aliases": ["forward proxy allow all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:forward_proxy_allow_all", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "forward_proxy_allow_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["global network list"], "anchor": "section", "description": "List of global network connections.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:global_network_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "global_network_list"], "syntax": "attribute", "type": "object"}, {"aliases": ["k8s cluster"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:k8s_cluster", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "k8s_cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["no dc cluster group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_dc_cluster_group", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["no forward proxy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_forward_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_forward_proxy"], "syntax": "attribute", "type": "object"}, {"aliases": ["no global network"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_global_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_global_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["no k8s cluster"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_k8s_cluster", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_k8s_cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["no network policy"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_network_policy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_network_policy"], "syntax": "attribute", "type": "object"}, {"aliases": ["no outside static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:no_outside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "no_outside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["outside static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:outside_static_routes", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "outside_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["sm connection public ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:sm_connection_public_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "sm_connection_public_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["sm connection pvt ip"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:sm_connection_pvt_ip", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "sm_connection_pvt_ip"], "syntax": "attribute", "type": "object"}, {"aliases": ["storage class list"], "anchor": "section", "description": "Add additional custom storage classes in Kubernetes for this site.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:voltstack_cluster:storage_class_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["voltstack_cluster", "storage_class_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/voltstack_cluster/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "App Stack cluster of single interface AWS nodes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- voltstack_cluster

<a id="section"></a>

Type: `"single"`. Computed.

App Stack cluster of single interface AWS nodes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

## Direct properties

- [active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/): complete subsection reference.

- [active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/): complete subsection reference.

- [active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_network_policies/): complete subsection reference.

- [allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/): complete subsection reference.

<a id="schema-voltstack_cluster--aws_certified_hw"></a>

### aws_certified_hw property

Type: `"string"`. Computed.

\[Enum: aws-byol-voltstack-combo\] AWS Certified Hardware. Name for AWS certified hardware. The only
possible value is \`aws-byol-voltstack-combo\`.

Upstream description:

Name for AWS certified hardware.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-voltstack-combo"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/): complete subsection reference.

- [dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/dc_cluster_group/): complete subsection reference.

- [default_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/default_storage/): complete subsection reference.

- [forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/forward_proxy_allow_all/): complete subsection reference.

- [global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/): complete subsection reference.

- [k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/k8s_cluster/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_dc_cluster_group/): complete subsection reference.

- [no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_forward_proxy/): complete subsection reference.

- [no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_global_network/): complete subsection reference.

- [no_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_k8s_cluster/): complete subsection reference.

- [no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_network_policy/): complete subsection reference.

- [no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_outside_static_routes/): complete subsection reference.

- [outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/): complete subsection reference.

- [sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/sm_connection_public_ip/): complete subsection reference.

- [sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/sm_connection_pvt_ip/): complete subsection reference.

- [storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/storage_class_list/): complete subsection reference.

## Next pages

- [voltstack_cluster.active_enhanced_firewall_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_enhanced_firewall_policies/)
- [voltstack_cluster.active_forward_proxy_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_forward_proxy_policies/)
- [voltstack_cluster.active_network_policies](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/active_network_policies/)
- [voltstack_cluster.allowed_vip_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/allowed_vip_port/)
- [voltstack_cluster.az_nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/az_nodes/)
- [voltstack_cluster.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/dc_cluster_group/)
- [voltstack_cluster.default_storage](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/default_storage/)
- [voltstack_cluster.forward_proxy_allow_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/forward_proxy_allow_all/)
- [voltstack_cluster.global_network_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/global_network_list/)
- [voltstack_cluster.k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/k8s_cluster/)
- [voltstack_cluster.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_dc_cluster_group/)
- [voltstack_cluster.no_forward_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_forward_proxy/)
- [voltstack_cluster.no_global_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_global_network/)
- [voltstack_cluster.no_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_k8s_cluster/)
- [voltstack_cluster.no_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_network_policy/)
- [voltstack_cluster.no_outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/no_outside_static_routes/)
- [voltstack_cluster.outside_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/outside_static_routes/)
- [voltstack_cluster.sm_connection_public_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/sm_connection_public_ip/)
- [voltstack_cluster.sm_connection_pvt_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/sm_connection_pvt_ip/)
- [voltstack_cluster.storage_class_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/voltstack_cluster/storage_class_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
