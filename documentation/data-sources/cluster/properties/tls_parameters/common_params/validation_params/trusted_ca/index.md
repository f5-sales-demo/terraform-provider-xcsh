---
page_title: "tls_parameters.common_params.validation_params.trusted_ca"
subcategory: ""
description: "Reference to Root CA Certificate."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls parameters common params validation params trusted ca"], "body_bytes": 2038, "body_sha256": "sha256:cca1c51027c737d8f7e3d2a7ae2c039932572e6483af5499cd2babfe86366c40", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": ["xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:validation_params:trusted_ca:trusted_ca_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:validation_params:trusted_ca", "parent_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:validation_params", "path": "documentation/data-sources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2131030032102121-2232130121132133-2222233320102200-3020321232333101-3320331002022311-0230030310122103-1122032113030020-1201030230302231", "registry_path": "docs/guides/data-sources--cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "common_params", "validation_params", "trusted_ca"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "trusted ca list"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:data-sources:cluster:properties:tls_parameters:common_params:validation_params:trusted_ca:trusted_ca_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["tls_parameters", "common_params", "validation_params", "trusted_ca", "trusted_ca_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Reference to Root CA Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["clusterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.validation_params.trusted_ca

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/)
- [tls_parameters.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/validation_params/)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="section"></a>

Type: `"single"`. Computed.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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

- [trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/): complete subsection reference.

## Next pages

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/)
- [tls_parameters.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/properties/tls_parameters/common_params/validation_params/)
- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cluster/)
