---
page_title: "http_health_check.disable_virtual_host"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["http health check disable virtual host"], "body_bytes": 1337, "body_sha256": "sha256:4ebd5bc4c07dde2858ea4879d10339139ce5fa7f82fdb9dbcc39dea5c7554c44", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check:disable_virtual_host", "parent_id": "xcsh-docs:resources:dns_lb_health_check:properties:http_health_check", "path": "documentation/resources/dns_lb_health_check/properties/http_health_check/disable_virtual_host/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1220313032231202-1002300301213000-1103213221323331-0322331110013232-0011321232220222-1112302020313320-0321110201220330-1101231220110331", "registry_path": "docs/guides/resources--dns_lb_health_check--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["http_health_check", "disable_virtual_host"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/properties/http_health_check/disable_virtual_host/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# http_health_check.disable_virtual_host

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/)
- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/)
- http_health_check.disable_virtual_host

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
disable_virtual_host = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [http_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/http_health_check/)
- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
