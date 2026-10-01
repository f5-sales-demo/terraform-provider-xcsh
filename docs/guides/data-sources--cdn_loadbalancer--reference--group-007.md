---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-c895c315c59f954267fb2396adb379339d973a6fc5336db825969de517df0a85"></a>

## Next pages — bot_defense.enable_cors_support / b3df8d8ecd36 / 4

- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3c94d6c74f07e5cfc3b6cbd24eb458799b2902e60e2f37fc0a0aad47f8166d0"></a>

## bot_defense.policy — bot_defense.policy / 75ee2bc8a687 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- bot_defense.policy

<a id="canonical-b8332217e1344d78482aff9803d8e868d50c80621c5a76abdc09f1076810eb68"></a>

Type: `"single"`. Computed.

Defines various configuration OPTIONS for Bot Defense policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

<a id="canonical-7c673160b8745b5e76ede156bbbfdd1c5d3e3ac0b5824f1ab8241bca02047c65"></a>

## Direct properties — bot_defense.policy / 75ee2bc8a687 / 3

- [disable_js_insert](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-bda9cbffef603f986393da44620a074413d17d114bd52fbc80d1b02bc697d55d): complete subsection reference.

- [disable_mobile_sdk](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a4b3eeb297eab3f74c59bf6e042051d42a33c4fe84b29c134e3412a13d15bf3a): complete subsection reference.

<a id="canonical-456e628ccd78188cd98a90c2857179c6a4543a798ac5b059812c4dadb8cd8e56"></a>

<a id="canonical-ac29b3066314779a2e4eba4480420e78c84c886d91c4994d36fbc21a86a15f55"></a>

## javascript_mode property — bot_defense.policy / 75ee2bc8a687 / 4

Type: `"string"`. Computed.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Upstream description:

Web Client JavaScript Mode.

Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable
Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable.

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-71a1c9f70e56a4abd59bffb923495511dee4c7acc38550f5d291037b9c11b4bd"></a>

<a id="canonical-c5c8fe378d2c00a60e98cc55f83df3a50c2ef70c2962d7df4ae98ec7cc8828dd"></a>

## js_download_path property — bot_defense.policy / 75ee2bc8a687 / 5

Type: `"string"`. Computed.

Customize Bot Defense Client JavaScript path. If not specified, default

Upstream description:

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1163add6993b1c929d4fe9831c23ca1c5981efd73bd5a1a61d2de6bc6e24b555): complete subsection reference.

- [js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8a9473a93a8f8218afddc636685505997cdb0a0a0fa8a2159405565215071628): complete subsection reference.

- [js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d): complete subsection reference.

- [mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a6d501fcf23450d0d973e5b1aa3f37d111c82769e7e6e2f1ca962105445bc4f3): complete subsection reference.

- [protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b): complete subsection reference.

<a id="canonical-6a3dc00ac5c005ab9d64baeb25fd6cce32544d6895306ad90208d5830ca56f7e"></a>

## Next pages — bot_defense.policy / 75ee2bc8a687 / 6

- [bot_defense.policy.disable_js_insert](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-bda9cbffef603f986393da44620a074413d17d114bd52fbc80d1b02bc697d55d)
- [bot_defense.policy.disable_mobile_sdk](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a4b3eeb297eab3f74c59bf6e042051d42a33c4fe84b29c134e3412a13d15bf3a)
- [bot_defense.policy.js_insert_all_pages](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1163add6993b1c929d4fe9831c23ca1c5981efd73bd5a1a61d2de6bc6e24b555)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8a9473a93a8f8218afddc636685505997cdb0a0a0fa8a2159405565215071628)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a6d501fcf23450d0d973e5b1aa3f37d111c82769e7e6e2f1ca962105445bc4f3)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-bda9cbffef603f986393da44620a074413d17d114bd52fbc80d1b02bc697d55d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1bd507d80c6c4b7ad27833184984f5a91f4f1c9b4eaa29f71ea459f96ae4f668"></a>

## bot_defense.policy.disable_js_insert — bot_defense.policy.disable_js_insert / f93114d2c736 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- bot_defense.policy.disable_js_insert

<a id="canonical-f00da27dac2f6841e19f9b5caea923825b68b0866a6ec353e330a9dd8433a61e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable js insert.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-a2c278124970d6baa952c9bf6262251e4d862d8e8b64ff5babff48c432cccb99"></a>

## Direct properties — bot_defense.policy.disable_js_insert / f93114d2c736 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a4ea0b249f256515f580294d9d0cc97809bd0961e87aa151838610c510217e7f"></a>

## Next pages — bot_defense.policy.disable_js_insert / f93114d2c736 / 4

- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a4b3eeb297eab3f74c59bf6e042051d42a33c4fe84b29c134e3412a13d15bf3a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f3e2659bd75bd08b8757cecac345f76e7afb6c612b183ae4686c26969adaf5d"></a>

## bot_defense.policy.disable_mobile_sdk — bot_defense.policy.disable_mobile_sdk / 1faf312df88f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-7b0a7043cf485414ee5e78260d81e589b79a87136c8eea7bddc4fb000bfd1ec5"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-f27b84b523b4165a530ca5dc75c6f6debfaf4bd6b2848bd630e19b14eb838656"></a>

## Direct properties — bot_defense.policy.disable_mobile_sdk / 1faf312df88f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-7b4a07fe3cb2e8e58f8dfc2e4ad414cbf029eacc8af684eb4dd400e5097110a9"></a>

## Next pages — bot_defense.policy.disable_mobile_sdk / 1faf312df88f / 4

- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-1163add6993b1c929d4fe9831c23ca1c5981efd73bd5a1a61d2de6bc6e24b555"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1d0751b2c89b9a62703430b203103620f099d10e8b8b1e26a3e7e4e752f94254"></a>

## bot_defense.policy.js_insert_all_pages — bot_defense.policy.js_insert_all_pages / 787e7f0470a4 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-ee63035870bec03840b1d62bba631dc248badc9b9fa7a77a1bab72452b00bae7"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages.

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

<a id="canonical-f4ded3a35d7b653bc1cad38a2d30dade74fe898794c52355b92bbb93aba1368a"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages / 787e7f0470a4 / 3

<a id="canonical-054fb0938db12b8f44a1ccab02493404cfe381dd5ab4adfff35722e6f9cc7e62"></a>

<a id="canonical-e55a48182a67ac12533bf2a3a9aceb31e5f44621b15b777f9c841a7de141c620"></a>

## javascript_location property — bot_defense.policy.js_insert_all_pages / 787e7f0470a4 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ce7c98008c988fcf0a23afc6a6d78b942911c3d64c6dd6684b37a0d31bdc68b4"></a>

## Next pages — bot_defense.policy.js_insert_all_pages / 787e7f0470a4 / 5

- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8a9473a93a8f8218afddc636685505997cdb0a0a0fa8a2159405565215071628"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-06a4ef9c2662457390e7bee19a81182a27d3f6caf62d598979b4def64ebda86c"></a>

## bot_defense.policy.js_insert_all_pages_except — bot_defense.policy.js_insert_all_pages_except / 4ce7e2818497 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- bot_defense.policy.js_insert_all_pages_except

<a id="canonical-6242aa4b439266663992a61536f3778c6dfb53365aeabaf2fde908ba82056cba"></a>

Type: `"single"`. Computed.

Insert Bot Defense JavaScript in all pages with the exceptions.

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

<a id="canonical-3263ad5cbefba1e1e552131282a8037f546a97a80d3d0ae1b5bf22c240bb89a1"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except / 4ce7e2818497 / 3

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4): complete subsection reference.

<a id="canonical-87c16d441a40b962ece9c0b1c890298c649e34d95bf0102b91b7899209502c24"></a>

<a id="canonical-5fdd6f037f46ae8b3bb6d0c6156fa0849b1ad7504083680ca59f365bcacf13f1"></a>

## javascript_location property — bot_defense.policy.js_insert_all_pages_except / 4ce7e2818497 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-5977664927612acfe425d0efaef26bb02ebdbb81dedcd39407f7cb4d148f9a54"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except / 4ce7e2818497 / 5

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe6c97911f2d7e68d00442f68d4d2d1125044243087428e8be0a444853947c1a"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list — bot_defense.policy.js_insert_all_pages_except.exclude_list / 779386283be6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8a9473a93a8f8218afddc636685505997cdb0a0a0fa8a2159405565215071628)
- bot_defense.policy.js_insert_all_pages_except.exclude_list

<a id="canonical-76a9bde21032863e06dbfcb5f7e3a34cc24789b621fa33b7f4f0457687743b4b"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-580aad9ddf6917ce8b99a987814133a0818a829a16e216e59ab112540c2a32c5"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list / 779386283be6 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-e9287d6473ff6510ed88f1a566391604c63d469b962567d3cd644a3db03e561e): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-f275fc60b6ff7b392d292d3a8cfef1b03caafbcb4c201a6e9c9b8e746018aac9): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-c206e315bcd5af1541c637d035ba7ec96a8b10899be84bcb603c4a1447af6c1a): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3378d3f112c6c7776f14117ce951ed99731e51a07dc56a9db0d615f6fe8575f3): complete subsection reference.

<a id="canonical-5a22a3ece8746d4781a2dcf8d58f4d93a5a25400f9358787c5afdf266fa35ad2"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list / 779386283be6 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-e9287d6473ff6510ed88f1a566391604c63d469b962567d3cd644a3db03e561e)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-f275fc60b6ff7b392d292d3a8cfef1b03caafbcb4c201a6e9c9b8e746018aac9)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-c206e315bcd5af1541c637d035ba7ec96a8b10899be84bcb603c4a1447af6c1a)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list.path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3378d3f112c6c7776f14117ce951ed99731e51a07dc56a9db0d615f6fe8575f3)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8a9473a93a8f8218afddc636685505997cdb0a0a0fa8a2159405565215071628)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-e9287d6473ff6510ed88f1a566391604c63d469b962567d3cd644a3db03e561e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97508880407e0e99ed5c1a26dd8125340a7982a158f7008806b92d8536f4370b"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 2fa15dfc9de3 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8a9473a93a8f8218afddc636685505997cdb0a0a0fa8a2159405565215071628)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain

