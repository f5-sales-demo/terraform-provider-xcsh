---
page_title: "captcha_challenge"
subcategory: "Load Balancing"
description: "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML"
xcsh_docs: {"aliases": ["captcha challenge", "succeeded", "success", "successful"], "body_bytes": 4853, "body_sha256": "sha256:d03bae5e3817f8cbc8019ee939627fd7b12dc4e3905d59aee448513efd350246", "capabilities": ["cdn"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:captcha_challenge", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:reference", "path": "documentation/data-sources/cdn_loadbalancer/properties/captcha_challenge/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1103002330300201-1231313031223333-0303202310013222-0023112211102302-3010003113123000-2212320100321323-2032133101121303-3112310212212213", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-009.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["captcha_challenge"], "schema_version": 1, "sections": [{"aliases": ["captcha challenge cookie expiry"], "anchor": "schema-captcha_challenge--cookie_expiry", "description": "Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new challenge.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:captcha_challenge", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["captcha_challenge", "cookie_expiry"], "syntax": "attribute", "type": "number"}, {"aliases": ["captcha challenge custom page"], "anchor": "schema-captcha_challenge--custom_page", "description": "Custom message is of type uri_ref. Currently supported URL schemes is string:///. For string:/// scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded plain text message e.g. \"Please Wait..\" or it can be HTML paragraph or a body string encoded as base64 string E.g. \"<p>", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:captcha_challenge", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["captcha_challenge", "custom_page"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/captcha_challenge/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# captcha_challenge

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- captcha_challenge

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: captcha\_challenge, enable\_challenge, js\_challenge, no\_challenge,
policy\_based\_challenge; Default: no\_challenge\] Enables loadbalancer to perform captcha challenge
Captcha challenge will be based on Google Recaptcha. With this feature enabled, only clients that
pass the captcha challenge will be allowed to complete the HTTP request. When loadbalancer is
configured to do Captcha Challenge, it will redirect..

Additional upstream details:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha. When loadbalancer is configured to do Captcha
Challenge, it will redirect the browser to an HTML page on every new HTTP request. This HTML page
will have captcha challenge embedded in it. Client will be allowed to make the request only if the
captcha challenge is successful. Loadbalancer will tag response header with a cookie to avoid
Captcha challenge for subsequent requests. CAPTCHA is mainly used as a security check to ensure only
human users can pass through. Generally, computers or bots are not capable of solving a captcha. You
can enable either Javascript challenge or Captcha challenge on a virtual host.

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

- [captcha_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/captcha_challenge/#section)
- [enable_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/enable_challenge/#section)
- [js_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/js_challenge/#section)
- [no_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/no_challenge/#section)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/policy_based_challenge/#section)

Select alternatives according to the provider validators above.

## Direct properties

<a id="schema-captcha_challenge--cookie_expiry"></a>

### cookie_expiry property

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="schema-captcha_challenge--custom_page"></a>

### custom_page property

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
