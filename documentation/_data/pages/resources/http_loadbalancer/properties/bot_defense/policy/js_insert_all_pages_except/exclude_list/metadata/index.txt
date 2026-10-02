---
page_title: "bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata"
subcategory: "Load Balancing"
description: "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs."
xcsh_docs: {"aliases": ["bot defense policy js insert all pages except exclude list metadata"], "body_bytes": 4875, "body_sha256": "sha256:01a2eadfde6a5fd6b2c21e4433ce58d78535e581e3b614e6641fec707d6f639c", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/metadata/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2311013103031000-3033100120030313-3300312013130000-1001203302010231-0102201032321131-0321022130232121-3221210322132101-3333033023221113", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-011.md", "relationships": [{"anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--metadata--name", "enforcement": "provider-schema", "group": "bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "metadata"], "schema_version": 1, "sections": [{"aliases": ["description spec"], "anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--metadata--description_spec", "description": "Description. Human readable description.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "metadata", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--metadata--name", "description": "This is the name of the message. The value of name has to follow DNS-1035 format.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except:exclude_list:metadata", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except", "exclude_list", "metadata", "name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/metadata/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name")}
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
metadata {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--metadata--description_spec"></a>

### description_spec property

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="schema-bot_defense--policy--js_insert_all_pages_except--exclude_list--metadata--name"></a>

### name property

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

## Next pages

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/exclude_list/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
