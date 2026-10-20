---
page_title: "f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param"
subcategory: ""
description: "Parameters for creating a new cloud subnet."
xcsh_docs: {"aliases": ["f5 big ip aws service nodes mgmt subnet subnet param"], "body_bytes": 2226, "body_sha256": "sha256:90be247ba8bf2c08b00d595741ac58c00fcf6913588f8eef69abb64ac39edfc1", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet:subnet_param", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet", "path": "documentation/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/mgmt_subnet/subnet_param/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2223120103320113-0332301010330122-0103122200032002-1121221101303201-0023220323101012-1221220220222321-3133210302000013-3212222003231121", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["f5_big_ip_aws_service", "nodes", "mgmt_subnet", "subnet_param"], "schema_version": 1, "sections": [{"aliases": ["f5 big ip aws service nodes mgmt subnet subnet param ipv4"], "anchor": "schema-f5_big_ip_aws_service--nodes--mgmt_subnet--subnet_param--ipv4", "description": "IPv4 subnet prefix for this subnet.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:f5_big_ip_aws_service:nodes:mgmt_subnet:subnet_param", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["f5_big_ip_aws_service", "nodes", "mgmt_subnet", "subnet_param", "ipv4"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/mgmt_subnet/subnet_param/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Parameters for creating a new cloud subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [f5_big_ip_aws_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/)
- [f5_big_ip_aws_service.nodes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/)
- [f5_big_ip_aws_service.nodes.mgmt_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/f5_big_ip_aws_service/nodes/mgmt_subnet/)
- f5_big_ip_aws_service.nodes.mgmt_subnet.subnet_param

<a id="section"></a>

Type: `"single"`. Computed.

Parameters for creating a new cloud subnet.

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

## Direct properties

<a id="schema-f5_big_ip_aws_service--nodes--mgmt_subnet--subnet_param--ipv4"></a>

### ipv4 property

Type: `"string"`. Computed.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-20T00:10:01+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```