<a id="canonical-2c364505ea022b7df29cda0a39419d48c164657dba44d5b36426689b9a9d7228"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-e245e6dbeaa4725ca85e032c8f9641a52ef835de239d58712989e7f9cb06d139"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 2fa15dfc9de3 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a5ae82b558124de3cec5debd4fac89df59aa0d79c767a2a1c24434f419d52ec2"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.any_domain / 2fa15dfc9de3 / 4

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f275fc60b6ff7b392d292d3a8cfef1b03caafbcb4c201a6e9c9b8e746018aac9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-194dadde00305089f121b0d13972693e1937b81e7ce7aa38f9a28d0d50d4f3c8"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.domain — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 808a10b5dc57 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8a9473a93a8f8218afddc636685505997cdb0a0a0fa8a2159405565215071628)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.domain

<a id="canonical-b3748364eecfd35fb2163f4948932498197c01d7c2433a04eb9cae957757724e"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-ebda14f6e7c21e55982c202c33a3453209b24687bbe3cc5c11200c06bac24ff6"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 808a10b5dc57 / 3

<a id="canonical-bc73ea9d2fdaa98d4d9bad5d6aec94c85a0f4320e0ae1bc49de2ce8e6e7d681c"></a>

<a id="canonical-215d7c0d888012cedc8b06c28125ce22ffb2888b0575e91beebcdb1a7d277db5"></a>

## exact_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 808a10b5dc57 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-764b1a30aa097a4c1db11636946612ac2ad782ea75b95c2f84bb6c967e97347e"></a>

<a id="canonical-12a715bf7c00616247f43ee222a444366c5dfef9cad1c95cf1d2e14f24493775"></a>

## regex_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 808a10b5dc57 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-98b476a3a702f9df910c47fbc4ce3535c43aa17afd35028e92e07a660500f815"></a>

<a id="canonical-e4f099cdf5357581783640440865eeb6df098d859afdda428e7547a02779873d"></a>

## suffix_value property — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 808a10b5dc57 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-df9eb1a28922552ce0cbc3e941ecbe4111c8d9657150c7510b98f8889ddd7713"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.domain / 808a10b5dc57 / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c206e315bcd5af1541c637d035ba7ec96a8b10899be84bcb603c4a1447af6c1a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-37816b99f06bd20aa9de04dd8f21bba4447c40b817910c1a96331d7b6c2030e7"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 41fb0e839e3f / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8a9473a93a8f8218afddc636685505997cdb0a0a0fa8a2159405565215071628)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata

<a id="canonical-489e10f23815be8e3426f61ef7f50d075f975d68581654afe96d4f4bcffd6d80"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-def762bf321315c29eb09c9cf90d6490bb745251c5e853464e299a61a147816c"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 41fb0e839e3f / 3

<a id="canonical-42a7e5c07cf9599ec43234662fa0701c33e222c30028daab2969885ba89db380"></a>

<a id="canonical-49f7bd2716f9d7f96bf6867286278a5e49656cc62f131ac4cd9ad241839c6e0c"></a>

## description_spec property — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 41fb0e839e3f / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-dff6adb3ed153356eba944904aec8f898966a76d431da2653c5dcf3ed19620d7"></a>

<a id="canonical-d9515fcf967fca496ef8a4db502f48d0eac940fc550262a5a771952c04fab65a"></a>

## name property — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 41fb0e839e3f / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-23606f1f05e06967043a6b6ca92e03391e5a30278b2c700c1058d3e39428ca7b"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.metadata / 41fb0e839e3f / 6

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3378d3f112c6c7776f14117ce951ed99731e51a07dc56a9db0d615f6fe8575f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7efc866925ae48bc376ac7cc464f5a79fe0618d7ca7c73753644f4a37f6dde13"></a>

## bot_defense.policy.js_insert_all_pages_except.exclude_list.path — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 2dde0ccb3f2b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8a9473a93a8f8218afddc636685505997cdb0a0a0fa8a2159405565215071628)
- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4)
- bot_defense.policy.js_insert_all_pages_except.exclude_list.path

<a id="canonical-7e6a0ec0ba9887807969546c520be5d52d7d91a660e5ea7fb6baf39a46e14b08"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-fad7e4ac6359e2d2a67f9c1c875442b470a414094ddb57100c7bd41db3eeb2a5"></a>

## Direct properties — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 2dde0ccb3f2b / 3

<a id="canonical-750d60538195af08ccf6d5abd623be371dec7e5f7ecdf79c3cedc13e39a581bb"></a>

<a id="canonical-a76c7173e1e9402c532fe918294d141dd01d9120bc353666efa73622ac3d6d28"></a>

## path property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 2dde0ccb3f2b / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-b02e0fb1fb2e8f0f91e37d31646936705c59c515b299195da639f7eba41821b2"></a>

<a id="canonical-db4ff1ec66dd2caac74c9290bbec6b57bf674788617758ebf03c9b696ab96118"></a>

## prefix property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 2dde0ccb3f2b / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-628b98d9cec3fa3e499339905d133fd01ab3e4e4c4c39b2f11a6c522fd9c2ef3"></a>

<a id="canonical-25d67396980581ebe4d95123f3a0b68aaa3c2b4c1e4fd52341e7381244b52e32"></a>

## regex property — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 2dde0ccb3f2b / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-470e47f26e65db6ce1431cb2670f8837042235c25250a543ad24ce2e69f1ac30"></a>

## Next pages — bot_defense.policy.js_insert_all_pages_except.exclude_list.path / 2dde0ccb3f2b / 7

- [bot_defense.policy.js_insert_all_pages_except.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-7ef50a69e88e7e4e43226a98640042d6964a3a973650410260f743fd715a91b4)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f12725bca3159818a5f9cd0e04fc375f6e4926d251057d846a655093407caafa"></a>

## bot_defense.policy.js_insertion_rules — bot_defense.policy.js_insertion_rules / a2ebae2a85d6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- bot_defense.policy.js_insertion_rules

<a id="canonical-8d22507f11356fd012f4156f790489cba602ad1ae15032fa12b5a89410706f38"></a>

Type: `"single"`. Computed.

Defines custom JavaScript insertion rules for Bot Defense Policy.

Upstream description:

This defines custom JavaScript insertion rules for Bot Defense Policy.

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

<a id="canonical-0679bafeda6f7487718471e82d78c662a52e02c78473673a9622f83c75c49da0"></a>

## Direct properties — bot_defense.policy.js_insertion_rules / a2ebae2a85d6 / 3

- [exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9): complete subsection reference.

- [rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23): complete subsection reference.

<a id="canonical-7ef47e147366b36efbc67941672699d053a7897db9c1fd94450e39938096a1ed"></a>

## Next pages — bot_defense.policy.js_insertion_rules / a2ebae2a85d6 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-067868d1207c69b769e339ae93b197700ef3cca2d9365e12dc1eedda81a3786b"></a>

## bot_defense.policy.js_insertion_rules.exclude_list — bot_defense.policy.js_insertion_rules.exclude_list / 3efcaf733291 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- bot_defense.policy.js_insertion_rules.exclude_list

<a id="canonical-39d30138931de9830787fc5d66a97dbcb7f35581655160fc9d51f161894e0066"></a>

Type: `"list"`. Computed.

Optional JavaScript insertions exclude list of domain and path matchers.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1b8b8e7a18bb2b66abaeba1a753a5938f0cd892302f67b791e98f07e9fc7ba4b"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list / 3efcaf733291 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-73130ed0e2297986368584c883cfbf7227bae9dafc28adb8d960c5b34ed725e3): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-61cff5e39aad0f82ed89466d256ea0c75569f7eb1eccb125286b386d31afcf4e): complete subsection reference.

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-846cfafd88e939fcb7bafc6c84385b60b818bab71eaacd584296735681a59830): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a9759d0999a3626a930f3b924b38edca39d4c10b40fb20043dad5fb59147a944): complete subsection reference.

<a id="canonical-5f5398bf8205c34e1cd888fe2f6faabf98b82949e30d900e4c571839c20e515e"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list / 3efcaf733291 / 4

- [bot_defense.policy.js_insertion_rules.exclude_list.any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-73130ed0e2297986368584c883cfbf7227bae9dafc28adb8d960c5b34ed725e3)
- [bot_defense.policy.js_insertion_rules.exclude_list.domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-61cff5e39aad0f82ed89466d256ea0c75569f7eb1eccb125286b386d31afcf4e)
- [bot_defense.policy.js_insertion_rules.exclude_list.metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-846cfafd88e939fcb7bafc6c84385b60b818bab71eaacd584296735681a59830)
- [bot_defense.policy.js_insertion_rules.exclude_list.path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a9759d0999a3626a930f3b924b38edca39d4c10b40fb20043dad5fb59147a944)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-73130ed0e2297986368584c883cfbf7227bae9dafc28adb8d960c5b34ed725e3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb0a51890472e73f49032202c38b858443c0c63b1c5787bf8038fd1efc2f4edc"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.any_domain — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / 4000ff2220ef / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9)
- bot_defense.policy.js_insertion_rules.exclude_list.any_domain

<a id="canonical-13c4c043564e537994725fed7b94796a9a0fb93dcd7ed6d656431bef88394b98"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-e7209eb1f0b7d53d6742ccf21bb88c197ff52b2b718c68bbde9e49d7968f5d23"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / 4000ff2220ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8c39dd26df4b7489ffea7b686af564d8560953d509cea446d77f9cf11b130952"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.any_domain / 4000ff2220ef / 4

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-61cff5e39aad0f82ed89466d256ea0c75569f7eb1eccb125286b386d31afcf4e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-058b9be169bdcdda46afb44d45dd64a08548dec75696ec2fd098c118213e30a2"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.domain — bot_defense.policy.js_insertion_rules.exclude_list.domain / eef56197f4d7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9)
- bot_defense.policy.js_insertion_rules.exclude_list.domain

