---
page_title: "tls_tcp.tls_parameters.tls_certificates.private_key"
subcategory: "Load Balancing"
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls tcp tls parameters tls certificates private key"], "body_bytes": 2866, "body_sha256": "sha256:8e2b4ba4a76e4d489c7055235f7bc92de0a4513ee61572248c9ec1142a944a3e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:clear_secret_info"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "path": "documentation/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0100333102101300-3220200132001223-0030120230300122-3331201232232302-1232301113321120-0330221011030203-2110332331210010-2021112323003013", "registry_path": "docs/guides/resources--tcp_loadbalancer--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates.private_key:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:clear_secret_info", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "private_key"], "schema_version": 1, "sections": [{"aliases": ["blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tls_tcp--tls_parameters--tls_certificates--private_key--blindfold_secret_info--location", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info:RequiredObjectAttributes:location", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:blindfold_secret_info", "type": "requires"}], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "private_key", "blindfold_secret_info"], "syntax": "block", "type": "object"}, {"aliases": ["clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:clear_secret_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-tls_tcp--tls_parameters--tls_certificates--private_key--clear_secret_info--url", "enforcement": "provider-schema", "group": "tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info:RequiredObjectAttributes:url", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:private_key:clear_secret_info", "type": "requires"}], "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "private_key", "clear_secret_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_parameters.tls_certificates.private_key

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/)
- [tls_tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/)
- [tls_tcp.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/)
- [tls_tcp.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/)
- tls_tcp.tls_parameters.tls_certificates.private_key

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/clear_secret_info/): complete subsection reference.

## Next pages

- [tls_tcp.tls_parameters.tls_certificates.private_key.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/blindfold_secret_info/)
- [tls_tcp.tls_parameters.tls_certificates.private_key.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/private_key/clear_secret_info/)
- [tls_tcp.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/)
- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
