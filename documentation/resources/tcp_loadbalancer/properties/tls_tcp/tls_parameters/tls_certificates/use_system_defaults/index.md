---
page_title: "tls_tcp.tls_parameters.tls_certificates.use_system_defaults"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls tcp tls parameters tls certificates use system defaults"], "body_bytes": 1737, "body_sha256": "sha256:59832b49283c4d1835cb3b31a253125aff5361581aa0c28f95d5f93dd0df4323", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:use_system_defaults", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "path": "documentation/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/use_system_defaults/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0102010120213331-3223223023110201-0122232110302212-0323110331103310-3031012201011103-1013011002200011-3202010123303321-1231222302131132", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "use_system_defaults"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/use_system_defaults/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_parameters.tls_certificates.use_system_defaults

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [tls_tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/)
- [tls_tcp.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/)
- [tls_tcp.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/)
- tls_tcp.tls_parameters.tls_certificates.use_system_defaults

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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
use_system_defaults = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_tcp.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