<a id="canonical-6746789889bf8b414163675f839f1c50b76fd46fc64d26dbaaecbc8af76a1362"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-262a502f74731d8556624d267f0c926dcf2da716521e546646bd90b7d3251682"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.domain / eef56197f4d7 / 3

<a id="canonical-f1dae1c1a9a994f7efdedeb0c40ea05bfa95ae2416630a603df9d7772757e855"></a>

<a id="canonical-d30f57babbad1e3e8347c9b9d69684bcfd3fd3d76dc0e84a6a03790dc38a9774"></a>

## exact_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / eef56197f4d7 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-c47a7279e8038e911677cdee567c39fb89de9577208bffdc1a983354b5ec4308"></a>

<a id="canonical-a27a55fe3b0158b50629aa82e611b90c8f3d3125deb9958ac19eef43b8a5963e"></a>

## regex_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / eef56197f4d7 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-b149d2be37fc6ef99b8e00e1abf56d1ef729c35c5579fda13f4a2fc678abc454"></a>

<a id="canonical-45ccccca4f8e719b8a5edb73ff0589c3aaed47ea16a78adef8e640fcd80ec06c"></a>

## suffix_value property — bot_defense.policy.js_insertion_rules.exclude_list.domain / eef56197f4d7 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-14f05fe30d114d461e2532e6718e679f1059287f64386ee1bdedfb26e5687439"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.domain / eef56197f4d7 / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-846cfafd88e939fcb7bafc6c84385b60b818bab71eaacd584296735681a59830"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c6c2de04432687f24a927bad68240ed6a08c248dcffb01cbbb8b4d25780b5e84"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.metadata — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 78b03a7cbbb7 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9)
- bot_defense.policy.js_insertion_rules.exclude_list.metadata

<a id="canonical-52df5461554bd1021bb12a2048268ebcc2e9830a71debc9cffe2b490ed18c296"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-87d84d192a8de96320b435173fe9fadc98868e4d2c918c8687d259c94f0295f3"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 78b03a7cbbb7 / 3

<a id="canonical-2281c7ed1ccf4df9c159e5737c48714918b8cdc047f6155cf7cf2a23e95a708b"></a>

<a id="canonical-79195927db67f257670a03b272cb3d9cfc8ea45ea0cf1cd3a880724bc78dd14c"></a>

## description_spec property — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 78b03a7cbbb7 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-620437636d1afe45b3aedbe3502946b958d028627e422c178d8c1284956cfcdc"></a>

<a id="canonical-e288050edda0e47031f88b93267f5c95c2827fe4cdb40c87706f33840fef4c96"></a>

## name property — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 78b03a7cbbb7 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-06721e5d5407697356ef73a1c853060d48182a8b88f06818b1f52baa59009b7b"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.metadata / 78b03a7cbbb7 / 6

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a9759d0999a3626a930f3b924b38edca39d4c10b40fb20043dad5fb59147a944"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b20ea09fe0e4311dbb6a3ee9c2e9047bc0c35fa2de1ef9d097c2592656ecab6c"></a>

## bot_defense.policy.js_insertion_rules.exclude_list.path — bot_defense.policy.js_insertion_rules.exclude_list.path / b084e4dbbfbc / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9)
- bot_defense.policy.js_insertion_rules.exclude_list.path

<a id="canonical-d95bb3f092248504244379f45dbe704185572cea9125e7c4a66194a230c426c9"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-50c1987b546c32d692cd413d612971094fda302be052dbc7e5a51bf3f8ea89e4"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.exclude_list.path / b084e4dbbfbc / 3

<a id="canonical-09b243f4d8c6e2ab5b884d98dfc707ade6931199c36d9816fce49be998e2a060"></a>

<a id="canonical-18729b18fd4ea6f888ceb3c1863bacefb4b17312c66afe7f2116dd889cfef28e"></a>

## path property — bot_defense.policy.js_insertion_rules.exclude_list.path / b084e4dbbfbc / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-49e1e5d63f4329137db31dc8a3bd6b8761a4e4d6df794699f6e77062f3280bcb"></a>

<a id="canonical-925d4f701a69dabe4993b27eb13547692a2c90c65f1a2ee4aa411154153ecc1f"></a>

## prefix property — bot_defense.policy.js_insertion_rules.exclude_list.path / b084e4dbbfbc / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-9b1e9d1cbabed84f033b2be3027bc4890b7b298f41766f9b601ccbb639474689"></a>

<a id="canonical-7daf9e2b099afef8c357eb7cf5600ba194891d2a9af9f809be9c97c85cdd1963"></a>

## regex property — bot_defense.policy.js_insertion_rules.exclude_list.path / b084e4dbbfbc / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-9d8e52057c4e13cf4fad21ce2850ae5be5560736145248f5ca636ec5914450ee"></a>

## Next pages — bot_defense.policy.js_insertion_rules.exclude_list.path / b084e4dbbfbc / 7

- [bot_defense.policy.js_insertion_rules.exclude_list](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0dd30628ee6ced18fabe1afd13754a55c3bd943cd9dffe7cc3d1122d039fcdf9)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ef7b9406655dd81dc2ccfa23ed4b225d2f46c28af5d161cdc6d62428ee8a1aab"></a>

## bot_defense.policy.js_insertion_rules.rules — bot_defense.policy.js_insertion_rules.rules / cb67433872e8 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- bot_defense.policy.js_insertion_rules.rules

<a id="canonical-15064984131b9e01263eeb6c6c3fbc131101cfe40d6c2de1ac9e206be8763e30"></a>

Type: `"list"`. Computed.

Required list of pages to insert Bot Defense client JavaScript.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-42ae27882b7943e62cff50fa4584375bebca13b4ab2af7f1c778bee136a4a07d"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules / cb67433872e8 / 3

- [any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-cfa58b0b8338e528997c9b6a93dad7b33063294ac38858e145eb274fd7f471df): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-d8d3ecac8a4506914baf5d5eed17b890bc7fc834fa10894502e92f303af042ab): complete subsection reference.

<a id="canonical-ae89e1051bd47fd6784c1dbddea2d53174474cb8dccd88810c1826507f495850"></a>

<a id="canonical-da7740e4ff261c85cc925638702c21828a9d87af42d47ae8f58f9c6f5ea9699e"></a>

## javascript_location property — bot_defense.policy.js_insertion_rules.rules / cb67433872e8 / 4

Type: `"string"`. Computed.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Upstream description:

All inside networks.

Insert JavaScript after &lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert
JavaScript before first tag.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-780898b370f368b0a744ec09bc57b714e84a4de38f9e086104de9c96cc386386): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8152dc5fef0b6e197b64d9675dc75b2fa8ee497802baad88588f748182984e60): complete subsection reference.

<a id="canonical-c3f7cececaef0bc84b8ed63f7c6699a809e7cc180dfb008bb47d588f3ef805f7"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules / cb67433872e8 / 5

- [bot_defense.policy.js_insertion_rules.rules.any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-cfa58b0b8338e528997c9b6a93dad7b33063294ac38858e145eb274fd7f471df)
- [bot_defense.policy.js_insertion_rules.rules.domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-d8d3ecac8a4506914baf5d5eed17b890bc7fc834fa10894502e92f303af042ab)
- [bot_defense.policy.js_insertion_rules.rules.metadata](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-780898b370f368b0a744ec09bc57b714e84a4de38f9e086104de9c96cc386386)
- [bot_defense.policy.js_insertion_rules.rules.path](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-8152dc5fef0b6e197b64d9675dc75b2fa8ee497802baad88588f748182984e60)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-cfa58b0b8338e528997c9b6a93dad7b33063294ac38858e145eb274fd7f471df"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-197bbded6b844ae2800fb98a2b54fff25114554e09298efc6894ba9a6e998b0c"></a>

## bot_defense.policy.js_insertion_rules.rules.any_domain — bot_defense.policy.js_insertion_rules.rules.any_domain / 3fe1031b22bd / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23)
- bot_defense.policy.js_insertion_rules.rules.any_domain

<a id="canonical-9d9d688ccf5f38a8e672e0fc0e1d72870a2a9679b8fa5fda642fa7a3f430cf53"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-06f48f6d4ccbe46eafad9841cc1ef0bda9e74131ee3aced640518afc0978b098"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.any_domain / 3fe1031b22bd / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e6a3cedc907400e5d4c5a904bb1c7dcc16892bec1001dc6a86d850aafccae0a3"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.any_domain / 3fe1031b22bd / 4

- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-d8d3ecac8a4506914baf5d5eed17b890bc7fc834fa10894502e92f303af042ab"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1472757e10ce34b2fd8230df539ef95dcf0c3a7041a8331069fc64972933e4b3"></a>

## bot_defense.policy.js_insertion_rules.rules.domain — bot_defense.policy.js_insertion_rules.rules.domain / a9ac0261b8f6 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23)
- bot_defense.policy.js_insertion_rules.rules.domain

<a id="canonical-aa5447068d56ad9ebde3ad33aaa763ad1bd54e9228ef98301f953d46d19f9e59"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-30979d35c003a99fd396f64fc11ce8ad6d3915f39bd33c8c2a3a19f307d3dc64"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.domain / a9ac0261b8f6 / 3

<a id="canonical-d260f9d738fcf86a279ff5b659cc84084841f962b03f4c83bd3f5a2d2d29b425"></a>

<a id="canonical-9287411282b6ef9e281e50be028e1fd4b1cfd8469b812018604ad085d8c4961b"></a>

## exact_value property — bot_defense.policy.js_insertion_rules.rules.domain / a9ac0261b8f6 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-5d0997864716864635f4e87d1dc21e2f29f173834097460040e94be6703d791d"></a>

<a id="canonical-05220151282ea1569900ec7e4e11b17b162176cd491eea9a9edd63c91243cefd"></a>

## regex_value property — bot_defense.policy.js_insertion_rules.rules.domain / a9ac0261b8f6 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-ac6045da5a8f1969f57a0fc3fdf457305e3958866596314b62617dd2f63783c6"></a>

<a id="canonical-4bb5eb04c1e01a4e84c98a2a7d39a57e4cc5bb6c68c969ca0fda4c88d0fb33a9"></a>

