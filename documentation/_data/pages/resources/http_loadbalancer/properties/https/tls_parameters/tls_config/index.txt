---
page_title: "https.tls_parameters.tls_config"
subcategory: "Load Balancing"
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["https tls parameters tls config"], "body_bytes": 3753, "body_sha256": "sha256:27de9c6b50c07fadb8e8a805350770137e363bff3ec039cfc875d7299a610e8a", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:custom_security", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:default_security", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:medium_security"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters", "path": "documentation/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1331301222322132-3310303011021332-0330022222232121-3232113210300302-3100323021101302-2323203211312232-2232012310223322-1333222331330000", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-019.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:custom_security,default_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:custom_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:custom_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:custom_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:custom_security,default_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:default_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:default_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:default_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:custom_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:default_security,low_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:low_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:custom_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:medium_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:default_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:medium_security", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config:ConflictingObjectAttributes:low_security,medium_security", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:medium_security", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["https", "tls_parameters", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["https tls parameters tls config custom security"], "anchor": "section", "description": "This defines TLS protocol config including min/max versions and allowed ciphers.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:custom_security", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-https--tls_parameters--tls_config--custom_security--cipher_suites", "enforcement": "provider-schema", "group": "https.tls_parameters.tls_config.custom_security:RequiredObjectAttributes:cipher_suites", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:custom_security", "type": "requires"}], "schema_path": ["https", "tls_parameters", "tls_config", "custom_security"], "syntax": "block", "type": "object"}, {"aliases": ["https tls parameters tls config default security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:default_security", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "tls_parameters", "tls_config", "default_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https tls parameters tls config low security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:low_security", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "tls_parameters", "tls_config", "low_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https tls parameters tls config medium security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:https:tls_parameters:tls_config:medium_security", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https", "tls_parameters", "tls_config", "medium_security"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https.tls_parameters.tls_config

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [https](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/)
- [https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/)
- https.tls_parameters.tls_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/medium_security/): complete subsection reference.

## Next pages

- [https.tls_parameters.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/custom_security/)
- [https.tls_parameters.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/default_security/)
- [https.tls_parameters.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/low_security/)
- [https.tls_parameters.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/tls_config/medium_security/)
- [https.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https/tls_parameters/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
