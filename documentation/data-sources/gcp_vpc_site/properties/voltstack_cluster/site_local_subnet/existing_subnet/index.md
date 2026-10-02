---
page_title: "voltstack_cluster.site_local_subnet.existing_subnet"
subcategory: "Infrastructure"
description: "Name of existing GCP subnet."
xcsh_docs: {"aliases": ["voltstack cluster site local subnet existing subnet"], "body_bytes": 2505, "body_sha256": "sha256:af7d1523353b1acfb02569df90a98fd5061342f652f7cc9ab4958e97e4ab0eee", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:existing_subnet", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet", "path": "documentation/data-sources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/existing_subnet/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3232023011332203-3320333301223100-1310313101233011-1231312322330121-0300303210122111-1220121102113133-0301231130003200-0013201322013220", "registry_path": "docs/guides/data-sources--gcp_vpc_site--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["voltstack_cluster", "site_local_subnet", "existing_subnet"], "schema_version": 1, "sections": [{"aliases": ["subnet name"], "anchor": "schema-voltstack_cluster--site_local_subnet--existing_subnet--subnet_name", "description": "Name of your subnet in VPC network.", "document_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:voltstack_cluster:site_local_subnet:existing_subnet", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["voltstack_cluster", "site_local_subnet", "existing_subnet", "subnet_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/existing_subnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Name of existing GCP subnet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# voltstack_cluster.site_local_subnet.existing_subnet

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/)
- [voltstack_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/)
- [voltstack_cluster.site_local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/)
- voltstack_cluster.site_local_subnet.existing_subnet

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for existing subnet.

Upstream description:

Name of existing GCP subnet.

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

<a id="schema-voltstack_cluster--site_local_subnet--existing_subnet--subnet_name"></a>

### subnet_name property

Type: `"string"`. Computed.

VPC Subnet Name. Name of your subnet in VPC network.

Upstream description:

Name of your subnet in VPC network.

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

## Next pages

- [voltstack_cluster.site_local_subnet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/properties/voltstack_cluster/site_local_subnet/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/gcp_vpc_site/)
