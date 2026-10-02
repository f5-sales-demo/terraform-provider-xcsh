---
page_title: "custom_anonymization"
subcategory: "Security"
description: "Anonymization settings which is a list of HTTP headers, parameters and cookies."
xcsh_docs: {"aliases": ["custom anonymization"], "body_bytes": 2367, "body_sha256": "sha256:b1d3e6769ad0514844ffff3b43f61fadac3670486875f92d932e137aa7aad4b1", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization", "parent_id": "xcsh-docs:resources:app_firewall:reference", "path": "documentation/resources/app_firewall/properties/custom_anonymization/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0211102231112321-3303232230133311-0000200300103313-2303321211111030-1013331030331131-0002213130220302-2131101302112300-0121333312111003", "registry_path": "docs/guides/resources--app_firewall--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization:RequiredObjectAttributes:anonymization_config", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_anonymization"], "schema_version": 1, "sections": [{"aliases": ["anonymization config"], "anchor": "section", "description": "List of HTTP headers, cookies and query parameters whose values will be masked.", "document_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,http_header", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:cookie", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,http_header", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:http_header,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:http_header", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:cookie,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_anonymization.anonymization_config:ConflictingListObjectAttributes:http_header,query_parameter", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:app_firewall:properties:custom_anonymization:anonymization_config:query_parameter", "type": "conflicts"}], "schema_path": ["custom_anonymization", "anonymization_config"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/properties/custom_anonymization/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Anonymization settings which is a list of HTTP headers, parameters and cookies.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_anonymization

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- custom_anonymization

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: custom\_anonymization, default\_anonymization, disable\_anonymization; Default:
default\_anonymization\] Anonymization settings which is a list of HTTP headers, parameters and
cookies.

Upstream description:

Anonymization settings which is a list of HTTP headers, parameters and cookies.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("anonymization_config")}
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

OneOf alternatives in this subsection:

- [custom_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/#section)
- [default_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/default_anonymization/#section)
- [disable_anonymization](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/disable_anonymization/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
custom_anonymization {
  # Configure direct properties listed below.
}
```

## Direct properties

- [anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/): complete subsection reference.

## Next pages

- [custom_anonymization.anonymization_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/custom_anonymization/anonymization_config/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/properties/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