## suffix_value property — bot_defense.policy.js_insertion_rules.rules.domain / a9ac0261b8f6 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-7671593886d3a6b8310b81c58f7a24ef1c6c00b8d8febb654488829134044563"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.domain / a9ac0261b8f6 / 7

- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-780898b370f368b0a744ec09bc57b714e84a4de38f9e086104de9c96cc386386"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5a1dc3b000fee7f9672c76415e8f60c2fd7a2f5abaeb93fc4d1020358db2c57"></a>

## bot_defense.policy.js_insertion_rules.rules.metadata — bot_defense.policy.js_insertion_rules.rules.metadata / 69403706a9aa / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23)
- bot_defense.policy.js_insertion_rules.rules.metadata

<a id="canonical-14edda2addd5d0b64b494e1c20d56acc433ad9ba1495c438a13e65a6ca0a5e8b"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-bef26beb9b597a360db667206cb9c1261a3cada508c8c29e5b663dc447521b1b"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.metadata / 69403706a9aa / 3

<a id="canonical-0c425c01639639cb022f8f49001d2d024d7eafecabf73d3505280cbdd7bd6964"></a>

<a id="canonical-c8f3f68d1f6e8d9341e3d4b094cb7504590095d360e7486743eb2624c33241b1"></a>

## description_spec property — bot_defense.policy.js_insertion_rules.rules.metadata / 69403706a9aa / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-569adcf12578ea5cb72264700d74151d366cb715ad03d8b9169eb048d049f586"></a>

<a id="canonical-f6218f1ba1db6a3930e108b990308f3ce3caf04556778550b00feb01baa7fa22"></a>

## name property — bot_defense.policy.js_insertion_rules.rules.metadata / 69403706a9aa / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

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

<a id="canonical-b90ac8e9aee0814dd3bd3e1f1ca09d637300aad228c48e84eb82421c4ac9a6b8"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.metadata / 69403706a9aa / 6

- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-8152dc5fef0b6e197b64d9675dc75b2fa8ee497802baad88588f748182984e60"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c78b967ef7530fc4c0dda58020e80a17c835cb7ebd177af2d4811933ca23e22f"></a>

## bot_defense.policy.js_insertion_rules.rules.path — bot_defense.policy.js_insertion_rules.rules.path / e06000a0dc33 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-27fbd8e37173de924d1c2b72a8766317dcf6fb27a223e53f47a8aea6e95cf87d)
- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23)
- bot_defense.policy.js_insertion_rules.rules.path

<a id="canonical-024a51b0fd1110d95c6a072edaecac75fa60fcc3499dec8fb56f53dab50d3ee1"></a>

Type: `"single"`. Computed.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

<a id="canonical-86bc4d4c3e7700d6a381bf54ee996ee6a8592752ce2d8ff25686831ecd3a53c3"></a>

## Direct properties — bot_defense.policy.js_insertion_rules.rules.path / e06000a0dc33 / 3

<a id="canonical-a6d6a6948ba1c4f0186735925e0b4fa2109ed78b31817f2ec8924e539dc8ae1b"></a>

<a id="canonical-7b341a124cf4b30d6d3f4fce1660f87b6f9d93c06082fd7b9892c26adbd3e4be"></a>

## path property — bot_defense.policy.js_insertion_rules.rules.path / e06000a0dc33 / 4

Type: `"string"`. Computed.

Exclusive with \[prefix regex\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regex\] Exact path value to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-be3179c4637643026c95d95e24172fceff3b22df90c36c85a27869c82ac6566c"></a>

<a id="canonical-6744cc14a143b6f0412bf20325e2cde6e897d0cd16891a5c08224d8492c9cb38"></a>

## prefix property — bot_defense.policy.js_insertion_rules.rules.path / e06000a0dc33 / 5

Type: `"string"`. Computed.

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regex\] Path prefix to match (e.g. The value / will match on all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-5d35076bc35ccb0db38c9d05d27824bb288d12c07a08d1da3ba0ad92c5567696"></a>

<a id="canonical-4c66ba63f8792b9863753080d6ecfabc94ba93e5947ef492d441ae4b74975cdd"></a>

## regex property — bot_defense.policy.js_insertion_rules.rules.path / e06000a0dc33 / 6

Type: `"string"`. Computed.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-8869d3012cd46485fc1b91c133886a9101008dba270ff805865cda2dfbb92b12"></a>

## Next pages — bot_defense.policy.js_insertion_rules.rules.path / e06000a0dc33 / 7

- [bot_defense.policy.js_insertion_rules.rules](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-26385857c7a20fd608953bb8a0372a8e83291158b425466fc0c60566280ecb23)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-a6d501fcf23450d0d973e5b1aa3f37d111c82769e7e6e2f1ca962105445bc4f3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9342443a988902405b30ea65dd1997b56855827d492d202c9f41306dddbe9f92"></a>

## bot_defense.policy.mobile_sdk_config — bot_defense.policy.mobile_sdk_config / 3c368b4df628 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- bot_defense.policy.mobile_sdk_config

<a id="canonical-49778224ba1b6fa6ff1fa1dfaa4591e64759ee72fbff7e15baec62f140282c4e"></a>

Type: `"single"`. Computed.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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

<a id="canonical-dfff5fb2d505a9c3f34d00c3b429975b4495e9995b746039ace22c12c7fc6fd7"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config / 3c368b4df628 / 3

- [mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3a46abfc4bd2dafae14f198e085a18ab152415cf0babcc7671a9c6238472535a): complete subsection reference.

<a id="canonical-449534cfaeace0cee2d2b6940566b7c435c03bef2f2a972bb752ff6796d6cebb"></a>

## Next pages — bot_defense.policy.mobile_sdk_config / 3c368b4df628 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3a46abfc4bd2dafae14f198e085a18ab152415cf0babcc7671a9c6238472535a)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3a46abfc4bd2dafae14f198e085a18ab152415cf0babcc7671a9c6238472535a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a1a054ae502dd0b72770a3ea95d927c97c647a17c716693d24a5aa072079abe3"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier — bot_defense.policy.mobile_sdk_config.mobile_identifier / 5fd1aee3c6c9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a6d501fcf23450d0d973e5b1aa3f37d111c82769e7e6e2f1ca962105445bc4f3)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="canonical-db33b705faab1c439318b10de58df16c5dccb85339a00af9b04e65e2ddcc1222"></a>

Type: `"single"`. Computed.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

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

<a id="canonical-d64d850f39667dc3bcf747b4ae049a23239b8bf88cff12f085817fb41d70f8eb"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier / 5fd1aee3c6c9 / 3

- [headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-18a485b0041c8541ba00079349017fa63975464b40b55ab7230e325ce24ba51a): complete subsection reference.

<a id="canonical-892cb86a4f6c4e536ebb57c48628f171143e66aec92fa60158361156f606faca"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier / 5fd1aee3c6c9 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-18a485b0041c8541ba00079349017fa63975464b40b55ab7230e325ce24ba51a)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a6d501fcf23450d0d973e5b1aa3f37d111c82769e7e6e2f1ca962105445bc4f3)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-18a485b0041c8541ba00079349017fa63975464b40b55ab7230e325ce24ba51a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-348e5e1905bb6481061db27960b844d3e505ea0ab8152d2ab03a1a966e871272"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / ae1f011e70c9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a6d501fcf23450d0d973e5b1aa3f37d111c82769e7e6e2f1ca962105445bc4f3)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3a46abfc4bd2dafae14f198e085a18ab152415cf0babcc7671a9c6238472535a)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers

<a id="canonical-7f4d3c070d5de1c8fb899e2108ccaf76ca0d9ef540404982e73b94e8b6dc11ca"></a>

Type: `"list"`. Computed.

Headers that can be used to identify mobile traffic.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-ed7be32272dc9bd3d4a35b6d743dece12bbcb7090cbcee79f5a8461482834ef0"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / ae1f011e70c9 / 3

- [check_not_present](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-ac6fd40696dc6c2f53e6ab9c580aa2be125ae1ed18887c5e286403b500ed0e9c): complete subsection reference.

- [check_present](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-47de8921b0bbcd067e99bb6cb728944154b7da90b008ceedfced42304df9d82d): complete subsection reference.

- [item](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-24f5644f34bfed3977ae4c668d54748e726694311416538706dbc7073f2554e2): complete subsection reference.

<a id="canonical-c2137c92139693dbe91533d5724f916eee486428d1e94e16fd295a42076eb118"></a>

<a id="canonical-db050f5bffbeb993462403f77685321b02acc01d163f81ed01b252d7c4b38d04"></a>

## name property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / ae1f011e70c9 / 4

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-6e15c68e8bcec159bb092f86d32d459fd77bcc9a9779f385ac93f4acbe3392dd"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers / ae1f011e70c9 / 5

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-ac6fd40696dc6c2f53e6ab9c580aa2be125ae1ed18887c5e286403b500ed0e9c)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-47de8921b0bbcd067e99bb6cb728944154b7da90b008ceedfced42304df9d82d)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-24f5644f34bfed3977ae4c668d54748e726694311416538706dbc7073f2554e2)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3a46abfc4bd2dafae14f198e085a18ab152415cf0babcc7671a9c6238472535a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ac6fd40696dc6c2f53e6ab9c580aa2be125ae1ed18887c5e286403b500ed0e9c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-602ee3f4cbeeb7c5249ef25143159b644c05d5212cf1903f28be84780981492b"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / 3cc999144f76 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a6d501fcf23450d0d973e5b1aa3f37d111c82769e7e6e2f1ca962105445bc4f3)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3a46abfc4bd2dafae14f198e085a18ab152415cf0babcc7671a9c6238472535a)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-18a485b0041c8541ba00079349017fa63975464b40b55ab7230e325ce24ba51a)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="canonical-c6f3b75b02c9dde0ed40399bc5dd47d30981f66e5606499426c507fc8e0a5629"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check not present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-c5a5de3533a7619d9a6f71be3437c49feba4af24edbe7a9e4cd93dd8a37c6636"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / 3cc999144f76 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d9b2b41beecec0a7ce696101baeb37097f7b70062ce02687b4fa4a417fba85e2"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present / 3cc999144f76 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-18a485b0041c8541ba00079349017fa63975464b40b55ab7230e325ce24ba51a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-47de8921b0bbcd067e99bb6cb728944154b7da90b008ceedfced42304df9d82d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-19befcea4d90249997d601cc590acf562bce6e2bf8d45fa4346e25532104ac38"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / ca5293171483 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a6d501fcf23450d0d973e5b1aa3f37d111c82769e7e6e2f1ca962105445bc4f3)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3a46abfc4bd2dafae14f198e085a18ab152415cf0babcc7671a9c6238472535a)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-18a485b0041c8541ba00079349017fa63975464b40b55ab7230e325ce24ba51a)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present

