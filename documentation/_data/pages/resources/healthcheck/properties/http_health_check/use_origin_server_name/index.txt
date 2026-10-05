---
page_title: "http_health_check.use_origin_server_name"
subcategory: "Monitoring"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["backend servers", "http health check use origin server name", "origin servers", "upstream servers"], "body_bytes": 1360, "body_sha256": "sha256:0b2d71f50e2a396044d78f2ac4c84d040e0248486e1ea24641e252eadc15f70c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:properties:http_health_check:use_origin_server_name", "parent_id": "xcsh-docs:resources:healthcheck:properties:http_health_check", "path": "documentation/resources/healthcheck/properties/http_health_check/use_origin_server_name/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2132133223202302-3212113232122201-0322200101231130-3023020122113221-1213112133312113-2313100320200122-0211103112310313-0300123102313030", "registry_path": "docs/guides/resources--healthcheck--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_health_check", "use_origin_server_name"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/properties/http_health_check/use_origin_server_name/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_health_check.use_origin_server_name

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/)
- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/)
- http_health_check.use_origin_server_name

<a id="section"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

This can be used for messages where no values are needed.

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

Terraform syntax:

```terraform
use_origin_server_name = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/properties/http_health_check/)
- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
