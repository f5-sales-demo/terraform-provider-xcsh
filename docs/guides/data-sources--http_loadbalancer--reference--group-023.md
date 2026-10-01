---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-72547ae6f78f65ca6c53c469cbe06e3e1a4e0d69bf6f5ef46cef5850dcb5ac09"></a>

## port property — routes.direct_response_route.incoming_port / 4dc940d7bce2 / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-915bb45197d7c82becb85848ba2a8ecc4b82f0bd1d010fcecb4bd8fa5083d316"></a>

<a id="canonical-0396a2ec45633f0864e94d5eb0c3b1000ca3294a409536378e3df29116341597"></a>

## port_ranges property — routes.direct_response_route.incoming_port / 4dc940d7bce2 / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-4b566231f4dffe0e66a9b8a07a160a972e17031de32aa0ce96f4298e20bf0bbc"></a>

## Next pages — routes.direct_response_route.incoming_port / 4dc940d7bce2 / 6

- [routes.direct_response_route.incoming_port.no_port_match](data-sources--http_loadbalancer--reference--group-023.md#canonical-3867f520a0380b2671da101048d3cf9916f3a6325803c683088abdec7a0de256)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-3867f520a0380b2671da101048d3cf9916f3a6325803c683088abdec7a0de256"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-41bb58cb4c2424ff84b69db397423c21c17b00f5a2c7d54f69e2b1bab658a6a7"></a>

## routes.direct_response_route.incoming_port.no_port_match — routes.direct_response_route.incoming_port.no_port_match / d100cbb24f8d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- [routes.direct_response_route.incoming_port](data-sources--http_loadbalancer--reference--group-022.md#canonical-5b994a44d4478081c4a6f627b3c96d54caa1db45d8d57058460cf3bb7daea7cf)
- routes.direct_response_route.incoming_port.no_port_match

<a id="canonical-6c8e5da329ae8a5790b3c00a21eb7662b24ef7ca12ce0eb0f25df0e12c5316bf"></a>

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

<a id="canonical-91904ae5fb8104d947dcfb9cf8caf5c83b38cdc6ba90b65cee768fa227c7eefa"></a>

## Direct properties — routes.direct_response_route.incoming_port.no_port_match / d100cbb24f8d / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-228eb53661c5c9fe3376aebc5ebc9b19376a9f1e68c429fc93744b65db3ed6fa"></a>

## Next pages — routes.direct_response_route.incoming_port.no_port_match / d100cbb24f8d / 4

- [routes.direct_response_route.incoming_port](data-sources--http_loadbalancer--reference--group-022.md#canonical-5b994a44d4478081c4a6f627b3c96d54caa1db45d8d57058460cf3bb7daea7cf)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-21be880e7260de75735f117af48f72d37af63bb1cce4eeddc20d5c9f05e80180"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7cba49da06371cb417771ab80c120d337563d07d8ad60fdea0cdf8a84e5a5354"></a>

## routes.direct_response_route.path — routes.direct_response_route.path / ba45c7dc5949 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- routes.direct_response_route.path

<a id="canonical-d0028cd8039051201d3c5384cdbcafc122ca65cbb11d66b767addd43fbc57433"></a>

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

<a id="canonical-606647da53c431332ced4146579b00cdeb1ed2f060eb2f3f4a3850a6c4d0502e"></a>

## Direct properties — routes.direct_response_route.path / ba45c7dc5949 / 3

<a id="canonical-8f145a2eff43a40b9dc08b6f7340bb1240e5633f39804915b5183f2caef3916f"></a>

<a id="canonical-7b29c89e8101178ee5c6867c96c9f1a11806e2b51df1c83d90ab0df04f57d14a"></a>

## path property — routes.direct_response_route.path / ba45c7dc5949 / 4

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

<a id="canonical-647765caaa047832d9e8f88c518a3ccb7c81dcdb5f8be29f9b19aeba58a7e43e"></a>

<a id="canonical-e301b2609ad62dbf03059ca4a6cbbd7ee7f0cd22bafdf82a92ee00ba2f6e6782"></a>

## prefix property — routes.direct_response_route.path / ba45c7dc5949 / 5

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

<a id="canonical-69e5bc89455ed4f13ee11a5614823d68881e8520f306c202be2def5bf67c772a"></a>

<a id="canonical-26f16f6b8600108e553124e84492760008bcff2a405bfe9f5b43cf771b0c3cf7"></a>

## regex property — routes.direct_response_route.path / ba45c7dc5949 / 6

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

<a id="canonical-722748a8e528bebe246cc276146f483ac5c8fd3fb94283e33b62ffef4fc04e63"></a>

## Next pages — routes.direct_response_route.path / ba45c7dc5949 / 7

- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-bdea64573060609049924cbdc60e279435e0ad3cfe1b6b6cfee5a37d53c29fc9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2dceaebd25ee2eadddfc78ea03734f042ceb715c3a7e6ac7d05f0a57bf48a21"></a>

## routes.direct_response_route.route_direct_response — routes.direct_response_route.route_direct_response / d977aa3174e3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- routes.direct_response_route.route_direct_response

<a id="canonical-ff5a8446c86a4375e121cf85eee07959cb455bbef27dc8aa8c7d339169a0fb2b"></a>

Type: `"single"`. Computed.

Send this direct response in case of route match action is direct response.

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

<a id="canonical-19f686140f0dde789559ee302f56250b9684a9cbfb3e503745c4f4a0e142f419"></a>

## Direct properties — routes.direct_response_route.route_direct_response / d977aa3174e3 / 3

<a id="canonical-0eba67f0a8a5748f1d9d0ab11905802983d18a380cf73835766423392ee6c199"></a>

<a id="canonical-6a2e10a235eb0fa386cadc2f7c6133c5a4c3e0b73beaffca00777bd335ef3f41"></a>

## response_body_encoded property — routes.direct_response_route.route_direct_response / d977aa3174e3 / 4

Type: `"string"`. Computed.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML.

Upstream description:

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in Base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". Base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 65536
    },
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
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-9b340050d93acf6583832570f8f469748b2edbd4119d2f2758c5b5a790cf2027"></a>

<a id="canonical-6c153161d6fe74cafc317dbb6a9beddfdcea96f1d0c50ecf3220a6c31137fdf9"></a>

## response_code property — routes.direct_response_route.route_direct_response / d977aa3174e3 / 5

Type: `"number"`. Computed.

Response Code. Response code to send.

Upstream description:

Response code to send.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 100
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "100",
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

<a id="canonical-36912588de515d579c3c404cc25aa4085dd943019af669f955cac25a42e5f140"></a>

## Next pages — routes.direct_response_route.route_direct_response / d977aa3174e3 / 6

- [routes.direct_response_route](data-sources--http_loadbalancer--reference--group-022.md#canonical-241694bc7b22e71a0e4b05da09ef2aaf3109f204b89ccd6b2d0d4b7eb4e2876e)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e6d8eed85ba395546a2a9d8218a7b745369555f5cec0dc2345b226d95c1b76bd"></a>

## routes.redirect_route — routes.redirect_route / 726b00165bd0 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- routes.redirect_route

<a id="canonical-54a8ab4afcd0acf2f687063310b47486c7dff5291e9766d7e76f657dc4d79f90"></a>

Type: `"single"`. Computed.

Redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects the
matching traffic to a different URL.

Upstream description:

A redirect route matches on path, incoming header, incoming port and/or HTTP method and redirects
the matching traffic to a different URL.

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

<a id="canonical-aadc9affd776de5e6baf5a4040aae8e6e779b01b03f4507c4f7fc1b007d16ea3"></a>

## Direct properties — routes.redirect_route / 726b00165bd0 / 3

- [headers](data-sources--http_loadbalancer--reference--group-023.md#canonical-6b47c6301e43271ca4a724ed92951ea7bb3ccb83c67a91a41a3184160f09c74a): complete subsection reference.

<a id="canonical-412224927d6f8a9c1df476d27e8831d52f7bff30472333fdc64d0fba99c03183"></a>

<a id="canonical-52dfbd84b65727ba8befc2ab7a8c2651aced0561192fa37d4ff4f23e2d31a0e5"></a>

## http_method property — routes.redirect_route / 726b00165bd0 / 4

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--http_loadbalancer--reference--group-023.md#canonical-15e1a0dd024da20cafb621e953bef3d481274bedbfd71c7bce2a7494784c5ed6): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-023.md#canonical-e869e4352b001069f47465c8f63ca7463cf63deb307d8bb4d0eae437c0da5e58): complete subsection reference.

- [route_redirect](data-sources--http_loadbalancer--reference--group-023.md#canonical-d4d8723e3b526dc9f70ed1a836f9645a1a4e136fc306a7d27581a02cd4dd2e6b): complete subsection reference.

<a id="canonical-355750e4c5dcf82adc72f03c13f8e289b0394d87f008ff231a8fdcb416979ce2"></a>

## Next pages — routes.redirect_route / 726b00165bd0 / 5

- [routes.redirect_route.headers](data-sources--http_loadbalancer--reference--group-023.md#canonical-6b47c6301e43271ca4a724ed92951ea7bb3ccb83c67a91a41a3184160f09c74a)
- [routes.redirect_route.incoming_port](data-sources--http_loadbalancer--reference--group-023.md#canonical-15e1a0dd024da20cafb621e953bef3d481274bedbfd71c7bce2a7494784c5ed6)
- [routes.redirect_route.path](data-sources--http_loadbalancer--reference--group-023.md#canonical-e869e4352b001069f47465c8f63ca7463cf63deb307d8bb4d0eae437c0da5e58)
- [routes.redirect_route.route_redirect](data-sources--http_loadbalancer--reference--group-023.md#canonical-d4d8723e3b526dc9f70ed1a836f9645a1a4e136fc306a7d27581a02cd4dd2e6b)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6b47c6301e43271ca4a724ed92951ea7bb3ccb83c67a91a41a3184160f09c74a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6633884e5ce475cc2c0d8bd7b5a2281371badf1420a5602a3c29b7523b11b297"></a>

## routes.redirect_route.headers — routes.redirect_route.headers / 160e029d4ee1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- routes.redirect_route.headers

<a id="canonical-595c2721aecb17a70a15ff218c97e84e502b767465211e3ee1409da331a760d1"></a>

Type: `"list"`. Computed.

Headers. List of (key, value) headers.

Upstream description:

List of (key, value) headers.

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-482c9fedf9ce655ecf1830e001d48fa8966e572b616dfb23f49a14d6032e459e"></a>

## Direct properties — routes.redirect_route.headers / 160e029d4ee1 / 3

<a id="canonical-4e636f95115d7fd72378a6e03d2b67cb5d0f9bfc1b1f61001b74fe2a114fe07c"></a>

<a id="canonical-2ae174f9b1044ae4c04f51050928e589735a881f8c19287e703e70187bd8d224"></a>

## exact property — routes.redirect_route.headers / 160e029d4ee1 / 4

Type: `"string"`. Computed.

Exclusive with \[presence regex\] Header value to match exactly.

Upstream description:

Exclusive with \[presence regex\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-10716407a0c44c8631f9a32fd29f37135fd2b324431ce3af808c4bd140de906f"></a>

<a id="canonical-c3f499d234cf4066644ec844b79f4b3d6d87ca8cdefaa750d02a2f76af83f972"></a>

## invert_match property — routes.redirect_route.headers / 160e029d4ee1 / 5

Type: `"bool"`. Computed.

Invert the result of the match to detect missing header or non-matching value.

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

<a id="canonical-a58f0fd3785415fc7dda7dd2cfce77763c013cf4e33dd7d59bf87fb85a7c17d7"></a>

<a id="canonical-92112595bc10e4d1bc255877015b1b2ff9f35e0a58c92a74355fac84a53f8c2e"></a>

## name property — routes.redirect_route.headers / 160e029d4ee1 / 6

Type: `"string"`. Computed.

Name. Name of the header.

Upstream description:

Name of the header.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-e821c0b0969245e8d7aeacf784b63f8f20f8ea4d622631f7c2b6c7691dadb897"></a>

<a id="canonical-d2048c94951e01825dfe3cf3d0c254d152f757bc624b27ef9db9dbb391a93048"></a>

## presence property — routes.redirect_route.headers / 160e029d4ee1 / 7

Type: `"bool"`. Computed.

Exclusive with \[exact regex\] If true, check for presence of header.

Upstream description:

Exclusive with \[exact regex\] If true, check for presence of header.

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

<a id="canonical-eafe2b0161fe2289459f058bbd015bba19b7085417c2ee9649ed51e9fdc9b618"></a>

<a id="canonical-a6ebdc08e01ec33d3357760822ed74f9e5ec573982023fff08d25f06e7dddc81"></a>

## regex property — routes.redirect_route.headers / 160e029d4ee1 / 8

Type: `"string"`. Computed.

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Upstream description:

Exclusive with \[exact presence\] Regex match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-6d0c6226eb45bf1041a276f2c89b366bff6e0ef9105baaa813fa9e2efacc0e1b"></a>

## Next pages — routes.redirect_route.headers / 160e029d4ee1 / 9

- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-15e1a0dd024da20cafb621e953bef3d481274bedbfd71c7bce2a7494784c5ed6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ca767bf885d76ae5ba99c6b519daf978bae23eaa55e195e5fded29f31c7f4d1a"></a>

## routes.redirect_route.incoming_port — routes.redirect_route.incoming_port / ffc5bdb5ecbc / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- routes.redirect_route.incoming_port

<a id="canonical-f9d0f01a30667bbb1ba12bedffd7f337ef6754c60495a96efd7b638a1a7f7423"></a>

Type: `"single"`. Computed.

Port match of the request can be a range or a specific port.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_match": "[\"no_port_match\",\"port\",\"port_ranges\"]"
}
```

<a id="canonical-85934be36b49e85076f9811c12cca1ada46182149ce1910ffd0378a9f87c6dca"></a>

## Direct properties — routes.redirect_route.incoming_port / ffc5bdb5ecbc / 3

- [no_port_match](data-sources--http_loadbalancer--reference--group-023.md#canonical-f1147e6bf6c465ce85170e8ffa7ecd9cb192c344f6cdc945148603dbd13514ca): complete subsection reference.

<a id="canonical-2f67657494fa319e8ef14b4e3342d72c1763bb410778bc0f7967fe6bf60f83bd"></a>

<a id="canonical-33ce89919dc7ff5fae5309b09b03eeeb3a59b8d00bb8fad861796b35453f6f0e"></a>

## port property — routes.redirect_route.incoming_port / ffc5bdb5ecbc / 4

Type: `"number"`. Computed.

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Upstream description:

Exclusive with \[no\_port\_match port\_ranges\] Exact Port to match.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-024102b0e2681ddab6c27c444e269d37d2c2179c5676c46814af5c2a868968d8"></a>

<a id="canonical-ab53cb92d102bcfaf54abbc62968696bfe7a7be7f997b6d99661b644cb31f394"></a>

## port_ranges property — routes.redirect_route.incoming_port / ffc5bdb5ecbc / 5

Type: `"string"`. Computed.

Exclusive with \[no\_port\_match port\] Port range to match.

Upstream description:

Exclusive with \[no\_port\_match port\] Port range to match.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 32,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 32,
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
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "32",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range": "true"
  }
}
```

<a id="canonical-7a4882c0d0e283dcac1d2df66f0bdeb6a48e74fd63b27e2fd9e76809a1eb3fd4"></a>

## Next pages — routes.redirect_route.incoming_port / ffc5bdb5ecbc / 6

- [routes.redirect_route.incoming_port.no_port_match](data-sources--http_loadbalancer--reference--group-023.md#canonical-f1147e6bf6c465ce85170e8ffa7ecd9cb192c344f6cdc945148603dbd13514ca)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-f1147e6bf6c465ce85170e8ffa7ecd9cb192c344f6cdc945148603dbd13514ca"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8e5e9627148e07d0fbc0ffd298b379c673c004a3433ccc52ea781bb10be0fd17"></a>

## routes.redirect_route.incoming_port.no_port_match — routes.redirect_route.incoming_port.no_port_match / 8123ae968539 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- [routes.redirect_route.incoming_port](data-sources--http_loadbalancer--reference--group-023.md#canonical-15e1a0dd024da20cafb621e953bef3d481274bedbfd71c7bce2a7494784c5ed6)
- routes.redirect_route.incoming_port.no_port_match

<a id="canonical-aae4d3d1e5e2ccfdea597bede704866691f810d797e3f1e1de06d278e76f965b"></a>

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

<a id="canonical-f32a6bd4ceed8796723d9fe2ecd7aecc1743544fb1aa3514892976bface13544"></a>

## Direct properties — routes.redirect_route.incoming_port.no_port_match / 8123ae968539 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-25fa8fef1cfb86ba7836851016f88d40caf7728713385fcb5cbeadb2d902db5f"></a>

## Next pages — routes.redirect_route.incoming_port.no_port_match / 8123ae968539 / 4

- [routes.redirect_route.incoming_port](data-sources--http_loadbalancer--reference--group-023.md#canonical-15e1a0dd024da20cafb621e953bef3d481274bedbfd71c7bce2a7494784c5ed6)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e869e4352b001069f47465c8f63ca7463cf63deb307d8bb4d0eae437c0da5e58"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-04a1e6d219ce85feeb31e2e3b58ee9aa97219ebbbad3dace99713457e4b766a2"></a>

## routes.redirect_route.path — routes.redirect_route.path / 1ad0aee5dec3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- routes.redirect_route.path

<a id="canonical-a4bfef0e20d6ce0e2444cb5df2704c113ba1a820a74216b48d3fae44409af13e"></a>

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

<a id="canonical-110513d704beb772a0b460a66640c5937ea52573385243d3089984f4799dea3f"></a>

## Direct properties — routes.redirect_route.path / 1ad0aee5dec3 / 3

<a id="canonical-d453f97c8c10c8f653b5010b4ad77794224dfbbe4ef15292f00b3304f8cfa2db"></a>

<a id="canonical-ae766fc4fdbca265b6a14bb8f02de10c2f2ba8b476b7101606bccf3f8b82a384"></a>

## path property — routes.redirect_route.path / 1ad0aee5dec3 / 4

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

<a id="canonical-012940f715aef5cb44f73a668681398db021f5fb2db21be465dced5f44318d81"></a>

<a id="canonical-2d9688b31865059c78dd9e1b103c2f04c2aff4c8372eeaf7114b954faf5d538e"></a>

## prefix property — routes.redirect_route.path / 1ad0aee5dec3 / 5

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

<a id="canonical-d8db817abd1c4bf86ca4e032f9c83f6b16196576255539ea6aa8cb03285972b9"></a>

<a id="canonical-a655982e15d7c6006912b6efbd116f6512e5bf58a401ce6a6cb2d0f8ebf37ee6"></a>

## regex property — routes.redirect_route.path / 1ad0aee5dec3 / 6

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

<a id="canonical-32780c41bad2524b3fd34e4fe13a7c915605168a22cce5bd989a66d31ccc1fc8"></a>

## Next pages — routes.redirect_route.path / 1ad0aee5dec3 / 7

- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d4d8723e3b526dc9f70ed1a836f9645a1a4e136fc306a7d27581a02cd4dd2e6b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-cca63663f808a8a3c41901e692f766fe525852d59877b419fe52c887cd5565ed"></a>

## routes.redirect_route.route_redirect — routes.redirect_route.route_redirect / 18232ee54f2a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- routes.redirect_route.route_redirect

<a id="canonical-c6797b5f9d8f99c8e104f9b9efa622f536445b4578aa1455706234f3ba3c51a6"></a>

Type: `"single"`. Computed.

Route redirect parameters when match action is redirect.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

<a id="canonical-140588e7c1283288e7cfd7f01fca7c1032a40f6c2af89ed5452adbe22bd642f1"></a>

## Direct properties — routes.redirect_route.route_redirect / 18232ee54f2a / 3

<a id="canonical-14b06d62bc3fc62e43c411dde174a009a8a6ab81faa39212fdfdcae758a3c98c"></a>

<a id="canonical-6c02e55316d03b40a806a13ec3d2c9b2ab21f2df530f352e7ba54cd794058cf2"></a>

## host_redirect property — routes.redirect_route.route_redirect / 18232ee54f2a / 4

Type: `"string"`. Computed.

Swap host part of incoming URL in redirect URL.

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
  }
}
```

<a id="canonical-6204acf62d5de252bf8fa7e42287d9d3846fe5199361dbc4ac92438870f5c164"></a>

<a id="canonical-e02ad6da131661a6a44a0d66489c3f380c91be2a23a94816761428ba5c82ae84"></a>

## path_redirect property — routes.redirect_route.route_redirect / 18232ee54f2a / 5

Type: `"string"`. Computed.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Upstream description:

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

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

<a id="canonical-5a0c7718411a10854b63700648b4ff9b3b821beed3e4a0f530650f0b03a65aec"></a>

<a id="canonical-355e8efe44ea10efad932544981f52c6c367c8224963f93ecf8a24307b5364e5"></a>

## prefix_rewrite property — routes.redirect_route.route_redirect / 18232ee54f2a / 6

Type: `"string"`. Computed.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Upstream description:

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

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

<a id="canonical-fe5330ff452274f6bf1f2081b3ddf9751aad252e4669cf34e8827651f3e12d61"></a>

<a id="canonical-fd8716fd3905d7896736efe24dd2b7339f28eaab72eb4369653d9946fa27d4ab"></a>

## proto_redirect property — routes.redirect_route.route_redirect / 18232ee54f2a / 7

Type: `"string"`. Computed.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Upstream description:

Swap protocol part of incoming URL in redirect URL The protocol can be swapped with either HTTP or
HTTPS When incoming-proto option is specified, swapping of protocol is not done.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "incoming-proto",
    "http",
    "https"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"incoming-proto\\\",\\\"http\\\",\\\"https\\\"]"
  }
}
```

- [remove_all_params](data-sources--http_loadbalancer--reference--group-023.md#canonical-8f2ceef6856f2ac3ba1c183245d4b6539e12c37af34d817e0b441566ca156a21): complete subsection reference.

<a id="canonical-a4f65ade36c5e0d33351e3ac5655e35484ccfb5a9f5aa9345a0d6f1b829cd76e"></a>

<a id="canonical-ee6716fefcc6d5b229071e139c4d6e543f84a6f1460db8ed8610a2b17def0fe7"></a>

## replace_params property — routes.redirect_route.route_redirect / 18232ee54f2a / 8

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1a07c3837c9e96630d6bcd860ea8f0425003298cd868f502ad61fe6248a959c1"></a>

<a id="canonical-c1f3cebf94cfd468ea969d72045af28e9207ca045bb8b45c6e35716cde206708"></a>

## response_code property — routes.redirect_route.route_redirect / 18232ee54f2a / 9

Type: `"number"`. Computed.

The HTTP status code to use in the redirect response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](data-sources--http_loadbalancer--reference--group-023.md#canonical-6e31736265dc182a40622a7b3058aeda9653fc575132fa498f8628e27b33e82c): complete subsection reference.

<a id="canonical-75f2a6278b80910734156b11a4d70f16185c249b33d6dea6c89e6393d706545a"></a>

## Next pages — routes.redirect_route.route_redirect / 18232ee54f2a / 10

- [routes.redirect_route.route_redirect.remove_all_params](data-sources--http_loadbalancer--reference--group-023.md#canonical-8f2ceef6856f2ac3ba1c183245d4b6539e12c37af34d817e0b441566ca156a21)
- [routes.redirect_route.route_redirect.retain_all_params](data-sources--http_loadbalancer--reference--group-023.md#canonical-6e31736265dc182a40622a7b3058aeda9653fc575132fa498f8628e27b33e82c)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-8f2ceef6856f2ac3ba1c183245d4b6539e12c37af34d817e0b441566ca156a21"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-82c0f3b3fe074789de80047e9692939c78a954745a5c81c0dfbe28b644cf3d3c"></a>

## routes.redirect_route.route_redirect.remove_all_params — routes.redirect_route.route_redirect.remove_all_params / 0bea0f077d79 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- [routes.redirect_route.route_redirect](data-sources--http_loadbalancer--reference--group-023.md#canonical-d4d8723e3b526dc9f70ed1a836f9645a1a4e136fc306a7d27581a02cd4dd2e6b)
- routes.redirect_route.route_redirect.remove_all_params

<a id="canonical-6d990014dad02e62c3bdd5c34991be14ac6f909dad973632a7ffd471d7bb7181"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-98d55adda693b212b083ba272b7a702c11524e8d9b43076062bcd68fcc3535ac"></a>

## Direct properties — routes.redirect_route.route_redirect.remove_all_params / 0bea0f077d79 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0a0c20875f0233d214647bf233ae33bb6c46bbad8078006fac85ae8b9177816c"></a>

## Next pages — routes.redirect_route.route_redirect.remove_all_params / 0bea0f077d79 / 4

- [routes.redirect_route.route_redirect](data-sources--http_loadbalancer--reference--group-023.md#canonical-d4d8723e3b526dc9f70ed1a836f9645a1a4e136fc306a7d27581a02cd4dd2e6b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6e31736265dc182a40622a7b3058aeda9653fc575132fa498f8628e27b33e82c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8416e1e24172d49888b7ce9e46833879ae897ec633279e7b9ff5b0a4ae523cc2"></a>

## routes.redirect_route.route_redirect.retain_all_params — routes.redirect_route.route_redirect.retain_all_params / da3e09faffe9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.redirect_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-cd706f2aae1efbae824fd3e0a5070d8ff662967e7b86604b485aa5d3b39b0ffe)
- [routes.redirect_route.route_redirect](data-sources--http_loadbalancer--reference--group-023.md#canonical-d4d8723e3b526dc9f70ed1a836f9645a1a4e136fc306a7d27581a02cd4dd2e6b)
- routes.redirect_route.route_redirect.retain_all_params

<a id="canonical-880278e73a1c6e52bcad74b908d819d23a5f8227202574797235e2ebb47b5047"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-3d9fbb1bc6f3d068903185b6be147c9c084aaf6eb97dc1b8fb52a8fde9dfd5f2"></a>

## Direct properties — routes.redirect_route.route_redirect.retain_all_params / da3e09faffe9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-18bf969e3aae5a8148e0c9d37d1c65dca4722d315c94fbcc0dad29cb0731f716"></a>

## Next pages — routes.redirect_route.route_redirect.retain_all_params / da3e09faffe9 / 4

- [routes.redirect_route.route_redirect](data-sources--http_loadbalancer--reference--group-023.md#canonical-d4d8723e3b526dc9f70ed1a836f9645a1a4e136fc306a7d27581a02cd4dd2e6b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-26a3458c018c5878797838910c1bbb3aac76cbcc3d77c81d2434c1334aa80760"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dd0af3e5eee60a61739c377d8d6ba2e280776fe85a2561d0471af6812be8a517"></a>

## routes.route_state_disabled — routes.route_state_disabled / 8f53ea62bc90 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- routes.route_state_disabled

<a id="canonical-af87e7c14af30bd2eb0e52920dbf8d052817f9e4aa5da46a88add3a950458645"></a>

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

<a id="canonical-53da75c8288a103ad8ebd2ed78e323a52c588ddf5d701ef8184e8839fbee55ac"></a>

## Direct properties — routes.route_state_disabled / 8f53ea62bc90 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3553c16d96daff0fff489535008d8036f9aefb5ee446b193f37305f2d6debd4a"></a>

## Next pages — routes.route_state_disabled / 8f53ea62bc90 / 4

- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c40cbd077ed664a4f6b634a5b335155e6e239298163d67a59031a515d5ba9eb3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e895fc497eea20ea423816936ec0c105f41df949e1f7fcdc793c0eac6b86c332"></a>

## routes.route_state_enabled — routes.route_state_enabled / 754073b31b4e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- routes.route_state_enabled

<a id="canonical-fa853fea18a23f5d9f6174751b4ade880f8659adb356f23ab8fd5f25dd40375c"></a>

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

<a id="canonical-ba1dac4cc9f982777aa894fdaeba88ed96f99bd2e331f4f76d55ef279cea44bf"></a>

## Direct properties — routes.route_state_enabled / 754073b31b4e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-98302441ddcb105234cbb0fc4349c3d0d458b05e7a396bc04050c7462526abfb"></a>

## Next pages — routes.route_state_enabled / 754073b31b4e / 4

- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7d2ea568afa7e8c29b02f60a0c505c9addd836a80ff07c3cdf07151271695bf7"></a>

## routes.simple_route — routes.simple_route / 2e53ae9d39de / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- routes.simple_route

<a id="canonical-e67119fa6334341a645e5d1a04af274a2fbb9bc46fd6328397c355a56f2f3c17"></a>

Type: `"single"`. Computed.

Simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the
matching traffic to the associated pools.

Upstream description:

A simple route matches on path, incoming header, incoming port and/or HTTP method and forwards the
matching traffic to the associated pools.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-caching": "[\"caching_disable\",\"caching_inherit\"]",
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"disable_host_rewrite\",\"host_rewrite\"]"
}
```

<a id="canonical-c5fa1057c71e3110b88c7214738856a64a06a5588ed9675a3799b0f457d48268"></a>

## Direct properties — routes.simple_route / 2e53ae9d39de / 3

- [advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f): complete subsection reference.

- [auto_host_rewrite](data-sources--http_loadbalancer--reference--group-024.md#canonical-3d492b3bdcca763d60b54e12c22e82299596d8500ba00a401784178986605aee): complete subsection reference.

- [caching_disable](data-sources--http_loadbalancer--reference--group-024.md#canonical-5922e2977a3d71bb322b00946c4b30952fb7fdd019e711ad38fe0694ef48aa48): complete subsection reference.

- [caching_inherit](data-sources--http_loadbalancer--reference--group-024.md#canonical-780d60a72d3993fb38d01eae6e6c0f445bab2ba40df7c10859410497869b5fce): complete subsection reference.

- [disable_host_rewrite](data-sources--http_loadbalancer--reference--group-024.md#canonical-bb64bd3e4bab0522d7301ed7c9bd4204e6ba4350583a5d3fde8eda7b2eaa92b4): complete subsection reference.

- [headers](data-sources--http_loadbalancer--reference--group-025.md#canonical-de255b4567a2f226a97c30866b898d06f983c5dc1b19f276a1c1a56219d65744): complete subsection reference.

<a id="canonical-55b6f6ee98fadcae7d962455ec938b6f99e162248d3570e3e44ba7281b3e9239"></a>

<a id="canonical-a6aa5c4e5f5a82ae4356bfc408d538236384fc2226db189559d3626195d4e649"></a>

## host_rewrite property — routes.simple_route / 2e53ae9d39de / 4

Type: `"string"`. Computed.

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Upstream description:

Exclusive with \[auto\_host\_rewrite disable\_host\_rewrite\] Host header will be swapped with this
value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-75f6d2b6dee6882212eff847c98c2719a65390b5076b477b007c3a05351757ad"></a>

<a id="canonical-9348764cea1ef25763331036f7b69decd708634db9ae0a8d5dc957e8c0c1d683"></a>

## http_method property — routes.simple_route / 2e53ae9d39de / 5

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Upstream description:

Specifies the HTTP method used to access a resource.

Any HTTP Method.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [incoming_port](data-sources--http_loadbalancer--reference--group-025.md#canonical-29cba187b019955171440e9188dcedbe57464be112234f015041aea135f6c35f): complete subsection reference.

- [origin_pools](data-sources--http_loadbalancer--reference--group-025.md#canonical-037a8c7968f0f9bfaae82eb0371c9ca19c0c75787702540b4d66be15b16efb5d): complete subsection reference.

- [path](data-sources--http_loadbalancer--reference--group-025.md#canonical-740eb0b78dceeed12b23d9b6594c906b796d3474248f1238748fe40172dadfc2): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-c72afd3fe56cf4c285c457636927542268eb685d6e9a616b7cb7dc6ecea0db38): complete subsection reference.

<a id="canonical-b2b6d8873d1668f52412d8e22a1c8d2893a034211476f64f535553d0011ba855"></a>

## Next pages — routes.simple_route / 2e53ae9d39de / 6

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.auto_host_rewrite](data-sources--http_loadbalancer--reference--group-024.md#canonical-3d492b3bdcca763d60b54e12c22e82299596d8500ba00a401784178986605aee)
- [routes.simple_route.caching_disable](data-sources--http_loadbalancer--reference--group-024.md#canonical-5922e2977a3d71bb322b00946c4b30952fb7fdd019e711ad38fe0694ef48aa48)
- [routes.simple_route.caching_inherit](data-sources--http_loadbalancer--reference--group-024.md#canonical-780d60a72d3993fb38d01eae6e6c0f445bab2ba40df7c10859410497869b5fce)
- [routes.simple_route.disable_host_rewrite](data-sources--http_loadbalancer--reference--group-024.md#canonical-bb64bd3e4bab0522d7301ed7c9bd4204e6ba4350583a5d3fde8eda7b2eaa92b4)
- [routes.simple_route.headers](data-sources--http_loadbalancer--reference--group-025.md#canonical-de255b4567a2f226a97c30866b898d06f983c5dc1b19f276a1c1a56219d65744)
- [routes.simple_route.incoming_port](data-sources--http_loadbalancer--reference--group-025.md#canonical-29cba187b019955171440e9188dcedbe57464be112234f015041aea135f6c35f)
- [routes.simple_route.origin_pools](data-sources--http_loadbalancer--reference--group-025.md#canonical-037a8c7968f0f9bfaae82eb0371c9ca19c0c75787702540b4d66be15b16efb5d)
- [routes.simple_route.path](data-sources--http_loadbalancer--reference--group-025.md#canonical-740eb0b78dceeed12b23d9b6594c906b796d3474248f1238748fe40172dadfc2)
- [routes.simple_route.query_params](data-sources--http_loadbalancer--reference--group-025.md#canonical-c72afd3fe56cf4c285c457636927542268eb685d6e9a616b7cb7dc6ecea0db38)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f602e6beb62063e89f12d8c2ac2eacdacb269f280475a367794b044b7277e32"></a>

## routes.simple_route.advanced_options — routes.simple_route.advanced_options / d4925c81b39a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- routes.simple_route.advanced_options

<a id="canonical-6948fd53a1698ac52136ee521d09b3478cd0b0dcccea9ecd6bb5a2b19adb802a"></a>

Type: `"single"`. Computed.

Configure advanced OPTIONS for route like path rewrite, hash policy, etc.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-bot_defense_javascript_injection_choice": "[\"bot_defense_javascript_injection\",\"inherited_bot_defense_javascript_injection\"]",
  "x-ves-oneof-field-buffer_choice": "[\"buffer_policy\",\"common_buffering\"]",
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-hash_policy_choice": "[\"common_hash_policy\",\"specific_hash_policy\"]",
  "x-ves-oneof-field-mirroring_choice": "[\"disable_mirroring\",\"mirror_policy\"]",
  "x-ves-oneof-field-retry_policy_choice": "[\"default_retry_policy\",\"no_retry_policy\",\"retry_policy\"]",
  "x-ves-oneof-field-rewrite_choice": "[\"disable_prefix_rewrite\",\"prefix_rewrite\",\"regex_rewrite\"]",
  "x-ves-oneof-field-spdy_choice": "[\"disable_spdy\",\"enable_spdy\"]",
  "x-ves-oneof-field-waf_choice": "[\"app_firewall\",\"disable_waf\",\"inherited_waf\"]",
  "x-ves-oneof-field-waf_exclusion_choice": "[\"inherited_waf_exclusion\",\"waf_exclusion_policy\"]",
  "x-ves-oneof-field-websocket_choice": "[\"disable_web_socket_config\",\"web_socket_config\"]"
}
```

<a id="canonical-a999d0eda6c4da5f328e440a123f2967f56e80829029207afe0dc3e144e62e11"></a>

## Direct properties — routes.simple_route.advanced_options / d4925c81b39a / 3

- [app_firewall](data-sources--http_loadbalancer--reference--group-023.md#canonical-ea0be713e0b12f1d2cb811acf8e18402ef3df3946ab48e95928943b3e189e0a6): complete subsection reference.

- [bot_defense_javascript_injection](data-sources--http_loadbalancer--reference--group-023.md#canonical-a070af4c4222824f640e8d41af3748c66a66c4c00ec31e66ad2fc77af4a68bcf): complete subsection reference.

- [buffer_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-2743af63f645df04a829535c2f29e231ca5840d639c980e3c546faf26ab17d95): complete subsection reference.

- [common_buffering](data-sources--http_loadbalancer--reference--group-023.md#canonical-a709e2446ff2c196cdbb532465c599ed281b51e0bf41a93e317acde28217e5ce): complete subsection reference.

- [common_hash_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-e99841d4ddb24eb447dfc07e539b8f5873bddb21e70fb1d663bd117d205346ac): complete subsection reference.

- [cors_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-6a8957e8d46322f43983c8385e733bc7cd2b3bbea376b8e597fa7e19e4c52b30): complete subsection reference.

- [csrf_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-60dfdbdf9bf9617f1060c279c98ad547e8cc931bfac5c03c99af2f86e586218b): complete subsection reference.

- [default_retry_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-7ccc0016d13300c87cf9ac20b8ad1096a47ed6e6ab41f976be9b60772e6e9e8d): complete subsection reference.

<a id="canonical-2eb0be98f0c45b61b34a67b87594feea4e0e34509a1982d548a6388e72f874bd"></a>

<a id="canonical-b92fb65a5ca88d831eed33dab3e2d9e19fdbd3be804ed3de46a3d755b71a6b27"></a>

## disable_location_add property — routes.simple_route.advanced_options / d4925c81b39a / 4

Type: `"bool"`. Computed.

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

Upstream description:

Disables append of x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; at route level, if it is
configured at virtual-host level. This configuration is ignored on CE sites.

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

- [disable_mirroring](data-sources--http_loadbalancer--reference--group-023.md#canonical-1e184ac2707318ef658f31d8a200477233ad6e840aa5a2932f952538c1bca0aa): complete subsection reference.

- [disable_prefix_rewrite](data-sources--http_loadbalancer--reference--group-023.md#canonical-33f80f837020152d3321818fb412613893962d256b640072c5b0e8b941ddf76b): complete subsection reference.

- [disable_spdy](data-sources--http_loadbalancer--reference--group-023.md#canonical-ab26f3a361f1791f0aa8e26c6cdf23037dcff177be9dae6225d312a921239c3b): complete subsection reference.

- [disable_waf](data-sources--http_loadbalancer--reference--group-023.md#canonical-64c28d52db0e7e954219acb1525f4209d51832102ab43625827897696852e095): complete subsection reference.

- [disable_web_socket_config](data-sources--http_loadbalancer--reference--group-023.md#canonical-1c3234e93b7a6e2eaeb38762b76dc7d09712c63d6d88f562bf7ed515aeba30ac): complete subsection reference.

- [do_not_retract_cluster](data-sources--http_loadbalancer--reference--group-023.md#canonical-a3a2380a30e85754a58a4815bfb5a8ee66ba994bacecd1dd6d6de3e1e9177dc4): complete subsection reference.

- [enable_spdy](data-sources--http_loadbalancer--reference--group-023.md#canonical-0b87e7a126800e2b39cb6e8f205aaca84e5cbdedab1fa6ec45862a01bb5efe99): complete subsection reference.

- [endpoint_subsets](data-sources--http_loadbalancer--reference--group-023.md#canonical-2df64617e248296615238f6a69b0c1b3ab5e8422f1c4d9878c534d8b84a9a80e): complete subsection reference.

- [inherited_bot_defense_javascript_injection](data-sources--http_loadbalancer--reference--group-023.md#canonical-c77539f53ba7dcaacbfeeddf31daa1e7ebf3fe61e95f992380c7ee1b332375de): complete subsection reference.

- [inherited_waf](data-sources--http_loadbalancer--reference--group-023.md#canonical-151e4e141e0b8478564e09810fd4293346787906af422e8900d0efc3844eb246): complete subsection reference.

- [inherited_waf_exclusion](data-sources--http_loadbalancer--reference--group-023.md#canonical-19688b6f66df4309a4e12e0c6d13b5df0c5761bb5a169371bacd566e9a161815): complete subsection reference.

- [mirror_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f769c35b2f994587ec953a41759d45ee49a616e417b0becf45983a71a456ef5): complete subsection reference.

- [no_retry_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-d694da319839efa801fda458340d5a2b9ddb71372337297699448b6d15476113): complete subsection reference.

<a id="canonical-cbd986f004cce54209e0e02bffa18d7703a862ed0bc1aa1b29eecf7d15dd4c66"></a>

<a id="canonical-f83d700aa453fa6ec7f6db5f2784f1cf3440b30a341c976f5c1f16dc78664e94"></a>

## prefix_rewrite property — routes.simple_route.advanced_options / d4925c81b39a / 5

Type: `"string"`. Computed.

Exclusive with \[disable\_prefix\_rewrite regex\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regex path
matching, the entire path (not including the query string) will be swapped with this value.

Upstream description:

Exclusive with \[disable\_prefix\_rewrite regex\_rewrite\] prefix\_rewrite indicates that during
forwarding, the matched prefix (or path) should be swapped with its value. When using regex path
matching, the entire path (not including the query string) will be swapped with this value.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-308f1d21cda97bdffb998ae14b3376be9b463f3419549b275089ba0da451c495"></a>

<a id="canonical-6c5730cd0a08d0950b64b0c2e7df828370c7c368c248e33ab940c5d1b480e8d3"></a>

## priority property — routes.simple_route.advanced_options / d4925c81b39a / 6

Type: `"string"`. Computed.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Upstream description:

Priority routing for each request. Different connection pools are used based on the priority
selected for the request. Also, circuit-breaker configuration at destination cluster is chosen based
on selected priority.

Default routing mechanism High-Priority routing mechanism.

Receipt-pinned upstream constraints:

```json
{
  "default": "DEFAULT",
  "enum": [
    "DEFAULT",
    "HIGH"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [regex_rewrite](data-sources--http_loadbalancer--reference--group-023.md#canonical-c0d0fc65461843d23c8820bb8f9422b6c455947b0710550459d44cb396fa94b5): complete subsection reference.

- [request_cookies_to_add](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f3a8a13d10671a490d330b6c70591cbd941c941b58f82b19df0fe4c42e76852): complete subsection reference.

<a id="canonical-10261d1375a4bbc41e4740da7cb3f003a4d682dc170697118651f9dc30c851a7"></a>

<a id="canonical-248ec23849eeb58fbad9be410ea46ab46977bef3f5a3e564621071d353c6b089"></a>

## request_cookies_to_remove property — routes.simple_route.advanced_options / d4925c81b39a / 7

Type: `["list", "string"]`. Computed.

List of keys of Cookies to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [request_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-9be9557dde347d2723c87163cf37b8d7322baf0f5e2adb7a120366f63e6a90a6): complete subsection reference.

<a id="canonical-e2e35aa484030d72265babb617e005f553bc3eabc5f008ee50747813623a1433"></a>

<a id="canonical-9103c15b935130366ad88ff75f3ae384af90f0cb5010dd805104d4b052e07268"></a>

## request_headers_to_remove property — routes.simple_route.advanced_options / d4925c81b39a / 8

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23): complete subsection reference.

<a id="canonical-19e86c086a5333468bc6eb53fa9e4b74688cb7061fc4348d9cfd2b88a5dfc83b"></a>

<a id="canonical-68a3a9dea68e593d04baa9467414e5383514f2cb6b3cf7dc2b0408e1570b1f03"></a>

## response_cookies_to_remove property — routes.simple_route.advanced_options / d4925c81b39a / 9

Type: `["list", "string"]`. Computed.

List of name of Cookies to be removed from the HTTP response being sent towards downstream. Entire
set-cookie header will be removed.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-e30971933d5007464e0b821668a3ae01b4a90a73b358f07cb04984decfa8ca95): complete subsection reference.

<a id="canonical-19054fb08b68243f83f9bcbcfc04dd5fb0763c87cfa3c7f927f2c6caacb1da01"></a>

<a id="canonical-63782ea9eb958ba8ec7aa86981a249f8ebac5621d55f6db9f5f2baee76f732b8"></a>

## response_headers_to_remove property — routes.simple_route.advanced_options / d4925c81b39a / 10

Type: `["list", "string"]`. Computed.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [retract_cluster](data-sources--http_loadbalancer--reference--group-024.md#canonical-629d3b101c247273fffdb2a8e94adacd50a906396d477f73469b9afa7539e6b4): complete subsection reference.

- [retry_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-c4fc0ead7e6779144e37321ac5084409dc1876214277674b2e7da8b00f43db70): complete subsection reference.

- [specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613): complete subsection reference.

<a id="canonical-b6f5efa8273e3f48c533fd60e8eea16e973f864b1f13c92c8ee4459b7e5f830a"></a>

<a id="canonical-c44ae8953e94846c1f6077bfbf2c23fde81fb087ddde951fda184bb2cd7a49d6"></a>

## timeout property — routes.simple_route.advanced_options / d4925c81b39a / 11

Type: `"number"`. Computed.

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Upstream description:

The timeout for the route including all retries, in milliseconds. Should be set to a high value or 0
(infinite timeout) for server-side streaming.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [waf_exclusion_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-2e206a249d6baf756f0d86a64ab9201567f3fc42453191a3795a30dee8e4fa96): complete subsection reference.

- [web_socket_config](data-sources--http_loadbalancer--reference--group-024.md#canonical-da8c7bfa640c00d39cdaa9f4cbbe5e56a6989c13a45e893fffe113fa834c178c): complete subsection reference.

<a id="canonical-f8f9ea5a02844f3ee49bc93ca6d9eaad27693d0f0cfe2181cb86b447f7d1e9dd"></a>

## Next pages — routes.simple_route.advanced_options / d4925c81b39a / 12

- [routes.simple_route.advanced_options.app_firewall](data-sources--http_loadbalancer--reference--group-023.md#canonical-ea0be713e0b12f1d2cb811acf8e18402ef3df3946ab48e95928943b3e189e0a6)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](data-sources--http_loadbalancer--reference--group-023.md#canonical-a070af4c4222824f640e8d41af3748c66a66c4c00ec31e66ad2fc77af4a68bcf)
- [routes.simple_route.advanced_options.buffer_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-2743af63f645df04a829535c2f29e231ca5840d639c980e3c546faf26ab17d95)
- [routes.simple_route.advanced_options.common_buffering](data-sources--http_loadbalancer--reference--group-023.md#canonical-a709e2446ff2c196cdbb532465c599ed281b51e0bf41a93e317acde28217e5ce)
- [routes.simple_route.advanced_options.common_hash_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-e99841d4ddb24eb447dfc07e539b8f5873bddb21e70fb1d663bd117d205346ac)
- [routes.simple_route.advanced_options.cors_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-6a8957e8d46322f43983c8385e733bc7cd2b3bbea376b8e597fa7e19e4c52b30)
- [routes.simple_route.advanced_options.csrf_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-60dfdbdf9bf9617f1060c279c98ad547e8cc931bfac5c03c99af2f86e586218b)
- [routes.simple_route.advanced_options.default_retry_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-7ccc0016d13300c87cf9ac20b8ad1096a47ed6e6ab41f976be9b60772e6e9e8d)
- [routes.simple_route.advanced_options.disable_mirroring](data-sources--http_loadbalancer--reference--group-023.md#canonical-1e184ac2707318ef658f31d8a200477233ad6e840aa5a2932f952538c1bca0aa)
- [routes.simple_route.advanced_options.disable_prefix_rewrite](data-sources--http_loadbalancer--reference--group-023.md#canonical-33f80f837020152d3321818fb412613893962d256b640072c5b0e8b941ddf76b)
- [routes.simple_route.advanced_options.disable_spdy](data-sources--http_loadbalancer--reference--group-023.md#canonical-ab26f3a361f1791f0aa8e26c6cdf23037dcff177be9dae6225d312a921239c3b)
- [routes.simple_route.advanced_options.disable_waf](data-sources--http_loadbalancer--reference--group-023.md#canonical-64c28d52db0e7e954219acb1525f4209d51832102ab43625827897696852e095)
- [routes.simple_route.advanced_options.disable_web_socket_config](data-sources--http_loadbalancer--reference--group-023.md#canonical-1c3234e93b7a6e2eaeb38762b76dc7d09712c63d6d88f562bf7ed515aeba30ac)
- [routes.simple_route.advanced_options.do_not_retract_cluster](data-sources--http_loadbalancer--reference--group-023.md#canonical-a3a2380a30e85754a58a4815bfb5a8ee66ba994bacecd1dd6d6de3e1e9177dc4)
- [routes.simple_route.advanced_options.enable_spdy](data-sources--http_loadbalancer--reference--group-023.md#canonical-0b87e7a126800e2b39cb6e8f205aaca84e5cbdedab1fa6ec45862a01bb5efe99)
- [routes.simple_route.advanced_options.endpoint_subsets](data-sources--http_loadbalancer--reference--group-023.md#canonical-2df64617e248296615238f6a69b0c1b3ab5e8422f1c4d9878c534d8b84a9a80e)
- [routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection](data-sources--http_loadbalancer--reference--group-023.md#canonical-c77539f53ba7dcaacbfeeddf31daa1e7ebf3fe61e95f992380c7ee1b332375de)
- [routes.simple_route.advanced_options.inherited_waf](data-sources--http_loadbalancer--reference--group-023.md#canonical-151e4e141e0b8478564e09810fd4293346787906af422e8900d0efc3844eb246)
- [routes.simple_route.advanced_options.inherited_waf_exclusion](data-sources--http_loadbalancer--reference--group-023.md#canonical-19688b6f66df4309a4e12e0c6d13b5df0c5761bb5a169371bacd566e9a161815)
- [routes.simple_route.advanced_options.mirror_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f769c35b2f994587ec953a41759d45ee49a616e417b0becf45983a71a456ef5)
- [routes.simple_route.advanced_options.no_retry_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-d694da319839efa801fda458340d5a2b9ddb71372337297699448b6d15476113)
- [routes.simple_route.advanced_options.regex_rewrite](data-sources--http_loadbalancer--reference--group-023.md#canonical-c0d0fc65461843d23c8820bb8f9422b6c455947b0710550459d44cb396fa94b5)
- [routes.simple_route.advanced_options.request_cookies_to_add](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f3a8a13d10671a490d330b6c70591cbd941c941b58f82b19df0fe4c42e76852)
- [routes.simple_route.advanced_options.request_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-9be9557dde347d2723c87163cf37b8d7322baf0f5e2adb7a120366f63e6a90a6)
- [routes.simple_route.advanced_options.response_cookies_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-fbf084299009928d815884c6c9511e04ee490ed3b3da9a2305f553c497c04c23)
- [routes.simple_route.advanced_options.response_headers_to_add](data-sources--http_loadbalancer--reference--group-024.md#canonical-e30971933d5007464e0b821668a3ae01b4a90a73b358f07cb04984decfa8ca95)
- [routes.simple_route.advanced_options.retract_cluster](data-sources--http_loadbalancer--reference--group-024.md#canonical-629d3b101c247273fffdb2a8e94adacd50a906396d477f73469b9afa7539e6b4)
- [routes.simple_route.advanced_options.retry_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-c4fc0ead7e6779144e37321ac5084409dc1876214277674b2e7da8b00f43db70)
- [routes.simple_route.advanced_options.specific_hash_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-129d70719eb85bb4b1c61f470f5ea78bacd54b1da60a7b2f1dcdfdd7d2da0613)
- [routes.simple_route.advanced_options.waf_exclusion_policy](data-sources--http_loadbalancer--reference--group-024.md#canonical-2e206a249d6baf756f0d86a64ab9201567f3fc42453191a3795a30dee8e4fa96)
- [routes.simple_route.advanced_options.web_socket_config](data-sources--http_loadbalancer--reference--group-024.md#canonical-da8c7bfa640c00d39cdaa9f4cbbe5e56a6989c13a45e893fffe113fa834c178c)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ea0be713e0b12f1d2cb811acf8e18402ef3df3946ab48e95928943b3e189e0a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a5b38c304393c2138f8b36c6528e6f9e8f3fc12e74d3bf93cb878864aad4e106"></a>

## routes.simple_route.advanced_options.app_firewall — routes.simple_route.advanced_options.app_firewall / 5d6020aa256c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.app_firewall

<a id="canonical-ae2ddb927415431e38e4abd4afdf6356761632167abb3a4ebe21bda51b6068bd"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-108152e5fa08060dfb658f2c0bb16802129bc98a721a63ec111b467f4e1ce77f"></a>

## Direct properties — routes.simple_route.advanced_options.app_firewall / 5d6020aa256c / 3

<a id="canonical-20f5b765a7e47b34fd1cb58add4249fa8868a40db41e1c5318614f18fefdb67c"></a>

<a id="canonical-b40e2d16472b68ce4d9ad1a53c5fbd2ffd472c74b3aa8c99cb912e538d15b6fa"></a>

## name property — routes.simple_route.advanced_options.app_firewall / 5d6020aa256c / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-ecfc1d52133a6b478d74c923c0498c121ff56a59bdfa3a25bba7a65a75654668"></a>

<a id="canonical-821c9cfef7bdbbd53a0cc5f098f8e8bad7be1c6eabad4c3cf38ff67c24f8144c"></a>

## namespace property — routes.simple_route.advanced_options.app_firewall / 5d6020aa256c / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-39b9bb37298e7821b59c18fc9424fd7dfc9cc9e1a5c81410522e13606091fedb"></a>

<a id="canonical-1e9d233cb6ffaac3e25d61e540e1a3cfc406ba5c2746d18c221b52ad5926154f"></a>

## tenant property — routes.simple_route.advanced_options.app_firewall / 5d6020aa256c / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-ea4784ece4fce350b32076f76c4d2ee117d38b8d6e6f0a09cb0df282a2224485"></a>

## Next pages — routes.simple_route.advanced_options.app_firewall / 5d6020aa256c / 7

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a070af4c4222824f640e8d41af3748c66a66c4c00ec31e66ad2fc77af4a68bcf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7395ff9ba64325ee404df53f9f5e538c60e649aa09044c506151517725c22167"></a>

## routes.simple_route.advanced_options.bot_defense_javascript_injection — routes.simple_route.advanced_options.bot_defense_javascript_injection / 03185fd97d93 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.bot_defense_javascript_injection

<a id="canonical-7d239c592b10a8c3fffbf0ca5251c02586b97c186f777f441b41f95122886d7b"></a>

Type: `"single"`. Computed.

Bot Defense Javascript Injection Configuration for inline bot defense deployments.

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

<a id="canonical-2c888f47f1bdd8807de305740456f8af839d6126575464d75ff1d8a46b562391"></a>

## Direct properties — routes.simple_route.advanced_options.bot_defense_javascript_injection / 03185fd97d93 / 3

<a id="canonical-419f10a037345bf366b626f717422c83429897351c2dbb2497ee0abfe83a367a"></a>

<a id="canonical-3c156ba36cc4cd3c7adb3e85df0fce62369d04ff1ee0e427a7611253f9dd38b6"></a>

## javascript_location property — routes.simple_route.advanced_options.bot_defense_javascript_injection / 03185fd97d93 / 4

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

- [javascript_tags](data-sources--http_loadbalancer--reference--group-023.md#canonical-9096e1c1b55637f7c333e9a4fce8c38ed9069b5ae4523dcef18e74d74cd379d0): complete subsection reference.

<a id="canonical-c89a631d50b5ed9f88cfc550814682c46d839919ad3feca46bcc0864095e3c0b"></a>

## Next pages — routes.simple_route.advanced_options.bot_defense_javascript_injection / 03185fd97d93 / 5

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](data-sources--http_loadbalancer--reference--group-023.md#canonical-9096e1c1b55637f7c333e9a4fce8c38ed9069b5ae4523dcef18e74d74cd379d0)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-9096e1c1b55637f7c333e9a4fce8c38ed9069b5ae4523dcef18e74d74cd379d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-34cf358fdf9d351611dffe10e1adbd92422b6898448bd91fc76c0401349c7da4"></a>

## routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / 39bde2f1d6f3 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](data-sources--http_loadbalancer--reference--group-023.md#canonical-a070af4c4222824f640e8d41af3748c66a66c4c00ec31e66ad2fc77af4a68bcf)
- routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags

<a id="canonical-e60bbb74866debc7ec1c200c6d9f0f447d3c1dde6b5aac830cb5d2a24cc8e3ae"></a>

Type: `"list"`. Computed.

Select Add item to configure your javascript tag. If adding both Bot Adv and Fraud, the Bot
Javascript should be added first.

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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-971737345384015f243dab0b42ec795b6b01e6bb93e89e808825f5b0c0ee54c8"></a>

## Direct properties — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / 39bde2f1d6f3 / 3

<a id="canonical-b96c6cdea28c80e0b803063c9386012ca710ca659995de0408b9148bd8cce276"></a>

<a id="canonical-80bd1968be095456b737fa176c68e5f7a4948463af2b6a7f8ef4ce81d928e7d7"></a>

## javascript_url property — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / 39bde2f1d6f3 / 4

Type: `"string"`. Computed.

Please enter the full URL (include domain and path), or relative path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 2048,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 2048,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 2048,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "2048",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [tag_attributes](data-sources--http_loadbalancer--reference--group-023.md#canonical-fc3ad4a1e002e15e653c71f5e350f872a6daa98adbe901883dab118c44b2bfbf): complete subsection reference.

<a id="canonical-14cab86be7a368ec04f9a6950d4306af1a691d66f3a593f1b28f69ba526abac4"></a>

## Next pages — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / 39bde2f1d6f3 / 5

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes](data-sources--http_loadbalancer--reference--group-023.md#canonical-fc3ad4a1e002e15e653c71f5e350f872a6daa98adbe901883dab118c44b2bfbf)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](data-sources--http_loadbalancer--reference--group-023.md#canonical-a070af4c4222824f640e8d41af3748c66a66c4c00ec31e66ad2fc77af4a68bcf)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-fc3ad4a1e002e15e653c71f5e350f872a6daa98adbe901883dab118c44b2bfbf"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-977ef37144e0b3b8c0d319efded3bfa00c082856f1df14ef70ae666ffbdce645"></a>

## routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / fd90dfba359d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection](data-sources--http_loadbalancer--reference--group-023.md#canonical-a070af4c4222824f640e8d41af3748c66a66c4c00ec31e66ad2fc77af4a68bcf)
- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](data-sources--http_loadbalancer--reference--group-023.md#canonical-9096e1c1b55637f7c333e9a4fce8c38ed9069b5ae4523dcef18e74d74cd379d0)
- routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags.tag_attributes

<a id="canonical-87be7bfaf1946d0fdaf7a783ce6d1dadf53fd2960f1d68ea034af8ab28f67540"></a>

Type: `"list"`. Computed.

Add the tag attributes you want to include in your Javascript tag.

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

<a id="canonical-4e89af6894c71f32dbd12f75bba3d7b0ceba12a113823ef12465f29d9c680f49"></a>

## Direct properties — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / fd90dfba359d / 3

<a id="canonical-f17641f8e4f5c6455a0e70d8fb0987e0161913b3f70925675796ad331a37f828"></a>

<a id="canonical-7af243a35c551794f33f60791277ca12e0535a24926e3ce21befcc7145b0b28d"></a>

## javascript_tag property — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / fd90dfba359d / 4

Type: `"string"`. Computed.

\[Enum:
JS\_ATTR\_ID|JS\_ATTR\_CID|JS\_ATTR\_CN|JS\_ATTR\_API\_DOMAIN|JS\_ATTR\_API\_URL|JS\_ATTR\_API\_PATH|JS\_ATTR\_ASYNC|JS\_ATTR\_DEFER\]
Select from one of the predefined tag attributes. Possible values are \`JS\_ATTR\_ID\`,
\`JS\_ATTR\_CID\`, \`JS\_ATTR\_CN\`, \`JS\_ATTR\_API\_DOMAIN\`, \`JS\_ATTR\_API\_URL\`,
\`JS\_ATTR\_API\_PATH\`, \`JS\_ATTR\_ASYNC\`, \`JS\_ATTR\_DEFER\`. Defaults to \`JS\_ATTR\_ID\`.

Upstream description:

Select from one of the predefined tag attributes.

Receipt-pinned upstream constraints:

```json
{
  "default": "JS_ATTR_ID",
  "enum": [
    "JS_ATTR_ID",
    "JS_ATTR_CID",
    "JS_ATTR_CN",
    "JS_ATTR_API_DOMAIN",
    "JS_ATTR_API_URL",
    "JS_ATTR_API_PATH",
    "JS_ATTR_ASYNC",
    "JS_ATTR_DEFER"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-ebb42ec65df2b97a18fdedb9b7f01279aeeae4adb28761866b88af3dc5b2e28d"></a>

<a id="canonical-ddb97abe8c5cb32388ab328844d13d7b5568c9bc8617274c735ba9ae037a1dd2"></a>

## tag_value property — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / fd90dfba359d / 5

Type: `"string"`. Computed.

Value. Add the tag attribute value.

Upstream description:

Add the tag attribute value.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.max_bytes": "1024"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1024"
  }
}
```

<a id="canonical-5e0dec5141be858047b7ed072f668c876d2ae0555f1ece4c0d1497bb9213ea5e"></a>

## Next pages — routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript / fd90dfba359d / 6

- [routes.simple_route.advanced_options.bot_defense_javascript_injection.javascript_tags](data-sources--http_loadbalancer--reference--group-023.md#canonical-9096e1c1b55637f7c333e9a4fce8c38ed9069b5ae4523dcef18e74d74cd379d0)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2743af63f645df04a829535c2f29e231ca5840d639c980e3c546faf26ab17d95"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d0a147cb0e7d9eb25ec2778054e3ae689fa7d6ec0fb36a2404dd7be02d62b740"></a>

## routes.simple_route.advanced_options.buffer_policy — routes.simple_route.advanced_options.buffer_policy / 64b0db92fcdd / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.buffer_policy

<a id="canonical-1808c7f39f6572770c24bbbffb70c2ed2856dc6d2ad0db479d9cd1ff9e174497"></a>

Type: `"single"`. Computed.

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Upstream description:

Some upstream applications are not capable of handling streamed data. This config enables buffering
the entire request before sending to upstream application. We can specify the maximum buffer size
and buffer interval with this config.

Buffering can be enabled and disabled at VirtualHost and Route levels Route level buffer
configuration takes precedence.

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

<a id="canonical-dea4f2e1ce7ae788fdd49c1f4bacaf24747385ae9a01a489e7d927d9f5744843"></a>

## Direct properties — routes.simple_route.advanced_options.buffer_policy / 64b0db92fcdd / 3

<a id="canonical-853c8abec5b98c71070f0d3919a09feb1c021e79137234aeb49084b528d34532"></a>

<a id="canonical-5d93d8d19690961d3d4c4f664331f477867eaadbacf831c1adf23ebf53c1e1a9"></a>

## disabled property — routes.simple_route.advanced_options.buffer_policy / 64b0db92fcdd / 4

Type: `"bool"`. Computed.

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable buffering for a particular route. This is useful when virtual-host has buffering, but we
need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-864e385ba267fba4d539812495a28d9e518d60f76a73e1837f0ca3d6aec68b8b"></a>

<a id="canonical-8039256888126906fe34b323ad8c3e3333f121e1b2784e519db0785177428830"></a>

## max_request_bytes property — routes.simple_route.advanced_options.buffer_policy / 64b0db92fcdd / 5

Type: `"number"`. Computed.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Upstream description:

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 10485760,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "10485760"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "10485760"
  }
}
```

<a id="canonical-759c08dc84d2cbac3a3bf9ad3918aaa1ad32889d36a0cd82166a559792eff329"></a>

## Next pages — routes.simple_route.advanced_options.buffer_policy / 64b0db92fcdd / 6

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a709e2446ff2c196cdbb532465c599ed281b51e0bf41a93e317acde28217e5ce"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ae16322bbf326b2138b00ca38b685dede52eff6581ed27e41d4544cd54744b8f"></a>

## routes.simple_route.advanced_options.common_buffering — routes.simple_route.advanced_options.common_buffering / 2335383c6dd8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.common_buffering

<a id="canonical-e353edbfe16711d01887e52d704cf74a13e81a6895ce1fe6d11e0c7332ddb4ee"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for common buffering.

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

<a id="canonical-ea6df4e43eef31b2ba3e770fd0bbb0c455b11869c22b3fb9ba796c20739f0c5b"></a>

## Direct properties — routes.simple_route.advanced_options.common_buffering / 2335383c6dd8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d49158c601a2d8be74400d2ce9052fa83d45c44d1a1055b16f63fe8908dd5b79"></a>

## Next pages — routes.simple_route.advanced_options.common_buffering / 2335383c6dd8 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-e99841d4ddb24eb447dfc07e539b8f5873bddb21e70fb1d663bd117d205346ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-306179c637da07abf07bc913398fe9113675abe69ff8d5f42d84af55a5a37c25"></a>

## routes.simple_route.advanced_options.common_hash_policy — routes.simple_route.advanced_options.common_hash_policy / b32786f22594 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.common_hash_policy

<a id="canonical-b9ff84856533e90245c319eb8d0dab033d331daf3ee726b5fea4b341e968aec0"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-1039fc1e933f05925e48da8b7e4d12f487386b005183d62ab753bd313dfdb28c"></a>

## Direct properties — routes.simple_route.advanced_options.common_hash_policy / b32786f22594 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8793c7f544b2a853b0c735569e26c5e630cccaecaaf19ef1246d65ea67617c1c"></a>

## Next pages — routes.simple_route.advanced_options.common_hash_policy / b32786f22594 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6a8957e8d46322f43983c8385e733bc7cd2b3bbea376b8e597fa7e19e4c52b30"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b97b7f973641f9bc5657dc1b5e5a2158a875ba474bce67a5b2b8ceef5384c2b4"></a>

## routes.simple_route.advanced_options.cors_policy — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.cors_policy

<a id="canonical-04c4af184cc4af28fccc6f96dea0499ee04ff60e8305575e52f8cc199aa91a6d"></a>

Type: `"single"`. Computed.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence. An example of an Cross origin HTTP request GET
/resources/public-data/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS
X 10.5..

Upstream description:

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.html Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
MAC OS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8 Accept-Language: en-us,en;q=0.5
Accept-Encoding: gzip,deflate Accept-Charset: ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive
Origin: http&#58;//foo.example Access-Control-Request-Method: POST Access-Control-Request-Headers:
X-PINGOTHER, Content-Type

HTTP/1.1 204 No Content Date: Mon, 01 Dec 2008 01:15:39 GMT Server: Apache/2.0.61 (Unix)
Access-Control-Allow-Origin: http&#58;//foo.example Access-Control-Allow-Methods: POST, GET, OPTIONS
Access-Control-Allow-Headers: X-PINGOTHER, Content-Type Access-Control-Max-Age: 86400 Vary:
Accept-Encoding, Origin Keep-Alive: timeout=2, max=100 Connection: Keep-Alive.

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

<a id="canonical-a469cd63c88664165c1332a0f97aed8fe93be68f6d6897339d9143727403aea3"></a>

## Direct properties — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 3

<a id="canonical-5021a756029bb7dbfef4efafe41e7585443e4a581a630d08a4387e5fabaa2a27"></a>

<a id="canonical-ea40d3f3f2b88da461ff687279369b5a94669c4232d7d9e9de7a2a683a4879e7"></a>

## allow_credentials property — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 4

Type: `"bool"`. Computed.

Specifies whether the resource allows credentials.

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

<a id="canonical-eb036a65113ba4bc7b9139b5b33b72aa2c51d4d3b07af4c274468f75c7268cff"></a>

<a id="canonical-4a8acfca29d0322f0f0cb6f954891e4863a6ec6676b7d62a70e2c31fcd4fdafa"></a>

## allow_headers property — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 5

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-headers header.

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
  }
}
```

<a id="canonical-fac813fd31eec0ad0ca22b48448fb9ee0c50ffdb767b169ad07b4e4df1b81497"></a>

<a id="canonical-fa13d09c04741bd3cd7b805c493de6a3cb47a8643555c2f5d775bbecafddefd5"></a>

## allow_methods property — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 6

Type: `"string"`. Computed.

Specifies the content for the access-control-allow-methods header.

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
    "ves.io.schema.rules.string.http_valid_methods": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_valid_methods": "true"
  }
}
```

<a id="canonical-6ca0928249d7121eb15d6bba9f485aaf2ff9f4fefcb6efa75f4c0c97471f5285"></a>

<a id="canonical-68ee074928b7f45db753038ff8a294b6983e080ec70e805fba790aad2df79229"></a>

## allow_origin property — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 7

Type: `["list", "string"]`. Computed.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Upstream description:

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b6ed754277686ec4fd324e041827cc27a9c05e8936dee3c057b4d74a454d7dcb"></a>

<a id="canonical-f1cbce3fbcdc7ff942d322efd5b2f31d77b032b3716cb9b2a924f4c17829f235"></a>

## allow_origin_regex property — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 8

Type: `["list", "string"]`. Computed.

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Upstream description:

Specifies regex patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

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
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-b20aa5502a9f3f500b38197cd50a4e1283f3476813c9fae71ec2a7a85b2426d1"></a>

<a id="canonical-97505ff7e9fdf5819b90014e1a328a90f030bf4afdba8e5102751c146e17cc71"></a>

## disabled property — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 9

Type: `"bool"`. Computed.

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

Upstream description:

Disable the CorsPolicy for a particular route. This is useful when virtual-host has CorsPolicy, but
we need to disable it on a specific route. The value of this field is ignored for virtual-host.

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

<a id="canonical-2640a87e8b8b2fc6fa0faa638ed43b137e9debb463d8efba32a7ccb323f7b861"></a>

<a id="canonical-fad936a150ee0fbd7de16e028e156e3124cfab679d6c2560a26e8de2b6d811d4"></a>

## expose_headers property — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 10

Type: `"string"`. Computed.

Specifies the content for the access-control-expose-headers header.

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
  }
}
```

<a id="canonical-68fd0f35335223151ce9efe66e14b53abb2ac9a66c816d02dcf5689a1e5c9dc8"></a>

<a id="canonical-d36e75529eac96390a50496ac3df39db61ff3ee94ad9e31cd8bd47753b605eb0"></a>

## maximum_age property — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 11

Type: `"number"`. Computed.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Upstream description:

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours)

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
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": -1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "-1",
    "ves.io.schema.rules.int32.lte": "86400"
  }
}
```

<a id="canonical-7492866f580f1d06c5f286817f3f23ed02c380c8cbdf6bd020b84e465e819c34"></a>

## Next pages — routes.simple_route.advanced_options.cors_policy / 8c2415c4e960 / 12

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-60dfdbdf9bf9617f1060c279c98ad547e8cc931bfac5c03c99af2f86e586218b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3aac80d2cfeabf05ca64ea4c22254e2daedf8c794653fd7277b6772811ed7146"></a>

## routes.simple_route.advanced_options.csrf_policy — routes.simple_route.advanced_options.csrf_policy / 4ba97de1822b / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.csrf_policy

<a id="canonical-d84b201108c651ba6cba9405fe30a45486a3ef7a422c0151782c97accf68c440"></a>

Type: `"single"`. Computed.

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host. 1. The origin that caused the user
agent..

Upstream description:

To mitigate CSRF attack , the policy checks where a request is coming from to determine if the
request's origin is the same as its destination.the policy relies on two pieces of information used
in determining if a request originated from the same host.

&#8203;1. The origin that caused the user agent to issue the request (source origin). &#8203;2. The
origin that the request is going to (target origin). When the policy evaluating a request, it
ensures both pieces of information are present and compare their values. If the source origin is
missing or origins do not match the request is rejected. The exception to this being if the
source-origin has been added to they policy as valid. Because CSRF attacks specifically target
state-changing requests, the policy only acts on the HTTP requests that have state-changing method
(PUT,POST, etc.).

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

<a id="canonical-52bf2ad56105de9ae253f930084aad1298925cd520c56968ece080e5d30db214"></a>

## Direct properties — routes.simple_route.advanced_options.csrf_policy / 4ba97de1822b / 3

- [all_load_balancer_domains](data-sources--http_loadbalancer--reference--group-023.md#canonical-6d92497562f971eabe97e56109785c25c3f919b0fc736f5a53eb323d0ac3c8d0): complete subsection reference.

- [custom_domain_list](data-sources--http_loadbalancer--reference--group-023.md#canonical-b313413f116e071d944da8f4fc88406629d95e560fdd61d8a8a64c281e71cc5e): complete subsection reference.

- [disabled](data-sources--http_loadbalancer--reference--group-023.md#canonical-62edfac03b5989488e5de4c1d5c48704938781d69afaea31c91138e285909bb9): complete subsection reference.

<a id="canonical-4745f22971bfe63bdeb6adbe53a5ef7f3a1776168f5250731898f8f971f4e82d"></a>

## Next pages — routes.simple_route.advanced_options.csrf_policy / 4ba97de1822b / 4

- [routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains](data-sources--http_loadbalancer--reference--group-023.md#canonical-6d92497562f971eabe97e56109785c25c3f919b0fc736f5a53eb323d0ac3c8d0)
- [routes.simple_route.advanced_options.csrf_policy.custom_domain_list](data-sources--http_loadbalancer--reference--group-023.md#canonical-b313413f116e071d944da8f4fc88406629d95e560fdd61d8a8a64c281e71cc5e)
- [routes.simple_route.advanced_options.csrf_policy.disabled](data-sources--http_loadbalancer--reference--group-023.md#canonical-62edfac03b5989488e5de4c1d5c48704938781d69afaea31c91138e285909bb9)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-6d92497562f971eabe97e56109785c25c3f919b0fc736f5a53eb323d0ac3c8d0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-deacccba6be58b04c1ab7963b9aaaa58a39ec6537a850408929710757b097ba4"></a>

## routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains — routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains / e14cf1024f40 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.csrf_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-60dfdbdf9bf9617f1060c279c98ad547e8cc931bfac5c03c99af2f86e586218b)
- routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains

<a id="canonical-c5e39beb5105a40b04467709c8bb718d48f8a0111378687bd72007dd91b8d983"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for all load balancer domains.

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

<a id="canonical-e1b87440ba7c0e60f95afccc29aaf2f920a05f5b0b7ce3af7194d526ecb1a1a7"></a>

## Direct properties — routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains / e14cf1024f40 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2567bb1c7227d0e58df284f68f964673ea0b629e0f1c79b18c49e9dc1438d80a"></a>

## Next pages — routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains / e14cf1024f40 / 4

- [routes.simple_route.advanced_options.csrf_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-60dfdbdf9bf9617f1060c279c98ad547e8cc931bfac5c03c99af2f86e586218b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-b313413f116e071d944da8f4fc88406629d95e560fdd61d8a8a64c281e71cc5e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-31d99ee83f59125981322da4d05a0e08a83d0508da81bf5c3363c83ba451f38f"></a>

## routes.simple_route.advanced_options.csrf_policy.custom_domain_list — routes.simple_route.advanced_options.csrf_policy.custom_domain_list / 5a5e4ca8d34d / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.csrf_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-60dfdbdf9bf9617f1060c279c98ad547e8cc931bfac5c03c99af2f86e586218b)
- routes.simple_route.advanced_options.csrf_policy.custom_domain_list

<a id="canonical-160b664804f947475fa20f61110bef80f71bd78fefd7ad3ab7d19421964566b4"></a>

Type: `"single"`. Computed.

List of domain names used for Host header matching.

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

<a id="canonical-25bb004ac771ab286e45c2d26b670c04ce82fd94e216bf1536efc97c6926d9d1"></a>

## Direct properties — routes.simple_route.advanced_options.csrf_policy.custom_domain_list / 5a5e4ca8d34d / 3

<a id="canonical-dc1818baeb0c9e1c073b25ffa146f09b397de913a112318d15e15572236b09e9"></a>

<a id="canonical-04e246ec38d07fe9e040e2d9e14a10e19db3028e155db6d68825f194c46fab4c"></a>

## domains property — routes.simple_route.advanced_options.csrf_policy.custom_domain_list / 5a5e4ca8d34d / 4

Type: `["list", "string"]`. Computed.

List of domain names that will be matched to loadbalancer. These domains are not used for SNI match.
Wildcard names are supported in the suffix or prefix form.

Upstream description:

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "256",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.items.string.vh_domain": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-9e4b3c78958d6edaccd7b191f69a22b7756dfea68d8d0847b0be57091ed01d22"></a>

## Next pages — routes.simple_route.advanced_options.csrf_policy.custom_domain_list / 5a5e4ca8d34d / 5

- [routes.simple_route.advanced_options.csrf_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-60dfdbdf9bf9617f1060c279c98ad547e8cc931bfac5c03c99af2f86e586218b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-62edfac03b5989488e5de4c1d5c48704938781d69afaea31c91138e285909bb9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-25ebdad1a36a61037956882c3a148cd0131ac5caad082e99388d1337a20952c1"></a>

## routes.simple_route.advanced_options.csrf_policy.disabled — routes.simple_route.advanced_options.csrf_policy.disabled / dc82e56d7733 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.csrf_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-60dfdbdf9bf9617f1060c279c98ad547e8cc931bfac5c03c99af2f86e586218b)
- routes.simple_route.advanced_options.csrf_policy.disabled

<a id="canonical-38464c737d700d25c027efc7761da67b8d517c48959c256a0e1644157da14f06"></a>

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

<a id="canonical-d623c76d64732a5828afa95b035377b9ee6033c0f67325e67215a6ad9babf66c"></a>

## Direct properties — routes.simple_route.advanced_options.csrf_policy.disabled / dc82e56d7733 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c9e07c4f0d323e421b107254ec0a20f506243e5f0abf21739f1c658f1fa455dd"></a>

## Next pages — routes.simple_route.advanced_options.csrf_policy.disabled / dc82e56d7733 / 4

- [routes.simple_route.advanced_options.csrf_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-60dfdbdf9bf9617f1060c279c98ad547e8cc931bfac5c03c99af2f86e586218b)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-7ccc0016d13300c87cf9ac20b8ad1096a47ed6e6ab41f976be9b60772e6e9e8d"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-46bddde162211a0763732ad44c0b5048842724c350b1101e1a08a29cf439da37"></a>

## routes.simple_route.advanced_options.default_retry_policy — routes.simple_route.advanced_options.default_retry_policy / d070cf0fd436 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.default_retry_policy

<a id="canonical-b89ad1f3cf8ff851838d81be0cf0d25f711443959b66237acbda6402c4b6ab79"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-a6963001cf12f28a461155d03255fec28b74a06a18b9d9c08e0455e08a3b8326"></a>

## Direct properties — routes.simple_route.advanced_options.default_retry_policy / d070cf0fd436 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-116ab93c95b68bc4521bbeb06be26adbe84ab1b2c288873732427f94bbf8ecb1"></a>

## Next pages — routes.simple_route.advanced_options.default_retry_policy / d070cf0fd436 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1e184ac2707318ef658f31d8a200477233ad6e840aa5a2932f952538c1bca0aa"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4893a8cc6f68635a8d08b0fb43304cf88103c66986cdcab6b3fb318a56107b2e"></a>

## routes.simple_route.advanced_options.disable_mirroring — routes.simple_route.advanced_options.disable_mirroring / 635084abe65a / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.disable_mirroring

<a id="canonical-3e50cb03ee097e2759c89bba97a590bb960b5698d77fbf9e52e33725b46c6f12"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable mirroring.

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

<a id="canonical-5577392b72772471a81d39ba6ba674ced7870cd7b876f63551eaa9c90930a663"></a>

## Direct properties — routes.simple_route.advanced_options.disable_mirroring / 635084abe65a / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-358aa166f8b755465a5dbaca35c462c6bb24ebb4eae8d8a83670c8d04e3f9420"></a>

## Next pages — routes.simple_route.advanced_options.disable_mirroring / 635084abe65a / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-33f80f837020152d3321818fb412613893962d256b640072c5b0e8b941ddf76b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b2027d19fc0c86f7aca0e155611b04550c858792974052fb1c086fcd3e342771"></a>

## routes.simple_route.advanced_options.disable_prefix_rewrite — routes.simple_route.advanced_options.disable_prefix_rewrite / 37f77f105fc9 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.disable_prefix_rewrite

<a id="canonical-153c057a4f2ccb4c92d0df8ed872b13a530733d5c0bdda81f39e4aa4da08f02e"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable prefix rewrite.

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

<a id="canonical-20ce2f4a61fa462400b6787f607b3d89ba8111acdfec82485412038f071e1c44"></a>

## Direct properties — routes.simple_route.advanced_options.disable_prefix_rewrite / 37f77f105fc9 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-e1f36cb6e7a6f56870cf3e9ac15abf6dda334034d1fa4f175cbe9be9e5bc65ea"></a>

## Next pages — routes.simple_route.advanced_options.disable_prefix_rewrite / 37f77f105fc9 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-ab26f3a361f1791f0aa8e26c6cdf23037dcff177be9dae6225d312a921239c3b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1670dc1db99c2bba4734b8c8897f42d768bf21fc0e7bb2450c3d5bacdc80982f"></a>

## routes.simple_route.advanced_options.disable_spdy — routes.simple_route.advanced_options.disable_spdy / 403d6e60d8ef / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.disable_spdy

<a id="canonical-aed1c1c68ea83b7b4b05d6afd4764ac6b122226e8b08021fc7dcf2d484ec4728"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable spdy.

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

<a id="canonical-fb5808d97e651a7f6378c3abec01aab081991e830a5ba68fdfd76b4d57216d26"></a>

## Direct properties — routes.simple_route.advanced_options.disable_spdy / 403d6e60d8ef / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-d99192ae63cd82a5da69d051f43c44207fbe03209cea1504597bbc83a24e680c"></a>

## Next pages — routes.simple_route.advanced_options.disable_spdy / 403d6e60d8ef / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-64c28d52db0e7e954219acb1525f4209d51832102ab43625827897696852e095"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fd890b3b438a9d9d960f7c322c79f7033de79e32a4212314327b0ee68d9a7549"></a>

## routes.simple_route.advanced_options.disable_waf — routes.simple_route.advanced_options.disable_waf / a36c024f64e2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.disable_waf

<a id="canonical-19d9b18360fb772e87b467976bbafbedef2a15b67b7be6fc2cdb414d5fa203c4"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf.

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

<a id="canonical-4f8ecec31dfa22418456cacc0886356242b5d54eb5bb763b489e5df7ce2ad0f1"></a>

## Direct properties — routes.simple_route.advanced_options.disable_waf / a36c024f64e2 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2b32154bbdff991f527850f861b27340c33268c998e1c6656f021c9f937782ee"></a>

## Next pages — routes.simple_route.advanced_options.disable_waf / a36c024f64e2 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1c3234e93b7a6e2eaeb38762b76dc7d09712c63d6d88f562bf7ed515aeba30ac"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-47dfcdab24dd7d4544783b08814e619e73287358a42ec446c29a8279d1503719"></a>

## routes.simple_route.advanced_options.disable_web_socket_config — routes.simple_route.advanced_options.disable_web_socket_config / c35f3202e5a8 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.disable_web_socket_config

<a id="canonical-abdbad02ee7c29d73d3293144c7d8430b87d56392fbfbb5ab8e8ae10b29a940e"></a>

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

<a id="canonical-17e24deb5c705b9843880ba890d25f64e231769e69ac6f402667fa359baa5a53"></a>

## Direct properties — routes.simple_route.advanced_options.disable_web_socket_config / c35f3202e5a8 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-a30c57c3622a99ace1ec938658cf001d195d069093e1109f06ab36af5c6644c3"></a>

## Next pages — routes.simple_route.advanced_options.disable_web_socket_config / c35f3202e5a8 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-a3a2380a30e85754a58a4815bfb5a8ee66ba994bacecd1dd6d6de3e1e9177dc4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9c4f7a3f1032ad2522b70c7745442d9385597920678b59bbfd578462775e49ac"></a>

## routes.simple_route.advanced_options.do_not_retract_cluster — routes.simple_route.advanced_options.do_not_retract_cluster / 9922b1c36926 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.do_not_retract_cluster

<a id="canonical-c1107cd0367354a54f38c678a47e04bfd9965980d6e825e2a4f3fd9dd807b163"></a>

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

<a id="canonical-22ab654ef3cf0311dd69e59d2478d043f5c1948746b8ecf09c090c033810f659"></a>

## Direct properties — routes.simple_route.advanced_options.do_not_retract_cluster / 9922b1c36926 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2f19768dacc46eaab90ead381355236c4aaacb04d8c50032478b0d13a4823ce6"></a>

## Next pages — routes.simple_route.advanced_options.do_not_retract_cluster / 9922b1c36926 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-0b87e7a126800e2b39cb6e8f205aaca84e5cbdedab1fa6ec45862a01bb5efe99"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6dea5d44f7c794b3fbd2dbf2f8c758916616e67963f9d26a1fa156a646409386"></a>

## routes.simple_route.advanced_options.enable_spdy — routes.simple_route.advanced_options.enable_spdy / 2859d873d8c4 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.enable_spdy

<a id="canonical-d0a178a0509c3d2b08167d67f38557e0e7f516916b3f2d45bfad842cb0f9f775"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable spdy.

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

<a id="canonical-ef58dc09feb755bb9dd12c8b172fae6858dc00ef980ca6a2d56fd30ee1c319b0"></a>

## Direct properties — routes.simple_route.advanced_options.enable_spdy / 2859d873d8c4 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-8351c0ec9a998bc515a78f4d51e6c558c5ac2b91b6e83e1cc6f64963e983cd54"></a>

## Next pages — routes.simple_route.advanced_options.enable_spdy / 2859d873d8c4 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-2df64617e248296615238f6a69b0c1b3ab5e8422f1c4d9878c534d8b84a9a80e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-caadf99f35cb11e69400fb6414091dbfe29185ca5085d7e781a5ab2f96240441"></a>

## routes.simple_route.advanced_options.endpoint_subsets — routes.simple_route.advanced_options.endpoint_subsets / 0842b2568b5f / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.endpoint_subsets

<a id="canonical-001c32121782c431093dd8f162bfd91541e95ca618ab37c3eb9a1775a9939e07"></a>

Type: `"single"`. Computed.

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer For origin servers which are discovered in K8s or Consul..

Upstream description:

Upstream origin pool may be configured to divide its origin servers into subsets based on metadata
attached to the origin servers. Routes may then specify the metadata that a endpoint must match in
order to be selected by the load balancer

For origin servers which are discovered in K8s or Consul cluster, the label of the service is merged
with endpoint's labels. In case of Consul, the label is derived from the "Tag" field. For labels
that are common between configured endpoint and discovered service, labels from discovered service
takes precedence.

List of key-value pairs that will be used as matching metadata. Only those origin servers of
upstream origin pool which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.max_pairs": "16"
  }
}
```

<a id="canonical-3f1cc5d89ec26a7a468f297754d4ef263633a50a411c24da35e5fc111d73ff2c"></a>

## Direct properties — routes.simple_route.advanced_options.endpoint_subsets / 0842b2568b5f / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-4e9a519e1caa07b63ff52e5c63b5433eedfefecd7f7a7968ca32d53e25e02933"></a>

## Next pages — routes.simple_route.advanced_options.endpoint_subsets / 0842b2568b5f / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c77539f53ba7dcaacbfeeddf31daa1e7ebf3fe61e95f992380c7ee1b332375de"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b8c40ed8627bcbdcdf6688a63cb3dda83879d1f0c745f18c487a1ba6f3d31bde"></a>

## routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection — routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection / 347302d935a1 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection

<a id="canonical-dc0154824c4d70b87fa446e8e14acb5cfbafa0d7291429c0c593f0ae13437083"></a>

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

<a id="canonical-35a9927fdc70c69f141039324c2a8dd956c6af3f535bd5d1417dbc897865148a"></a>

## Direct properties — routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection / 347302d935a1 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-17c5907f23e7e14bade6de6fa9cb049e194a6b9c035db7d0d0fecc656ee03769"></a>

## Next pages — routes.simple_route.advanced_options.inherited_bot_defense_javascript_injection / 347302d935a1 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-151e4e141e0b8478564e09810fd4293346787906af422e8900d0efc3844eb246"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5370f046d7b47549cb9e74eea8acb39d9cb775870a774baac77f0c6d0539b154"></a>

## routes.simple_route.advanced_options.inherited_waf — routes.simple_route.advanced_options.inherited_waf / 4504e1c7206e / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.inherited_waf

<a id="canonical-f19f07304602a375998540b8e4a6daf4029e528efb191661faed637166cb596a"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherited waf.

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

<a id="canonical-ad94a122030260ec9c061ba1f7f306bd7188247b4cbb6d67efd74539155a57b5"></a>

## Direct properties — routes.simple_route.advanced_options.inherited_waf / 4504e1c7206e / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3d5a6e392f2673fca76ceb7740ae7171b0666f988b904ffdd350237588a6c8bd"></a>

## Next pages — routes.simple_route.advanced_options.inherited_waf / 4504e1c7206e / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-19688b6f66df4309a4e12e0c6d13b5df0c5761bb5a169371bacd566e9a161815"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c60d3d68ea76c748a0420ac8dad030f442af7cbe6ffab9f381ed938a36e4287e"></a>

## routes.simple_route.advanced_options.inherited_waf_exclusion — routes.simple_route.advanced_options.inherited_waf_exclusion / 0f184edda782 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.inherited_waf_exclusion

<a id="canonical-559cd46c61f9e7066e6752aeeaba28e6d4869bcd8a8882398b9ad82ed6538d2f"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherited waf exclusion.

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

<a id="canonical-dbe41cad7b0d32107364a01b400fa395a1a0d10f92753df8373f40579fca9b28"></a>

## Direct properties — routes.simple_route.advanced_options.inherited_waf_exclusion / 0f184edda782 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-9fea26472c66f6c991cdf3fd8aa3066cb71b04b5fcf03af5a217c88ee1cfd6c4"></a>

## Next pages — routes.simple_route.advanced_options.inherited_waf_exclusion / 0f184edda782 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1f769c35b2f994587ec953a41759d45ee49a616e417b0becf45983a71a456ef5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a9e414a64394bb965cafd5052facbfa8bd2825771e43ba9023d0db08720e064c"></a>

## routes.simple_route.advanced_options.mirror_policy — routes.simple_route.advanced_options.mirror_policy / 61fe967f9fe2 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.mirror_policy

<a id="canonical-b75d0eefe3229aa7b180d79116399392ea80fe4d8cd6ca9b86099c2f9ae86e42"></a>

Type: `"single"`. Computed.

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow
origin..

Upstream description:

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
"fire and forget", meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow origin
pool making this feature useful for testing and troubleshooting.

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

<a id="canonical-52abd71d71c712db6f39751a01f08bc1676f2a42e7152cff60a2c18fd900287a"></a>

## Direct properties — routes.simple_route.advanced_options.mirror_policy / 61fe967f9fe2 / 3

- [origin_pool](data-sources--http_loadbalancer--reference--group-023.md#canonical-724e681ebbef6c6db761549b7419a28556d2a25c61841913a288572c6b3a97cc): complete subsection reference.

- [percent](data-sources--http_loadbalancer--reference--group-023.md#canonical-46a99fdbd457ac3a92f95702f4e9b1ae7b3464f703124dae045113759a77e3d3): complete subsection reference.

<a id="canonical-9a5a06526aff216fb97a6e6fc7afda9d567796f18f17f771738534f2790d946a"></a>

## Next pages — routes.simple_route.advanced_options.mirror_policy / 61fe967f9fe2 / 4

- [routes.simple_route.advanced_options.mirror_policy.origin_pool](data-sources--http_loadbalancer--reference--group-023.md#canonical-724e681ebbef6c6db761549b7419a28556d2a25c61841913a288572c6b3a97cc)
- [routes.simple_route.advanced_options.mirror_policy.percent](data-sources--http_loadbalancer--reference--group-023.md#canonical-46a99fdbd457ac3a92f95702f4e9b1ae7b3464f703124dae045113759a77e3d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-724e681ebbef6c6db761549b7419a28556d2a25c61841913a288572c6b3a97cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-fe264161f77e6d994a308c51e9d70f2ea8ebd6a82e5417e7174eaa108a09cda5"></a>

## routes.simple_route.advanced_options.mirror_policy.origin_pool — routes.simple_route.advanced_options.mirror_policy.origin_pool / 60dd0486dd11 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.mirror_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f769c35b2f994587ec953a41759d45ee49a616e417b0becf45983a71a456ef5)
- routes.simple_route.advanced_options.mirror_policy.origin_pool

<a id="canonical-650711dc5c3a37a4fb338f97a2203aea0b241976058b09bac3a2fd6ed8930f84"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-c62e36699e04eae1fb25e37eeba10f7e43cc39858fff9ce719154a1e9417149b"></a>

## Direct properties — routes.simple_route.advanced_options.mirror_policy.origin_pool / 60dd0486dd11 / 3

<a id="canonical-9e643bca4f5a2f06164fec2531898b68988fa28e203c6ffcaeafe26a75955583"></a>

<a id="canonical-886dcea9bddd413c6f30f76394d7ceb4686c525668ec83ba5478c8beee5694d4"></a>

## name property — routes.simple_route.advanced_options.mirror_policy.origin_pool / 60dd0486dd11 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2b28d86d6e97fdf2fffef9d485ba7ba53ca8c072c245cfcc11fa130072d3753d"></a>

<a id="canonical-1654eef28569a2505309fb0b38aa0f48dc34fc73ecfc9c36730d8b6e540d1c89"></a>

## namespace property — routes.simple_route.advanced_options.mirror_policy.origin_pool / 60dd0486dd11 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-cfd4f661f057376e68e7da4b635a40d8f8d4d1e3ddbe7dfd5d96e6ceccb3e76b"></a>

<a id="canonical-c50c291fc6d0d48aea4019cfaaee22c271965a309252bc2cdea29cc5e14bfd6c"></a>

## tenant property — routes.simple_route.advanced_options.mirror_policy.origin_pool / 60dd0486dd11 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-7cbc35da83e5cf7955229ddeb2b68ac428200a126c519153cfa3e430f2b5ea94"></a>

## Next pages — routes.simple_route.advanced_options.mirror_policy.origin_pool / 60dd0486dd11 / 7

- [routes.simple_route.advanced_options.mirror_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f769c35b2f994587ec953a41759d45ee49a616e417b0becf45983a71a456ef5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-46a99fdbd457ac3a92f95702f4e9b1ae7b3464f703124dae045113759a77e3d3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f1f9349eb6b39555fde8630bed20984b883ed363ed8816c40c4c5bcb4e1ec227"></a>

## routes.simple_route.advanced_options.mirror_policy.percent — routes.simple_route.advanced_options.mirror_policy.percent / c85a1a463990 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.mirror_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f769c35b2f994587ec953a41759d45ee49a616e417b0becf45983a71a456ef5)
- routes.simple_route.advanced_options.mirror_policy.percent

<a id="canonical-93db97643737a18271b746f5b2cec7fa59dc0c6ca0754fa654b47058f8af0a6e"></a>

Type: `"single"`. Computed.

Fraction used where sampling percentages are needed. Example sampled requests.

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

<a id="canonical-9f11347541cca34c11f1f16a6812b833b3d2e8392b1da133f6bbdf78a22bce34"></a>

## Direct properties — routes.simple_route.advanced_options.mirror_policy.percent / c85a1a463990 / 3

<a id="canonical-029cd336a77efc614a1de0d0cb6d729a97039115e19c62ea29d8de2c11136c72"></a>

<a id="canonical-8ae5e814d3ac99fa5087ee6b2625185e1c02f053a9725ae1996bb3aec5516ed6"></a>

## denominator property — routes.simple_route.advanced_options.mirror_policy.percent / c85a1a463990 / 4

Type: `"string"`. Computed.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Upstream description:

Denominator used in fraction where sampling percentages are needed. Example sampled requests

Use hundred as denominator Use ten thousand as denominator Use million as denominator.

Receipt-pinned upstream constraints:

```json
{
  "default": "HUNDRED",
  "enum": [
    "HUNDRED",
    "TEN_THOUSAND",
    "MILLION"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-f8af8d183cbb410ffd0742f260a042d7348b9051ad4b0331927405d70dfae4de"></a>

<a id="canonical-092923b34eb3c82ee069a454a4c94e5757bd9b74f1d88ae93d2c7bd1e70209f7"></a>

## numerator property — routes.simple_route.advanced_options.mirror_policy.percent / c85a1a463990 / 5

Type: `"number"`. Computed.

Sampled parts per denominator. If denominator was 10000, then value of 5 will be 5 in 10000.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-bcc7a44efb3d253d295c562fd266868fe9f2bc735aed317859293d645dd56867"></a>

## Next pages — routes.simple_route.advanced_options.mirror_policy.percent / c85a1a463990 / 6

- [routes.simple_route.advanced_options.mirror_policy](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f769c35b2f994587ec953a41759d45ee49a616e417b0becf45983a71a456ef5)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-d694da319839efa801fda458340d5a2b9ddb71372337297699448b6d15476113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-876cddb058f3e9fdc7a512abd700c90e240a38ef3346ef00a12cb914d697b073"></a>

## routes.simple_route.advanced_options.no_retry_policy — routes.simple_route.advanced_options.no_retry_policy / a89380bcea06 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.no_retry_policy

<a id="canonical-ec108fc1550761b0e8cd7b7b4ddb0370362895cd1772f2ebb0c6af947a9a8271"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-431094914134c7a1a76b202be41eb4525fec2c84633da4d1ddae8f958c184ab9"></a>

## Direct properties — routes.simple_route.advanced_options.no_retry_policy / a89380bcea06 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-c908408e2f756c0255e90fdbf04a2baefcb2872f8a84317387de0256727802d3"></a>

## Next pages — routes.simple_route.advanced_options.no_retry_policy / a89380bcea06 / 4

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-c0d0fc65461843d23c8820bb8f9422b6c455947b0710550459d44cb396fa94b5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0280eff6f4fed19832f3edc4ce86431da20578ab2f1b121409129e26b2ed7290"></a>

## routes.simple_route.advanced_options.regex_rewrite — routes.simple_route.advanced_options.regex_rewrite / dcd9e4174763 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.regex_rewrite

<a id="canonical-32d21b023b2dbd5fc7edd9591881a496837b212d5b2b4d2a8407ef95415841ba"></a>

Type: `"single"`. Computed.

RegexMatchRewrite describes how to match a string and then produce a new string using a regular
expression and a substitution string.

Upstream description:

RegexMatchRewrite describes how to match a string and then produce a new string using a regular
expression and a substitution string.

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

<a id="canonical-a3f310616b49099ff85620b215ae93b017924947cb7d3ab194dd89fcbd45c23b"></a>

## Direct properties — routes.simple_route.advanced_options.regex_rewrite / dcd9e4174763 / 3

<a id="canonical-1d9404835a2533af59ceeb94bb28c24a8eabcadc4a98b6ad71299affc90af20c"></a>

<a id="canonical-698181dc5e61819fcc4e41ae0164d5e225d38c55c24f3e9389e82a2b51ecb61d"></a>

## pattern property — routes.simple_route.advanced_options.regex_rewrite / dcd9e4174763 / 4

Type: `"string"`. Computed.

The regular expression used to find portions of a string that should be replaced.

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

<a id="canonical-42b086ec59482abc09e83ff341ca2f285efa57c7b889e2f21b8a0c3782040e9b"></a>

<a id="canonical-10653d8e96b0422714e2edfe4fba616e0659db96a8f89459cd5fa82cac86853a"></a>

## substitution property — routes.simple_route.advanced_options.regex_rewrite / dcd9e4174763 / 5

Type: `"string"`. Computed.

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

Upstream description:

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-ff5c30d6ca1b4bca62e366055adf6446435c45ce686d7afaafabc8f52984371a"></a>

## Next pages — routes.simple_route.advanced_options.regex_rewrite / dcd9e4174763 / 6

- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-1f3a8a13d10671a490d330b6c70591cbd941c941b58f82b19df0fe4c42e76852"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f72ffe9564372889e084763d2c197489ec4ef93b0e6bd8c21da0ee68f394fd22"></a>

## routes.simple_route.advanced_options.request_cookies_to_add — routes.simple_route.advanced_options.request_cookies_to_add / 7feb09fa1aee / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- routes.simple_route.advanced_options.request_cookies_to_add

<a id="canonical-a8c4cf8ab8518648a8578c5da9951fe674059bea83b89e021c700f7e87c42164"></a>

Type: `"list"`. Computed.

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

Upstream description:

Cookies are key-value pairs to be added to HTTP request being routed towards upstream. Cookies
specified at this level are applied after cookies from matched Route are applied.

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

<a id="canonical-2a3452efc8ea69758101977e253c71e3c4e5d3e91d6169a516126c1f501df1bd"></a>

## Direct properties — routes.simple_route.advanced_options.request_cookies_to_add / 7feb09fa1aee / 3

<a id="canonical-d922fca5e1c0a7fce09abc1a6aa7ab3aa99384b0134a15a62cbcd1a0971c4b36"></a>

<a id="canonical-2c1ae9246af07cdc71a653679cf599f0dce203e82e21f9bbb85bb85f9f6fb357"></a>

## name property — routes.simple_route.advanced_options.request_cookies_to_add / 7feb09fa1aee / 4

Type: `"string"`. Computed.

Name. Name of the cookie in Cookie header.

Upstream description:

Name of the cookie in Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
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
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.cookie_name": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-b9d6f4cd2d4cfa14f204f8e76d247594d60226b2b2fc6d5226e7719f91783c69"></a>

<a id="canonical-dc81a4453620be4999cbe97f81a2e0a6ecceea0cf158a1bf7f072655e2749096"></a>

## overwrite property — routes.simple_route.advanced_options.request_cookies_to_add / 7feb09fa1aee / 5

Type: `"bool"`. Computed.

Should the value be overwritten? If true, the value is overwritten to existing values. not
overwrite. Defaults to \`do\`.

Upstream description:

Should the value be overwritten? If true, the value is overwritten to existing values. Default value
is do not overwrite.

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

- [secret_value](data-sources--http_loadbalancer--reference--group-023.md#canonical-60ac953e6fa03b699f98b4339db942385909c5a499261287d7eceba236c45b09): complete subsection reference.

<a id="canonical-771b7f1f2bb355bbc5be2c427d8532d65133cf9278b180f6764db41353a4cc12"></a>

<a id="canonical-686c334bd4296ae40382fc28fd3469eea3921b9efdc0acb2b5b91c8ace7d8c2d"></a>

## value property — routes.simple_route.advanced_options.request_cookies_to_add / 7feb09fa1aee / 6

Type: `"string"`. Computed.

Exclusive with \[secret\_value\] Value of the Cookie header.

Upstream description:

Exclusive with \[secret\_value\] Value of the Cookie header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-256f1d16e3b649169cab221cc1ae28e3daeab410ea02a31d9678675167593fe3"></a>

## Next pages — routes.simple_route.advanced_options.request_cookies_to_add / 7feb09fa1aee / 7

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-023.md#canonical-60ac953e6fa03b699f98b4339db942385909c5a499261287d7eceba236c45b09)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-60ac953e6fa03b699f98b4339db942385909c5a499261287d7eceba236c45b09"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-9f6cec9a6daf382874eaf28b2ca24b124122fbea368cabff39608934291b3a9a"></a>

## routes.simple_route.advanced_options.request_cookies_to_add.secret_value — routes.simple_route.advanced_options.request_cookies_to_add.secret_value / b6d7dccbfd01 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.request_cookies_to_add](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f3a8a13d10671a490d330b6c70591cbd941c941b58f82b19df0fe4c42e76852)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value

<a id="canonical-c94fdd0e1cf6e090176102adcc3d0ef547589417699bf618726a97fbd5e05ce6"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

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

<a id="canonical-2fea22fe8f07f7074770097db77edb315eb944fe52d6fbecc7d6c503598e92c5"></a>

## Direct properties — routes.simple_route.advanced_options.request_cookies_to_add.secret_value / b6d7dccbfd01 / 3

- [blindfold_secret_info](data-sources--http_loadbalancer--reference--group-023.md#canonical-70f53161cfff1bdbb47faef78245abcfb479619fbc39a8fab8def2b5de882784): complete subsection reference.

- [clear_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-390d69dcd8d4a77f3e202a8691d36a9b88f36e46e52c49adad332a3cfbdd4251): complete subsection reference.

<a id="canonical-56d6b9d60d156fb7104c0b07d3f8a4667051db5f643e86201866c002f5999d1c"></a>

## Next pages — routes.simple_route.advanced_options.request_cookies_to_add.secret_value / b6d7dccbfd01 / 4

- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info](data-sources--http_loadbalancer--reference--group-023.md#canonical-70f53161cfff1bdbb47faef78245abcfb479619fbc39a8fab8def2b5de882784)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value.clear_secret_info](data-sources--http_loadbalancer--reference--group-024.md#canonical-390d69dcd8d4a77f3e202a8691d36a9b88f36e46e52c49adad332a3cfbdd4251)
- [routes.simple_route.advanced_options.request_cookies_to_add](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f3a8a13d10671a490d330b6c70591cbd941c941b58f82b19df0fe4c42e76852)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)

<a id="canonical-70f53161cfff1bdbb47faef78245abcfb479619fbc39a8fab8def2b5de882784"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1346aaa8332a856a44e43a96f15b883a7d2ff189126bdcab776a207fd51616e6"></a>

## routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / 233f051eb79c / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-b9265d3d0725df0ed2c4fa4994941446ef82cd791555a70b1901a5649c1cce80)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-58128a4ee6bcaa29c55314ef6e094489064f6b317a0ae1e19d044979c56872b8)
- [routes](data-sources--http_loadbalancer--reference--group-022.md#canonical-ee9f4251c12a0cc9ca02d6a6f1516fdacdc974707099302cc5b877040f3b35c0)
- [routes.simple_route](data-sources--http_loadbalancer--reference--group-023.md#canonical-479b1f6f287bfb3ed88ffb35231975f9f9b685b55a0d14e312c962fa5a0232d3)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--reference--group-023.md#canonical-93c39b6bf7d436a3f4e2ce51e4a4ada25751d61de48f7a50a2c1db9115e3ea4f)
- [routes.simple_route.advanced_options.request_cookies_to_add](data-sources--http_loadbalancer--reference--group-023.md#canonical-1f3a8a13d10671a490d330b6c70591cbd941c941b58f82b19df0fe4c42e76852)
- [routes.simple_route.advanced_options.request_cookies_to_add.secret_value](data-sources--http_loadbalancer--reference--group-023.md#canonical-60ac953e6fa03b699f98b4339db942385909c5a499261287d7eceba236c45b09)
- routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfold_secret_info

<a id="canonical-fe948a4e2d4b9d82df635c807369b92ac7c7bba3d7a7430a2acb6ddb93de34df"></a>

Type: `"single"`. Computed.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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

<a id="canonical-c96c243100860abcdf1dca238fcaa0e9476d7f4877703cbd7742f34ecd752d0f"></a>

## Direct properties — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / 233f051eb79c / 3

<a id="canonical-f8a2c75bdae67968a005bc83b18620d32f00a7d119d651c8a39f7db8dd804880"></a>

<a id="canonical-ad5b187caed078fee8b2bc1be162405d89bc96cccede4c77098748bffaea3d0f"></a>

## decryption_provider property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / 233f051eb79c / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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
  }
}
```

<a id="canonical-9626ccd03a6817bc254484cf4c7b1e34c6fd8ed1e0606f0ee21d92681cf1cdc7"></a>

<a id="canonical-9cca23fe4df99b0e1b4c260888b897cdc5e0491841a2fc675d1adf8fbc78acb5"></a>

## location property — routes.simple_route.advanced_options.request_cookies_to_add.secret_value.blindfo / 233f051eb79c / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-a4448016eda635d934776d8e8a03ea8a395137f077d352693aaf89cc183d277d"></a>