<a id="canonical-f23a2a36033a0d280a300886cfd0d305b191af6ef92c03cc7eb0241dfc560689"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for check present.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-645918ca68dfca725ceec71e991604f1575f6fe7ec9d6dcbeba5430437d786c3"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / ca5293171483 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-5d22fc04bf9587c8577b227b7c7aafa10c940923c1f2ede39da61dbe09f05cba"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_present / ca5293171483 / 4

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-18a485b0041c8541ba00079349017fa63975464b40b55ab7230e325ce24ba51a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-24f5644f34bfed3977ae4c668d54748e726694311416538706dbc7073f2554e2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d31b3c820caa86dff26f14274f5d1383e9c3407512619e8a498c2dac4cd55cf7"></a>

## bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / a6185f0fbdbb / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-a6d501fcf23450d0d973e5b1aa3f37d111c82769e7e6e2f1ca962105445bc4f3)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3a46abfc4bd2dafae14f198e085a18ab152415cf0babcc7671a9c6238472535a)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-18a485b0041c8541ba00079349017fa63975464b40b55ab7230e325ce24ba51a)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item

<a id="canonical-c9d2819affdcab5776dddd0489ba0c8937bdd2237b9952b737fbe6c851937b7d"></a>

Type: `"single"`. Computed.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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

<a id="canonical-8e7c7d044a30813dd17947424894e6ec7ed039f3ad0c8e621fdad60359dbd30c"></a>

## Direct properties — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / a6185f0fbdbb / 3

<a id="canonical-5ae9485becb40cc197c37db7413ad243be138eb7a5b39ba6b6b5871e4adf4fa7"></a>

<a id="canonical-fc6551d12732afe51a57790c26feab5623a2f4c468d78dbf35409e356e4de319"></a>

## exact_values property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / a6185f0fbdbb / 4

Type: `["list", "string"]`. Computed.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-f871da7550f5c0c7e78e8a3dc90a8563b039f0a6b484f635f43ab41d94216d68"></a>

<a id="canonical-35c6b81671296efa81f567638d859104678ddab9228931839b924d93d3bd9c86"></a>

## regex_values property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / a6185f0fbdbb / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-8c523c9458bf33c624660ae55de3293c5affc8ba7878cbe5ff72e06d27fd745e"></a>

<a id="canonical-9050310bbb344e1a266d410c86cf0e3bf4fc0da08046f02fe712df199dba7875"></a>

## transformers property — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / a6185f0fbdbb / 6

Type: `["list", "string"]`. Computed.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b2590f2e36e8091e9f6509e87e5a29a8e19c9555631c12f1ac652a013acdce01"></a>

## Next pages — bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.item / a6185f0fbdbb / 7

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-18a485b0041c8541ba00079349017fa63975464b40b55ab7230e325ce24ba51a)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d2b9ca5053ce5ed57a26f6e5fa221e660ddf1f2582c4fcac3c4a8f93d090104d"></a>

## bot_defense.policy.protected_app_endpoints — bot_defense.policy.protected_app_endpoints / e050f0a6318d / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- bot_defense.policy.protected_app_endpoints

<a id="canonical-d251a48daab6db7a1390563acbfeb856424c177c32d7f208ab4e0ef8b27a8b3a"></a>

Type: `"list"`. Computed.

List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32
endpoints per LB' after 4 LBs.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0c1607339657ff5d12c22287302a0aa9ebcfe505591fc98389f1f65d44e7e9bf"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints / e050f0a6318d / 3

- [allow_good_bots](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-78fa576449e762eadb908d1e07902dbab1f4677dddb980a55b170ae061621962): complete subsection reference.

- [any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3b357a6210248e2f56fcede333b85e2220b35e6eac6c4f5fa366e4672b164f99): complete subsection reference.

- [domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-896c2a4280d8ac1334b7283632266afeb91254882a690b126e7e325807c00808): complete subsection reference.

- [flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b): complete subsection reference.

- [headers](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1dc75be54f1cdd14964ea980f8f5f331ef6bcc97102c2edc2924da6407b58355): complete subsection reference.

<a id="canonical-f5650852512654e9aaf3aa73f981217afba1b369586bcca0b9c3ed4f5a6ad573"></a>

<a id="canonical-2385f368675983c4cb31895f9b1332bf35a2d9853608fae13962d2744fa13954"></a>

## http_methods property — bot_defense.policy.protected_app_endpoints / e050f0a6318d / 4

Type: `["list", "string"]`. Computed.

