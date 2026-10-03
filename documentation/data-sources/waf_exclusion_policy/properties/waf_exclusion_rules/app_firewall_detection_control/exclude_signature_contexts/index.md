---
page_title: "waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts"
subcategory: ""
description: "Signature IDs to be excluded for the defined match criteria."
xcsh_docs: {"aliases": ["waf exclusion rules app firewall detection control exclude signature contexts"], "body_bytes": 7145, "body_sha256": "sha256:61b187bf33b119089d882575643864ea22ee5cc0f5890f592c68026d938e9b88", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "parent_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "path": "documentation/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3123010201210221-2232013322033120-2302013100133031-0132121120300123-2033023321220211-3230031213013302-1220221121313113-2010022230022302", "registry_path": "docs/guides/data-sources--waf_exclusion_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_signature_contexts"], "schema_version": 1, "sections": [{"aliases": ["waf exclusion rules app firewall detection control exclude signature contexts context"], "anchor": "schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--context", "description": "The available contexts for Exclusion rules. - CONTEXT_ANY: CONTEXT_ANY Detection will be excluded for all contexts. - CONTEXT_BODY: CONTEXT_BODY Detection will be excluded for the request body. - CONTEXT_REQUEST: CONTEXT_REQUEST Detection will be excluded for the request. - CONTEXT_RESPONSE: CONTEXT_RESPONSE -", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_signature_contexts", "context"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf exclusion rules app firewall detection control exclude signature contexts context name"], "anchor": "schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--context_name", "description": "Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an wildcard asterisk (*).", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_signature_contexts", "context_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf exclusion rules app firewall detection control exclude signature contexts signature id"], "anchor": "schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--signature_id", "description": "The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that all signatures will be excluded for the specified context.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control:exclude_signature_contexts", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control", "exclude_signature_contexts", "signature_id"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/exclude_signature_contexts/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Signature IDs to be excluded for the defined match criteria.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts

Breadcrumbs:

- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/)
- [waf_exclusion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/)
- [waf_exclusion_rules.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/)
- waf_exclusion_rules.app_firewall_detection_control.exclude_signature_contexts

<a id="section"></a>

Type: `"list"`. Computed.

Signature IDs to be excluded for the defined match criteria.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1024,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1024",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--context"></a>

### context property

Type: `"string"`. Computed.

\[Enum:
CONTEXT\_ANY|CONTEXT\_BODY|CONTEXT\_REQUEST|CONTEXT\_RESPONSE|CONTEXT\_PARAMETER|CONTEXT\_HEADER|CONTEXT\_COOKIE|CONTEXT\_URL|CONTEXT\_URI\]
The available contexts for Exclusion rules. - CONTEXT\_ANY: CONTEXT\_ANY Detection will be excluded
for all contexts. - CONTEXT\_BODY: CONTEXT\_BODY Detection will be excluded for the request body. -
CONTEXT\_REQUEST: CONTEXT\_REQUEST Detection will be excluded for the request. - CONTEXT\_RESPONSE..
Possible values are \`CONTEXT\_ANY\`, \`CONTEXT\_BODY\`, \`CONTEXT\_REQUEST\`,
\`CONTEXT\_RESPONSE\`, \`CONTEXT\_PARAMETER\`, \`CONTEXT\_HEADER\`, \`CONTEXT\_COOKIE\`,
\`CONTEXT\_URL\`, \`CONTEXT\_URI\`. Defaults to \`CONTEXT\_ANY\`.

Upstream description:

The available contexts for Exclusion rules.

&#8203;- CONTEXT\_ANY: CONTEXT\_ANY

Detection will be excluded for all contexts. &#8203;- CONTEXT\_BODY: CONTEXT\_BODY

Detection will be excluded for the request body. &#8203;- CONTEXT\_REQUEST: CONTEXT\_REQUEST

Detection will be excluded for the request. &#8203;- CONTEXT\_RESPONSE: CONTEXT\_RESPONSE

&#8203;- CONTEXT\_PARAMETER: CONTEXT\_PARAMETER

Detection will be excluded for the parameters. The parameter name is required in the Context name
field. If the field is left empty, the detection will be excluded for all parameters. &#8203;-
CONTEXT\_HEADER: CONTEXT\_HEADER

Detection will be excluded for the headers. The header name is required in the Context name field.
If the field is left empty, the detection will be excluded for all headers. &#8203;-
CONTEXT\_COOKIE: CONTEXT\_COOKIE

Detection will be excluded for the cookies. The cookie name is required in the Context name field.
If the field is left empty, the detection will be excluded for all cookies. &#8203;- CONTEXT\_URL:
CONTEXT\_URL

Detection will be excluded for the request URL. &#8203;- CONTEXT\_URI: CONTEXT\_URI.

Receipt-pinned upstream constraints:

```json
{
  "default": "CONTEXT_ANY",
  "enum": [
    "CONTEXT_ANY",
    "CONTEXT_BODY",
    "CONTEXT_REQUEST",
    "CONTEXT_RESPONSE",
    "CONTEXT_PARAMETER",
    "CONTEXT_HEADER",
    "CONTEXT_COOKIE",
    "CONTEXT_URL",
    "CONTEXT_URI"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--context_name"></a>

### context_name property

Type: `"string"`. Computed.

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Upstream description:

Relevant only for contexts: Header, Cookie and Parameter. Name of the Context that the WAF Exclusion
Rules will check. Wildcard matching can be used by prefixing or suffixing the context name with an
wildcard asterisk (\*).

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="schema-waf_exclusion_rules--app_firewall_detection_control--exclude_signature_contexts--signature_id"></a>

### signature_id property

Type: `"number"`. Computed.

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Upstream description:

The allowed values for signature ID are 0 and in the range of 200000001-299999999. 0 implies that
all signatures will be excluded for the specified context.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 299999999,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "299999999"
  }
}
```

## Next pages

- [waf_exclusion_rules.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/)
- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/)
