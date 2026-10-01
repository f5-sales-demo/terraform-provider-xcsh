---
page_title: "cloudflare.protected_endpoints.web_mobile_client.continue_mobile"
subcategory: ""
description: "cloudflare.protected_endpoints.web_mobile_client.continue_mobile for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2484, "body_sha256": "sha256:03b2eafcbb745680dbbf0b19bc8ec38c5d243716e85afd67dd9cc1f3d84fe4b7", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile:add_header", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile:no_header"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client:continue_mobile", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_mobile_client", "path": "docs/guides/resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_mobile_client", "continue_mobile"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/protected_endpoints/web_mobile_client/continue_mobile/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.protected_endpoints.web_mobile_client.continue_mobile for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudflare.protected_endpoints.web_mobile_client.continue_mobile

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudflare](resources--protected_application--properties--cloudflare.md)
- [cloudflare.protected_endpoints](resources--protected_application--properties--cloudflare--protected_endpoints.md)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client.md)
- cloudflare.protected_endpoints.web_mobile_client.continue_mobile

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Continue Bot Mitigation Action. Continue mitigation action.

Upstream description:

Continue mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("add_header",
    "no_header")}
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
  "x-ves-oneof-field-add_header_choice": "[\"add_header\",\"no_header\"]"
}
```

Terraform syntax:

```terraform
continue_mobile {
  # Configure direct properties listed below.
}
```

## Direct properties

- [add_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile--add_header.md): complete subsection reference.

- [no_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile--no_header.md): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.add_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile--add_header.md)
- [cloudflare.protected_endpoints.web_mobile_client.continue_mobile.no_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client--continue_mobile--no_header.md)
- [cloudflare.protected_endpoints.web_mobile_client](resources--protected_application--properties--cloudflare--protected_endpoints--web_mobile_client.md)
- [xcsh_protected_application](../resources/protected_application.md)
