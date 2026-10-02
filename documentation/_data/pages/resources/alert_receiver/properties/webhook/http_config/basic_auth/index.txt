---
page_title: "webhook.http_config.basic_auth"
subcategory: ""
description: "Authorization parameters to access HTPP alert Receiver Endpoint."
xcsh_docs: {"aliases": ["webhook http config basic auth"], "body_bytes": 2925, "body_sha256": "sha256:e3df6eb510527ea2192d2d9b2f2872a704f0fe898e37318f0ddab03f92ae0473", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:alert_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "parent_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config", "path": "documentation/resources/alert_receiver/properties/webhook/http_config/basic_auth/index.md", "product": "distributed-cloud", "provider_name": "alert_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2203121002000122-2002301012020130-3132012311001303-1113313133330121-2131210012023303-0113321023210102-3012313020303322-1333300320111233", "registry_path": "docs/guides/resources--alert_receiver--reference--group-001.md", "relationships": [{"anchor": "schema-webhook--http_config--basic_auth--user_name", "enforcement": "provider-schema", "group": "webhook.http_config.basic_auth:RequiredObjectAttributes:user_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["webhook", "http_config", "basic_auth"], "schema_version": 1, "sections": [{"aliases": ["password"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.basic_auth.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "webhook.http_config.basic_auth.password:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth:password:clear_secret_info", "type": "conflicts"}], "schema_path": ["webhook", "http_config", "basic_auth", "password"], "syntax": "block", "type": "object"}, {"aliases": ["user name"], "anchor": "schema-webhook--http_config--basic_auth--user_name", "description": "HTTP Basic Auth User Name.", "document_id": "xcsh-docs:resources:alert_receiver:properties:webhook:http_config:basic_auth", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["webhook", "http_config", "basic_auth", "user_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/alert_receiver/properties/webhook/http_config/basic_auth/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Authorization parameters to access HTPP alert Receiver Endpoint.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["alert_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# webhook.http_config.basic_auth

Breadcrumbs:

- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/)
- [webhook](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- webhook.http_config.basic_auth

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Authorization parameters to access HTPP alert Receiver Endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("user_name")}
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
basic_auth {
  # Configure direct properties listed below.
}
```

## Direct properties

- [password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/): complete subsection reference.

<a id="schema-webhook--http_config--basic_auth--user_name"></a>

### user_name property

Type: `"string"`. Optional.

User Name. HTTP Basic Auth User Name.

Upstream description:

HTTP Basic Auth User Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

## Next pages

- [webhook.http_config.basic_auth.password](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/basic_auth/password/)
- [webhook.http_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/properties/webhook/http_config/)
- [xcsh_alert_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/alert_receiver/)