\[Enum:
METHOD\_ANY|METHOD\_GET|METHOD\_POST|METHOD\_PUT|METHOD\_PATCH|METHOD\_DELETE|METHOD\_GET\_DOCUMENT\]
HTTP Methods. List of HTTP methods. Possible values are \`METHOD\_ANY\`, \`METHOD\_GET\`,
\`METHOD\_POST\`, \`METHOD\_PUT\`, \`METHOD\_PATCH\`, \`METHOD\_DELETE\`, \`METHOD\_GET\_DOCUMENT\`.
Defaults to \`METHOD\_ANY\`.

Upstream description:

List of HTTP methods.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[0,1,3,4,10]",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [metadata](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-34ee1626a5c040af1deb14d877aca49ce0c86424b622a03d79976de29c7f96f2): complete subsection reference.

- [mitigate_good_bots](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-e8b7705fee7dbed497825f916860d2b3cf50706ff39430aa6b3aa0d679b28438): complete subsection reference.

- [mitigation](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-208af0e212f3bb964ecd00b829e4d20d7034b97ef7170aa7f5b69e817776ce56): complete subsection reference.

- [mobile](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-9dd309e9d42e0dc6cff1eaa4866fcb1eb77b732f837e22050e230f20a91d7e4a): complete subsection reference.

- [path](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-c7d7f2d98dd0f4ee9431998c2b48129f7de75411910865f5fde4344bad631be2): complete subsection reference.

<a id="canonical-9c4a76535ad3e9bfc53f4cd1b54911408438b8fe52b03409b434ee7723ea6740"></a>

<a id="canonical-854b4c19641b4487a58d5e4d7e25fbef48ed22efbbd7cb6039d6155e8687798d"></a>

## protocol property — bot_defense.policy.protected_app_endpoints / e050f0a6318d / 5

Type: `"string"`. Computed.

\[Enum: BOTH|HTTP|HTTPS\] SchemeType is used to indicate URL scheme. - BOTH: BOTH URL scheme for
HTTPS:// or HTTP://. - HTTP: HTTP URL scheme HTTP:// only. - HTTPS: HTTPS URL scheme HTTPS:// only.
Possible values are \`BOTH\`, \`HTTP\`, \`HTTPS\`. Defaults to \`BOTH\`.

Upstream description:

SchemeType is used to indicate URL scheme.

&#8203;- BOTH: BOTH

URL scheme for HTTPS:// or HTTP://. &#8203;- HTTP: HTTP

URL scheme HTTP:// only. &#8203;- HTTPS: HTTPS

URL scheme HTTPS:// only.

Receipt-pinned upstream constraints:

```json
{
  "default": "BOTH",
  "enum": [
    "BOTH",
    "HTTP",
    "HTTPS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [query_params](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-8f9ed574617e0603808d1f263733a7becf36d50fee87d9e57b097235d850ff20): complete subsection reference.

- [undefined_flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-430e5217d3dbdfa993dd379d57e9dc024d9bd69ac64e862dceaae59eb354a3bb): complete subsection reference.

- [web](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-aac9008059c1454fb30ebf7517eb42ca5bbee2e7d3bbbf0602e2c4dc55a33e0f): complete subsection reference.

- [web_mobile](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-a6bd53beba7104217652dde3d3aab8f07b7dde74f425fe8b193c95f66da684d1): complete subsection reference.

<a id="canonical-edcafd6039a9e54fcff3fc7fcbd78994e43cbe64957a35031889c5686664fe02"></a>

## Next pages — bot_defense.policy.protected_app_endpoints / e050f0a6318d / 6

- [bot_defense.policy.protected_app_endpoints.allow_good_bots](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-78fa576449e762eadb908d1e07902dbab1f4677dddb980a55b170ae061621962)
- [bot_defense.policy.protected_app_endpoints.any_domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3b357a6210248e2f56fcede333b85e2220b35e6eac6c4f5fa366e4672b164f99)
- [bot_defense.policy.protected_app_endpoints.domain](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-896c2a4280d8ac1334b7283632266afeb91254882a690b126e7e325807c00808)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.headers](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-1dc75be54f1cdd14964ea980f8f5f331ef6bcc97102c2edc2924da6407b58355)
- [bot_defense.policy.protected_app_endpoints.metadata](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-34ee1626a5c040af1deb14d877aca49ce0c86424b622a03d79976de29c7f96f2)
- [bot_defense.policy.protected_app_endpoints.mitigate_good_bots](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-e8b7705fee7dbed497825f916860d2b3cf50706ff39430aa6b3aa0d679b28438)
- [bot_defense.policy.protected_app_endpoints.mitigation](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-208af0e212f3bb964ecd00b829e4d20d7034b97ef7170aa7f5b69e817776ce56)
- [bot_defense.policy.protected_app_endpoints.mobile](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-9dd309e9d42e0dc6cff1eaa4866fcb1eb77b732f837e22050e230f20a91d7e4a)
- [bot_defense.policy.protected_app_endpoints.path](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-c7d7f2d98dd0f4ee9431998c2b48129f7de75411910865f5fde4344bad631be2)
- [bot_defense.policy.protected_app_endpoints.query_params](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-8f9ed574617e0603808d1f263733a7becf36d50fee87d9e57b097235d850ff20)
- [bot_defense.policy.protected_app_endpoints.undefined_flow_label](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-430e5217d3dbdfa993dd379d57e9dc024d9bd69ac64e862dceaae59eb354a3bb)
- [bot_defense.policy.protected_app_endpoints.web](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-aac9008059c1454fb30ebf7517eb42ca5bbee2e7d3bbbf0602e2c4dc55a33e0f)
- [bot_defense.policy.protected_app_endpoints.web_mobile](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-a6bd53beba7104217652dde3d3aab8f07b7dde74f425fe8b193c95f66da684d1)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-78fa576449e762eadb908d1e07902dbab1f4677dddb980a55b170ae061621962"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9da32db7196d9c85b8e085a32c4e97b47739168a5151d613630e10249610b984"></a>

## bot_defense.policy.protected_app_endpoints.allow_good_bots — bot_defense.policy.protected_app_endpoints.allow_good_bots / 513143d9ae44 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- bot_defense.policy.protected_app_endpoints.allow_good_bots

<a id="canonical-616b5995f37a3cca39e9b2ccac19d3029c08347f4fd336cffc7c0d42e489e556"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for allow good bots.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-25b94e55b64e5a39c920e1f7edbe7d33d7aa669ccc9c7a13926a2a7822215849"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.allow_good_bots / 513143d9ae44 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-20271001dc3de84c8b7bdb862742da187b6a0c17d5038de667b0117cbfc3438f"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.allow_good_bots / 513143d9ae44 / 4

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-3b357a6210248e2f56fcede333b85e2220b35e6eac6c4f5fa366e4672b164f99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25ec90dd4303fc2432583e8858067ee0f85514eee0bf1cf1bd42fa4eb5c31f63"></a>

## bot_defense.policy.protected_app_endpoints.any_domain — bot_defense.policy.protected_app_endpoints.any_domain / fd28c509b546 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- bot_defense.policy.protected_app_endpoints.any_domain

<a id="canonical-29b2e0a219357ac0d2fbd1a64f6f9d1982461264a2f4c28576c92be4d2204eee"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-29c39ff8fa11cadec9d00cb99bf8f50d8de97f252bec5ed55020e7e3597d463a"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.any_domain / fd28c509b546 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1f2dfed989ee27c8d0359bf34eb6694d574d12174812bac552d91a9ecc4a56a0"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.any_domain / fd28c509b546 / 4

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-896c2a4280d8ac1334b7283632266afeb91254882a690b126e7e325807c00808"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a42607a29b6ca56d3c41d510933a5b491478908c35642a883221e9f6effeabaf"></a>

## bot_defense.policy.protected_app_endpoints.domain — bot_defense.policy.protected_app_endpoints.domain / ba707b39b135 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- bot_defense.policy.protected_app_endpoints.domain

<a id="canonical-a0d4475c70d09a0072f4b4b4fea4f11c207cfed671567d4353f25b0f4b8f1277"></a>

Type: `"single"`. Computed.

Domain name for routing and identification.

Upstream description:

Domains names.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-domain_choice": "[\"exact_value\",\"regex_value\",\"suffix_value\"]"
}
```

<a id="canonical-038ebb18fe865f4777770bb91f8dc30a686e55a2e954a7493dcab566ce1338a4"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.domain / ba707b39b135 / 3

<a id="canonical-3c09dc908edcc97c9cc3a909081cad1b55bcc3fa4516becd70a7ea7df9921133"></a>

<a id="canonical-b5cdd61f35b760b8c70afc430fc874390867e88046f412d24027053de8316c4a"></a>

## exact_value property — bot_defense.policy.protected_app_endpoints.domain / ba707b39b135 / 4

Type: `"string"`. Computed.

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regex\_value suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-a4ab15170ebd200dffbece434fab52f90d571b907f818837b4b51d89c43a926c"></a>

<a id="canonical-ed5e7e70cd6ca59efedf8271b6f386e2e3cb8b2fc655065525aa98aaed9b6111"></a>

## regex_value property — bot_defense.policy.protected_app_endpoints.domain / ba707b39b135 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-ad95177b1edf72fcb9ab37072fe223b92f9bcfdc16e89ec57a180c1ff66ce671"></a>

<a id="canonical-f58c84b2026c5bdab926955418e6d2b4353dea52631f7a17cd83b0594092b090"></a>

## suffix_value property — bot_defense.policy.protected_app_endpoints.domain / ba707b39b135 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regex\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-4ef6ce6ce421d22cd998ca59f8285f89834d2f876b73d17730223dedc1dcbcd1"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.domain / ba707b39b135 / 7

- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c5f9ea6c6212b83e66d46bf0e42d2bcc21aecbf159770741c14933f2fe87bc61"></a>

## bot_defense.policy.protected_app_endpoints.flow_label — bot_defense.policy.protected_app_endpoints.flow_label / bd3fbc191ec1 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- bot_defense.policy.protected_app_endpoints.flow_label

<a id="canonical-d6f35852042b18cc24f8a3c31e6755dd7238fb4d503b884e3751fbebe995bbf8"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Category allows to associate traffic with selected category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-flow_label_choice": "[\"account_management\",\"authentication\",\"financial_services\",\"flight\",\"profile_management\",\"search\",\"shopping_gift_cards\"]"
}
```

<a id="canonical-2b7efcbb054c4112d968b054e1df7cd6cec2e6889743d65368bedfa5d49f9822"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label / bd3fbc191ec1 / 3

- [account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-c6b953f8ad5cae00c78dc957e52803169be304e99b2b530486800bb78fc8b1a3): complete subsection reference.

- [authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b): complete subsection reference.

- [financial_services](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-fca6fa937721cea5d501ad026efab4fedd74f171b2c575d6aeec7051baa0351f): complete subsection reference.

- [flight](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-599c029197bdb8f0e31d9a7c671fc9fc61db11691e0c58ebbf6a7b999457dc63): complete subsection reference.

- [profile_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2f7d496b97ac5ae9bc4bb293ca00d529de9c636f5b9579b2dc533afb10779eb4): complete subsection reference.

- [search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-67430b5b6a6f9ce11e7ec2f3bdd6397b18818c7c0232279f4f0f6cd6aa3f338a): complete subsection reference.

- [shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4ee3089e74f9ff898307ef6a6ab1a9d8d8150e1f285db46a416047f952b1691d): complete subsection reference.

<a id="canonical-117ac126044bc386579ab7724ebcd43f6049255c8dd05481b3d3d0e67585131e"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label / bd3fbc191ec1 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-c6b953f8ad5cae00c78dc957e52803169be304e99b2b530486800bb78fc8b1a3)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- [bot_defense.policy.protected_app_endpoints.flow_label.financial_services](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-fca6fa937721cea5d501ad026efab4fedd74f171b2c575d6aeec7051baa0351f)
- [bot_defense.policy.protected_app_endpoints.flow_label.flight](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-599c029197bdb8f0e31d9a7c671fc9fc61db11691e0c58ebbf6a7b999457dc63)
- [bot_defense.policy.protected_app_endpoints.flow_label.profile_management](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-2f7d496b97ac5ae9bc4bb293ca00d529de9c636f5b9579b2dc533afb10779eb4)
- [bot_defense.policy.protected_app_endpoints.flow_label.search](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-67430b5b6a6f9ce11e7ec2f3bdd6397b18818c7c0232279f4f0f6cd6aa3f338a)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-4ee3089e74f9ff898307ef6a6ab1a9d8d8150e1f285db46a416047f952b1691d)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-c6b953f8ad5cae00c78dc957e52803169be304e99b2b530486800bb78fc8b1a3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5ba3244df99bac1b1ee376ac9a88c47d64a5fccd850fca65832b59d888edd42a"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management — bot_defense.policy.protected_app_endpoints.flow_label.account_management / 8934d25055ff / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management

<a id="canonical-8843e807e0977db1c3c84ae595cb80b12700229d0e12ce862c4f4aa46b2a7aac"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Account Management Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"create\",\"password_reset\"]"
}
```

<a id="canonical-f383ab9ed13fd46d93ec7908ab3d14820a909c6ea862ee05f78023526ab62cb6"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management / 8934d25055ff / 3

- [create](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0d3d4eb4bcef18a4feaa897158309e230792ac359316afa0f66f0ec8a6525893): complete subsection reference.

- [password_reset](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-ee307f56c1d3eaba2d742ee3a88c1f2a4dc1aac7d0e76574cfe7ba8fb417f7aa): complete subsection reference.

<a id="canonical-f93f1812044c6edeb58363977c0d789a6d43715d7ef35c09a19b2008e4ccd5d8"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management / 8934d25055ff / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.create](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-0d3d4eb4bcef18a4feaa897158309e230792ac359316afa0f66f0ec8a6525893)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-ee307f56c1d3eaba2d742ee3a88c1f2a4dc1aac7d0e76574cfe7ba8fb417f7aa)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-0d3d4eb4bcef18a4feaa897158309e230792ac359316afa0f66f0ec8a6525893"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dfd132a506078a19ab8d8c012c5c18f1a413acc4ff24f6a2ce29c97da0e937b8"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.create — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 98a6db4ccd50 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-c6b953f8ad5cae00c78dc957e52803169be304e99b2b530486800bb78fc8b1a3)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.create

<a id="canonical-a90b558def0b3f6cd91c8cda5243da3e4d79d21c0f85143bae20849d21a656f6"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-253d3a72430b9ac2a6341d226cbc22594ea509b7c1ef1bee27906af928055588"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 98a6db4ccd50 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-b2b17ee5fa0466e836faad49759d3a2c535c6cc2e30c9a22f39b975f8f7c556e"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management.create / 98a6db4ccd50 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-c6b953f8ad5cae00c78dc957e52803169be304e99b2b530486800bb78fc8b1a3)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ee307f56c1d3eaba2d742ee3a88c1f2a4dc1aac7d0e76574cfe7ba8fb417f7aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97cd1e337a101367a7fee2a7a5c581aff62b88c863422384b8d62d65bdaef303"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / ff2d9bd1fed9 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-c6b953f8ad5cae00c78dc957e52803169be304e99b2b530486800bb78fc8b1a3)
- bot_defense.policy.protected_app_endpoints.flow_label.account_management.password_reset

<a id="canonical-0385d41f467410199ebc18b4ade65d242d778a2fd51552d17dd495163c972368"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for password reset.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-cd56e7fe726d033f6bf5164301daabd46eb32d0a5bac8319843bd48b5c7411bd"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / ff2d9bd1fed9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d67c53948c0868de550d30bf6cda9bd86efea9a88e6f8a854248557775125fbd"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.account_management.passwor / ff2d9bd1fed9 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.account_management](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-c6b953f8ad5cae00c78dc957e52803169be304e99b2b530486800bb78fc8b1a3)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3dfbfa172d99625fb9b9eb0d33b2472c00e463a17e9c9a6593d25a83c6f614de"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 22e72749077e / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication

<a id="canonical-216d3804a48b19b5b6f51e8ffcca398f9fdc72812db8eec5d4b2559df000b62e"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Authentication Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"login\",\"login_mfa\",\"login_partner\",\"logout\",\"token_refresh\"]"
}
```

