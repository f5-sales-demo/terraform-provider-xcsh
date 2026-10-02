---
page_title: "http_health_check.use_origin_server_name"
subcategory: "Monitoring"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["backend servers", "http health check use origin server name", "origin servers", "upstream servers"], "body_bytes": 1360, "body_sha256": "sha256:0b2d71f50e2a396044d78f2ac4c84d040e0248486e1ea24641e252eadc15f70c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "id": "xcsh-docs:resources:healthcheck:properties:http_health_check:use_origin_server_name", "parent_id": "xcsh-docs:resources:healthcheck:properties:http_health_check", "path": "documentation/resources/healthcheck/properties/http_health_check/use_origin_server_name/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2132133223202302-3212113232122201-0322200101231130-3023020122113221-1213112133312113-2313100320200122-0211103112310313-0300123102313030", "registry_path": "docs/guides/resources--healthcheck--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_health_check", "use_origin_server_name"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/properties/http_health_check/use_origin_server_name/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
