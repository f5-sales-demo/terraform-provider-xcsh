---
page_title: "use_tls.use_mtls"
subcategory: "Load Balancing"
description: "MTLS Client Certificate."
xcsh_docs: {"aliases": ["use tls use mtls"], "body_bytes": 1661, "body_sha256": "sha256:17d178ac6efec940c9b3555f50cf1cdd89d6b8ea4b9aa5fe5303f4e18d5f58c7", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls", "parent_id": "xcsh-docs:resources:origin_pool:properties:use_tls", "path": "documentation/resources/origin_pool/properties/use_tls/use_mtls/index.md", "product": "distributed-cloud", "provider_name": "origin_pool", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1120222222033012-2312123120002033-1111012231100123-0022300321132031-1012231010221201-0202330220031110-1332122100233330-3222113110201011", "registry_path": "docs/guides/resources--origin_pool--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.use_mtls:RequiredObjectAttributes:tls_certificates", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["use_tls", "use_mtls"], "schema_version": 1, "sections": [{"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "use tls use mtls tls certificates"], "anchor": "section", "description": "MTLS Client Certificate.", "document_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:custom_hash_algorithms", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,disable_ocsp_stapling", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:disable_ocsp_stapling", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:custom_hash_algorithms,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates:ConflictingListObjectAttributes:disable_ocsp_stapling,use_system_defaults", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates:use_system_defaults", "type": "conflicts"}, {"anchor": "schema-use_tls--use_mtls--tls_certificates--certificate_url", "enforcement": "provider-schema", "group": "use_tls.use_mtls.tls_certificates:RequiredListObjectAttributes:certificate_url", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:origin_pool:properties:use_tls:use_mtls:tls_certificates", "type": "requires"}], "schema_path": ["use_tls", "use_mtls", "tls_certificates"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/properties/use_tls/use_mtls/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "MTLS Client Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# use_tls.use_mtls

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/)
- use_tls.use_mtls

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates")}
```

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
use_mtls {
  # Configure direct properties listed below.
}
```

## Direct properties

- [tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/): complete subsection reference.

## Next pages

- [use_tls.use_mtls.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/use_mtls/tls_certificates/)
- [use_tls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/use_tls/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