<a id="canonical-90cae97e87ca4f3b7d61f98a1f8e8e95010e0be453b5d8de3b5576f107ab28ac"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 22e72749077e / 3

- [login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-dd5cca1a22f44ec19aea136ab7717c84b8578474577a4a7c642c1dd44d7748a8): complete subsection reference.

- [login_mfa](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-baabdc52c1d5b1f42fc40e72f1eb5d5b47ee8bc8fff09c4147719cad280b992b): complete subsection reference.

- [login_partner](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-f9e26cd39c537a153b929a6703a6b58fb9a91fedce446d2cee51ebe6c41d26c9): complete subsection reference.

- [logout](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-b7bb71295854a49e9abf973614308f317ed7802bf5264e93405912102899edc8): complete subsection reference.

- [token_refresh](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-b6f98cc4285ebea36f6adc3a06aabc967ccac7990d530f322deb5b03c836c053): complete subsection reference.

<a id="canonical-77f05493823557f30550615ea9fa24bc3801306a3016988984b258c992deeca3"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication / 22e72749077e / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-dd5cca1a22f44ec19aea136ab7717c84b8578474577a4a7c642c1dd44d7748a8)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-baabdc52c1d5b1f42fc40e72f1eb5d5b47ee8bc8fff09c4147719cad280b992b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-f9e26cd39c537a153b929a6703a6b58fb9a91fedce446d2cee51ebe6c41d26c9)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.logout](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-b7bb71295854a49e9abf973614308f317ed7802bf5264e93405912102899edc8)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.token_refresh](data-sources--cdn_loadbalancer--reference--group-008.md#canonical-b6f98cc4285ebea36f6adc3a06aabc967ccac7990d530f322deb5b03c836c053)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-dd5cca1a22f44ec19aea136ab7717c84b8578474577a4a7c642c1dd44d7748a8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-786b9455f75ffb79db48e1d0be3422f24b60aaa0b927a8e5f04bd1e596f0999e"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / f76b7c879061 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login

<a id="canonical-053e1e4b3722b388cbf5b64ab8d55810a79d5264cd14bf4ea2f51e1ab2fc7ea3"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result. Bot Defense Transaction Result.

Upstream description:

Bot Defense Transaction Result.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-transaction_result_choice": "[\"disable_transaction_result\",\"transaction_result\"]"
}
```

<a id="canonical-d684fd1b7bb69dc117284316e1c45da639cee65ca1e5a66c4c6f18ad42aa7eb9"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / f76b7c879061 / 3

- [disable_transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-ef49159ca909f061d82e052eb487d2e51de7eabba93dbd5bdaf987b205a1578d): complete subsection reference.

- [transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-00a706bfd10d4686bdd7d0ceff3d11aae27a27656d89803ab297bd27699362f7): complete subsection reference.

<a id="canonical-282e003108dfe74dde1faa93f8fa495b7bd658d41310d62a18fc6e6a85ba7c5e"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login / f76b7c879061 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-ef49159ca909f061d82e052eb487d2e51de7eabba93dbd5bdaf987b205a1578d)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-00a706bfd10d4686bdd7d0ceff3d11aae27a27656d89803ab297bd27699362f7)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-ef49159ca909f061d82e052eb487d2e51de7eabba93dbd5bdaf987b205a1578d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5a2bf5c87ae81e4fe0a9c9b8a9fc40912ecbdcbc9b27c138d9aa770ddd0a266f"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disab / 23f500075789 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-dd5cca1a22f44ec19aea136ab7717c84b8578474577a4a7c642c1dd44d7748a8)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disable_transaction_result

<a id="canonical-b1a6ef4a4ee5d7da61dd2a89ecd02b9c34b57b790975811f5bc87d6164c7b088"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-dac00764a37ebf4d60983a26066f5e3a860e57dd052495df52897ad8d3202f0e"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disab / 23f500075789 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-359e6236893efc3057b46996cccb0644915150974efd25a8d63778eba6f27a7c"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.disab / 23f500075789 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-dd5cca1a22f44ec19aea136ab7717c84b8578474577a4a7c642c1dd44d7748a8)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-00a706bfd10d4686bdd7d0ceff3d11aae27a27656d89803ab297bd27699362f7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0d914549208205da4db3e77df77dd57508d26cebacaf3310fd03e8d9fee54cf"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / a140fda04a08 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-dd5cca1a22f44ec19aea136ab7717c84b8578474577a4a7c642c1dd44d7748a8)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result

<a id="canonical-703e649938ab82d7b6dc157f0bfaf84ac28577ee3db8dfc51d4d057335d9798b"></a>

Type: `"single"`. Computed.

Bot Defense Transaction Result Type. Bot Defense Transaction ResultType.

Upstream description:

Bot Defense Transaction ResultType.

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

<a id="canonical-6376f335157f4c7bb5dc810f0a8c0cf0958b428f509156619030d7c5d5b3f183"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / a140fda04a08 / 3

- [failure_conditions](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-fffb7bc355473ac8b21f6ac6fbd794eba53510b853374b94150571fdfaeb48c2): complete subsection reference.

- [success_conditions](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-fac9b7d8ac38a440aab0feaf36d44e921c8f6737d731b6a6de84f0697a8088a0): complete subsection reference.

<a id="canonical-dba5bb10e2ad1dbf55342f50b2b9a9588bcb16611d0d13ad3d678541ac73d217"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / a140fda04a08 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-fffb7bc355473ac8b21f6ac6fbd794eba53510b853374b94150571fdfaeb48c2)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-fac9b7d8ac38a440aab0feaf36d44e921c8f6737d731b6a6de84f0697a8088a0)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-dd5cca1a22f44ec19aea136ab7717c84b8578474577a4a7c642c1dd44d7748a8)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-fffb7bc355473ac8b21f6ac6fbd794eba53510b853374b94150571fdfaeb48c2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6976e40d7a9f531f55d5ccf1e3be5d8a6082cb0b21b426947ba07029ef856a92"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 2dc671cd7a3b / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-dd5cca1a22f44ec19aea136ab7717c84b8578474577a4a7c642c1dd44d7748a8)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-00a706bfd10d4686bdd7d0ceff3d11aae27a27656d89803ab297bd27699362f7)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.failure_conditions

<a id="canonical-ef5c27dd7aa49fdd4925d97f5d43ff28c57ece962d88c9742531c7783ba481ce"></a>

Type: `"list"`. Computed.

Failure Conditions. Failure Conditions.

Upstream description:

Failure Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-dc97fe0eac467aa9d7e3a7096f6b3a7cb74b2f317a66011d0a52244ec3427e14"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 2dc671cd7a3b / 3

<a id="canonical-90a06aa77cba1d044bde8ca970cb47a4077b3f75a1bd8d164d99867c4a4c1ed5"></a>

<a id="canonical-a161dedf34774fb3a766afedbf78b1d38e06dc559481f96a80e9c05e33f41fae"></a>

## name property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 2dc671cd7a3b / 4

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-7325f87392a0681aef5a9ae6bb9a6d807d44b6e1affd45956c37616440100045"></a>

<a id="canonical-e971cb80b7e54653b20a3e637c895d8b5f0b75b09c4388193846fb218e6bc485"></a>

## regex_values property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 2dc671cd7a3b / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-aaa3f2b50f408ecd6840bfb1c31044593bfa6d9e4c29043bdb96c1ea86f2ee45"></a>

<a id="canonical-54667a101b822fbc41904672bf8d8544f61e3afb9876c3485568d14392f26222"></a>

## status property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 2dc671cd7a3b / 6

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-fd9c52e8b8c46597eaf3b4ea493289d2218fefc3df35ab7d6b3b1293d6d802a3"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 2dc671cd7a3b / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-00a706bfd10d4686bdd7d0ceff3d11aae27a27656d89803ab297bd27699362f7)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-fac9b7d8ac38a440aab0feaf36d44e921c8f6737d731b6a6de84f0697a8088a0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c4e5cb0ff5671edae5a6c08c879183449722adb95c576d11a29c0777174ba5d8"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 9e398a30801c / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-dd5cca1a22f44ec19aea136ab7717c84b8578474577a4a7c642c1dd44d7748a8)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-00a706bfd10d4686bdd7d0ceff3d11aae27a27656d89803ab297bd27699362f7)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result.success_conditions

<a id="canonical-409f4a8e575367db83416eef04c5ccd5060618c993a0f5b2ebee664c6086e453"></a>

Type: `"list"`. Computed.

Success Conditions. Success Conditions.

Upstream description:

Success Conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
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
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-cdd22db9a60e183cb845c6a304a8afc1abf9722f8fe2f8c1750e9446c4f24873"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 9e398a30801c / 3

<a id="canonical-dc73add5c49d1c7c939ae36b03470c1791b0a0355140c7887eeebe561a90cdf2"></a>

<a id="canonical-a6059b169ca74e3b2353548741beb5c0cde715e349ea6e5e5401f16cff23b259"></a>

## name property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 9e398a30801c / 4

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-af06500342558cbe9f772540686c3b8102f7712b49011aafdabbef47c8bb229d"></a>

<a id="canonical-1e37746b73e4a696b4e9c90456e484c9402e2aca1f202ba9e7bd0a8bc8e92bb9"></a>

## regex_values property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 9e398a30801c / 5

Type: `["list", "string"]`. Computed.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-015b4aaca089035fdf3fcae515bce73fbe227d6886dfb4c1dabadcfd4b520f8f"></a>

<a id="canonical-5feb532bcd0938a66e1b7450d67f27c05767ceb43f6d63b9aa77dcbd980964d1"></a>

## status property — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 9e398a30801c / 6

Type: `"string"`. Computed.

\[Enum:
EmptyStatusCode|Continue|OK|Created|Accepted|NonAuthoritativeInformation|NoContent|ResetContent|PartialContent|MultiStatus|AlreadyReported|IMUsed|MultipleChoices|MovedPermanently|Found|SeeOther|NotModified|UseProxy|TemporaryRedirect|PermanentRedirect|BadRequest|Unauthorized|PaymentRequired|Forbidden|NotFound|MethodNotAllowed|NotAcceptable|ProxyAuthenticationRequired|RequestTimeout|Conflict|Gone|LengthRequired|PreconditionFailed|PayloadTooLarge|URITooLong|UnsupportedMediaType|RangeNotSatisfiable|ExpectationFailed|MisdirectedRequest|UnprocessableEntity|Locked|FailedDependency|UpgradeRequired|PreconditionRequired|TooManyRequests|RequestHeaderFieldsTooLarge|InternalServerError|NotImplemented|BadGateway|ServiceUnavailable|GatewayTimeout|HTTPVersionNotSupported|VariantAlsoNegotiates|InsufficientStorage|LoopDetected|NotExtended|NetworkAuthenticationRequired\]
HTTP response status codes EmptyStatusCode response codes means it is not specified Continue status
code OK status code Created status code Accepted status code Non Authoritative Information status
code No Content status code Reset Content status code Partial Content status code Multi Status..
Possible values are \`EmptyStatusCode\`, \`Continue\`, \`OK\`, \`Created\`, \`Accepted\`,
\`NonAuthoritativeInformation\`, \`NoContent\`, \`ResetContent\`, \`PartialContent\`,
\`MultiStatus\`, \`AlreadyReported\`, \`IMUsed\`, \`MultipleChoices\`, \`MovedPermanently\`,
\`Found\`, \`SeeOther\`, \`NotModified\`, \`UseProxy\`, \`TemporaryRedirect\`,
\`PermanentRedirect\`, \`BadRequest\`, \`Unauthorized\`, \`PaymentRequired\`, \`Forbidden\`,
\`NotFound\`, \`MethodNotAllowed\`, \`NotAcceptable\`, \`ProxyAuthenticationRequired\`,
\`RequestTimeout\`, \`Conflict\`, \`Gone\`, \`LengthRequired\`, \`PreconditionFailed\`,
\`PayloadTooLarge\`, \`URITooLong\`, \`UnsupportedMediaType\`, \`RangeNotSatisfiable\`,
\`ExpectationFailed\`, \`MisdirectedRequest\`, \`UnprocessableEntity\`, \`Locked\`,
\`FailedDependency\`, \`UpgradeRequired\`, \`PreconditionRequired\`, \`TooManyRequests\`,
\`RequestHeaderFieldsTooLarge\`, \`InternalServerError\`, \`NotImplemented\`, \`BadGateway\`,
\`ServiceUnavailable\`, \`GatewayTimeout\`, \`HTTPVersionNotSupported\`, \`VariantAlsoNegotiates\`,
\`InsufficientStorage\`, \`LoopDetected\`, \`NotExtended\`, \`NetworkAuthenticationRequired\`.
Defaults to \`EmptyStatusCode\`.

Upstream description:

HTTP response status codes

EmptyStatusCode response codes means it is not specified Continue status code OK status code Created
status code Accepted status code Non Authoritative Information status code No Content status code
Reset Content status code Partial Content status code Multi Status status code Already Reported
status code Im Used status code Multiple Choices status code Moved Permanently status code Found
status code See Other status code Not Modified status code Use Proxy status code Temporary Redirect
status code Permanent Redirect status code Bad Request status code Unauthorized status code Payment
Required status code Forbidden status code Not Found status code Method Not Allowed status code Not
Acceptable status code Proxy Authentication Required status code Request Timeout status code
Conflict status code Gone status code Length Required status code Precondition Failed status code
Payload Too Large status code URI Too Long status code Unsupported Media Type status code Range Not
Satisfiable status code Expectation Failed status code Misdirected Request status code Unprocessable
Entity status code Locked status code Failed Dependency status code Upgrade Required status code
Precondition Required status code Too Many Requests status code Request Header Fields Too Large
status code Internal Server Error status code Not Implemented status code Bad Gateway status code
Service Unavailable status code Gateway Timeout status code HTTP Version Not Supported status code
Variant Also Negotiates status code Insufficient Storage status code Loop Detected status code Not
Extended status code Network Authentication Required status code.

Receipt-pinned upstream constraints:

```json
{
  "default": "EmptyStatusCode",
  "enum": [
    "EmptyStatusCode",
    "Continue",
    "OK",
    "Created",
    "Accepted",
    "NonAuthoritativeInformation",
    "NoContent",
    "ResetContent",
    "PartialContent",
    "MultiStatus",
    "AlreadyReported",
    "IMUsed",
    "MultipleChoices",
    "MovedPermanently",
    "Found",
    "SeeOther",
    "NotModified",
    "UseProxy",
    "TemporaryRedirect",
    "PermanentRedirect",
    "BadRequest",
    "Unauthorized",
    "PaymentRequired",
    "Forbidden",
    "NotFound",
    "MethodNotAllowed",
    "NotAcceptable",
    "ProxyAuthenticationRequired",
    "RequestTimeout",
    "Conflict",
    "Gone",
    "LengthRequired",
    "PreconditionFailed",
    "PayloadTooLarge",
    "URITooLong",
    "UnsupportedMediaType",
    "RangeNotSatisfiable",
    "ExpectationFailed",
    "MisdirectedRequest",
    "UnprocessableEntity",
    "Locked",
    "FailedDependency",
    "UpgradeRequired",
    "PreconditionRequired",
    "TooManyRequests",
    "RequestHeaderFieldsTooLarge",
    "InternalServerError",
    "NotImplemented",
    "BadGateway",
    "ServiceUnavailable",
    "GatewayTimeout",
    "HTTPVersionNotSupported",
    "VariantAlsoNegotiates",
    "InsufficientStorage",
    "LoopDetected",
    "NotExtended",
    "NetworkAuthenticationRequired"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-07396143dcc13a5dbf80e393b0b24901ed40c17ac1040278c848daa1b817bce8"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.trans / 9e398a30801c / 7

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication.login.transaction_result](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-00a706bfd10d4686bdd7d0ceff3d11aae27a27656d89803ab297bd27699362f7)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-baabdc52c1d5b1f42fc40e72f1eb5d5b47ee8bc8fff09c4147719cad280b992b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-abd757608ed19beee38437e7448271247933087fd6a3de3b64e2951de2e6dfc1"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa / 8ae06e163a40 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa

<a id="canonical-8fcfa828f606bae791172233b132cd7e6df1c63bf02f02bcca39c8edc98aa839"></a>

Type: `["object", {}]`. Computed.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-22c83e89effcf1737a04d8e4cba5d74ee27a537fe9bc754b4230bdcea57fd171"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa / 8ae06e163a40 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-74779653e05352aef455ced255faba0dcecbc31bca45d4801ec0f0762a68e0cc"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_mfa / 8ae06e163a40 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-f9e26cd39c537a153b929a6703a6b58fb9a91fedce446d2cee51ebe6c41d26c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-df790918bdc3954cdfe317c6182efbee62f3dfda7a48493c9473a1b0c40428ab"></a>

## bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partn / e0fcf2a30634 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)
- [Property reference](data-sources--cdn_loadbalancer--reference--group-001.md#canonical-5db8d4b5757acfc29311ca7ae528f64484791ce86769b3c942b2e0b2a6218c6b)
- [bot_defense](data-sources--cdn_loadbalancer--reference--group-006.md#canonical-80c830837258747b14c29edc361da71ddaace2df95d3141cd7182f22161b76e0)
- [bot_defense.policy](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-815a14c0461913304dcf64135e0a57386d7e3574b225352489744a1d18614d55)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-3f4ceda9c727384e58e6e476d7b5e1a9934f3e9367ba87405eedc60d3d54f85b)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-4802025440d9dc323d0579980497f1538ee24874fbd845a954f1f209e38d265b)
- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partner

<a id="canonical-694d58ea68f2b2fb26c9ffe1788af584f2f56f52ab6ad497aba839ec9d608685"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for login partner.

Upstream description:

This can be used for messages where no values are needed.

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

<a id="canonical-54d45c11d106adc3008229435bbffb5f947eb907a1cd4edc21aef23177c639ed"></a>

## Direct properties — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partn / e0fcf2a30634 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-78c5b43efd98f63331e7e86f0fa178a813ce8cbbacfab4ae515adf354f411da5"></a>

## Next pages — bot_defense.policy.protected_app_endpoints.flow_label.authentication.login_partn / e0fcf2a30634 / 4

- [bot_defense.policy.protected_app_endpoints.flow_label.authentication](data-sources--cdn_loadbalancer--reference--group-007.md#canonical-1f72995f704a663ee1ce7cebcda8c7c88c67b90e1f62139912468ed99763814b)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md#canonical-1d377b95e8f5bf8636f98b5c9b8e1221f466f7bdda35a73abf8605ed5b3e8eea)

<a id="canonical-b7bb71295854a49e9abf973614308f317ed7802bf5264e93405912102899edc8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
