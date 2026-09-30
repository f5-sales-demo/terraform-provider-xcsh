---
page_title: "cloudflare.protected_endpoints.web_client.continue"
subcategory: ""
description: "cloudflare.protected_endpoints.web_client.continue for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2238, "body_sha256": "sha256:e1416a5a52935a1b9471953fef6526f7dc900de82851ef13cf5ef39e7a232c3e", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue:add_header", "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue:no_header"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client:continue", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudflare:protected_endpoints:web_client", "path": "docs/guides/resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudflare", "protected_endpoints", "web_client", "continue"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudflare/protected_endpoints/web_client/continue/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudflare.protected_endpoints.web_client.continue for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# cloudflare.protected_endpoints.web_client.continue

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudflare](resources--protected_application--properties--cloudflare.md)
- [cloudflare.protected_endpoints](resources--protected_application--properties--cloudflare--protected_endpoints.md)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--properties--cloudflare--protected_endpoints--web_client.md)
- cloudflare.protected_endpoints.web_client.continue

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
continue {
  # Configure direct properties listed below.
}
```

## Direct properties

- [add_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--add_header.md): complete subsection reference.

- [no_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--no_header.md): complete subsection reference.

## Next pages

- [cloudflare.protected_endpoints.web_client.continue.add_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--add_header.md)
- [cloudflare.protected_endpoints.web_client.continue.no_header](resources--protected_application--properties--cloudflare--protected_endpoints--web_client--continue--no_header.md)
- [cloudflare.protected_endpoints.web_client](resources--protected_application--properties--cloudflare--protected_endpoints--web_client.md)
- [xcsh_protected_application](../resources/protected_application.md)
