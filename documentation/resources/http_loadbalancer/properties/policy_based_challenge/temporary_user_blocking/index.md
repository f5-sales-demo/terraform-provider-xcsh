---
page_title: "policy_based_challenge.temporary_user_blocking"
subcategory: "Load Balancing"
description: "Specifies configuration for temporary user blocking resulting from user behavior analysis. When Malicious User Mitigation is enabled from service policy rules, users' accessing the application will be analyzed for malicious activity and the configured mitigation actions will be taken on identified malicious users."
xcsh_docs: {"aliases": ["policy based challenge temporary user blocking"], "body_bytes": 3834, "body_sha256": "sha256:fe43ce8fd97931183498840cb3b90e4845e31ec9678e5c631b628096551e346c", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:temporary_user_blocking", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge", "path": "documentation/resources/http_loadbalancer/properties/policy_based_challenge/temporary_user_blocking/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0332322110110101-3123223013302023-0321200302320201-2312111021001220-0032301331120201-2310003212230313-2113130213200312-2022222232311001", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-023.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "temporary_user_blocking"], "schema_version": 1, "sections": [{"aliases": ["custom page"], "anchor": "schema-policy_based_challenge--temporary_user_blocking--custom_page", "description": "Custom message is of type `uri_ref`. Currently supported URL schemes is `string:///`. For `string:///` scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. \"Blocked..\" or it can be HTML paragraph or a body string encoded as base64 string E.g. \"<p>", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:temporary_user_blocking", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy_based_challenge", "temporary_user_blocking", "custom_page"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/temporary_user_blocking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Specifies configuration for temporary user blocking resulting from user behavior analysis. When Malicious User Mitigation is enabled from service policy rules, users' accessing the application will be analyzed for malicious activity and the configured mitigation actions will be taken on identified malicious users.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.temporary_user_blocking

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- policy_based_challenge.temporary_user_blocking

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specifies configuration for temporary user blocking resulting from user behavior analysis. When
Malicious User Mitigation is enabled from service policy rules, users' accessing the application
will be analyzed for malicious activity and the configured mitigation actions will be taken on..

Upstream description:

Specifies configuration for temporary user blocking resulting from user behavior analysis.

When Malicious User Mitigation is enabled from service policy rules, users' accessing the
application will be analyzed for malicious activity and the configured mitigation actions will be
taken on identified malicious users. These mitigation actions include setting up temporary blocking
on that user. This configuration specifies settings on how that blocking should be done by the
loadbalancer.

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
temporary_user_blocking {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-policy_based_challenge--temporary_user_blocking--custom_page"></a>

### custom_page property

Type: `"string"`. Optional.

Custom message is of type . Currently supported URL schemes is . For scheme, message needs to be
encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g.
'Blocked.' or it can be HTML paragraph or a body string encoded as base64 string E.g. '&lt;p&gt;
Blocked..

Upstream description:

Custom message is of type \`uri\_ref\`. Currently supported URL schemes is \`string:///\`. For
\`string:///\` scheme, message needs to be encoded in Base64 format. You can specify this message as
base64 encoded plain text message e.g. "Blocked.." or it can be HTML paragraph or a body string
encoded as base64 string E.g. "&lt;p&gt; Blocked &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

## Next pages

- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
