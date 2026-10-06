---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-3203333320300320-1001303300122003-1302121103333112-2100212201013022-3332320200103031-0011130011222123-0132200132321021-1031303233123221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_headers_to_add.secret_value.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_headers_to_add](resources--route--reference--group-002.md#canonical-1033033030111003-2022222120022003-1123301032332002-1313030103020021-3112122133001003-0321311203111100-2123212000030110-2132031210122220)
- [routes.response_headers_to_add.secret_value](resources--route--reference--group-002.md#canonical-3010111232312331-2003220221321131-0322032331220012-3211121312110333-1002012232113320-1010130133031122-0013112202333213-0232202132220103)
- routes.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1332301321300232-0033033312301202-3300011230112200-3322100333133323-3122312110121023-1323311120232300-2333020000202212-0301120130332001"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201130113321210-2112121022201300-1131100021101133-3030003022211102-3330311002001023-2302033223312022-0301113112111202-0033332313230112"></a>

### Direct properties for `routes.response_headers_to_add.secret_value.blindfold_secret_info`

<a id="canonical-0133130110130102-3311201331123012-1021303033003210-2022323123022233-3103003303333220-0231233231211212-2012123121323222-1103120300310301"></a>

#### `routes.response_headers_to_add.secret_value.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3300223022133313-0000020321001221-1032203231013200-2102212221113210-2101231013130322-2301013002332322-1032220211211321-3021213302331013"></a>

<a id="canonical-0301220322200301-1233302102300201-1120201113130023-0001130211012101-0120103230131021-2031232111000022-2111122001013011-3233202013131103"></a>

#### `routes.response_headers_to_add.secret_value.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2302200200012210-2113003013310212-0133000132100012-1031033101020122-3200323021312002-3031122220023131-2132230200331132-2333230100231022"></a>

<a id="canonical-1312113112312020-2132201302033020-1201213012103332-2003101120032333-0112233110301211-2121320021311301-1330231221123012-0103330322233102"></a>

#### `routes.response_headers_to_add.secret_value.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2031221013021230-3200000002021300-0320211122130203-3022021200310300-3220131112131311-2200200130031121-1031030211313012-2333031122022233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.response_headers_to_add.secret_value.clear_secret_info` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.response_headers_to_add](resources--route--reference--group-002.md#canonical-1033033030111003-2022222120022003-1123301032332002-1313030103020021-3112122133001003-0321311203111100-2123212000030110-2132031210122220)
- [routes.response_headers_to_add.secret_value](resources--route--reference--group-002.md#canonical-3010111232312331-2003220221321131-0322032331220012-3211121312110333-1002012232113320-1010130133031122-0013112202333213-0232202132220103)
- routes.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-1331312122013013-1330201332131122-0222020212331112-2022112310002112-0130110322213010-2101223300121211-2002213200103203-3301003332123232"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123012333013033-0021030032113300-3132323330013130-1313132320131203-3130300013003233-3210203030313201-1322220111222302-1223213110201200"></a>

### Direct properties for `routes.response_headers_to_add.secret_value.clear_secret_info`

<a id="canonical-0131001300012132-3022201030022021-0123312023221113-2112131110200210-2003303100011121-3312122011002020-0032220030232300-2030100013232311"></a>

#### `routes.response_headers_to_add.secret_value.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0223010103300110-0211201203103003-2012121323000313-3013103331332330-0211222120010132-2001231301011120-2211111322231213-0002220013303111"></a>

<a id="canonical-0100012101130332-3132032003020023-0332032310221220-0210020010203303-2333101020001112-2233310201220110-0302011122101222-2210030201333133"></a>

#### `routes.response_headers_to_add.secret_value.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.route_destination

<a id="canonical-3202133011020222-0010230313321311-0131111301313033-1310221232100321-0213231323010232-3301302020002202-3103333310010032-0320130202302201"></a>

Type: `"object"`. single nested block, Optional.

List of destination to choose if the route is match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("destinations"),
  validators.ConflictingObjectAttributes("auto_host_rewrite",
    "host_rewrite"),
  validators.ConflictingObjectAttributes("do_not_retract_cluster",
    "retract_cluster"),
  validators.ConflictingObjectAttributes("prefix_rewrite",
    "regex_rewrite")}
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
  "x-ves-oneof-field-cluster_retract_choice": "[\"do_not_retract_cluster\",\"retract_cluster\"]",
  "x-ves-oneof-field-host_rewrite_params": "[\"auto_host_rewrite\",\"host_rewrite\"]",
  "x-ves-oneof-field-route_destination_rewrite": "[\"prefix_rewrite\",\"regex_rewrite\"]"
}
```

Terraform syntax:

```terraform
route_destination {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232120123000213-0012120220200012-2111233313030331-0122313113000233-0012111102110001-1011032322210222-2212000113120022-2030321321112011"></a>

### Direct properties for `routes.route_destination`

<a id="canonical-1020132210222313-3231102003332113-3110211221330212-0030001130323103-1011302201323323-2212011001130231-2231210021230010-3221133321011012"></a>

#### `routes.route_destination.auto_host_rewrite` property

Type: `"bool"`. Optional.

Exclusive with \[host\_rewrite\] Indicates that during forwarding, the host header will be swapped
with the hostname of the upstream host chosen by the cluster.

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

- [buffer_policy](resources--route--reference--group-003.md#canonical-1101330032333303-3102321110222222-0320031323312213-2030010101032331-0332221210202312-2022102132200133-3310103022110310-1230121322213333): complete subsection reference.

- [cors_policy](resources--route--reference--group-003.md#canonical-3020020023122211-2131012103232233-2103022303301200-1310232123121322-2023220300101201-3013230110022110-2123210133331133-0210102331213233): complete subsection reference.

- [csrf_policy](resources--route--reference--group-003.md#canonical-1101321211132233-2202003002310130-3123211010001032-2023310312100303-3310321001200211-2001220013330302-2200133210032112-2213001112022232): complete subsection reference.

- [destinations](resources--route--reference--group-003.md#canonical-2033310032010003-2013013212023023-3331003021220120-3211312230110102-0332232123000011-2322023233233003-2010313220111220-3203133211033012): complete subsection reference.

- [do_not_retract_cluster](resources--route--reference--group-003.md#canonical-2232131110032213-2031333201002023-1223031002323001-0020002002100232-1113320131001233-1033123003220021-0030130030111321-2323233220010332): complete subsection reference.

- [endpoint_subsets](resources--route--reference--group-003.md#canonical-1111031211001032-1233310121003212-2023210031222203-2221132210120000-3013102123000323-3111031112232231-1202110122322030-1221020103030213): complete subsection reference.

- [hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123): complete subsection reference.

<a id="canonical-0213120322322211-2323332001212230-3002223111223201-1101113121101231-3333221201301321-3200122111111020-3322032111111131-0232230332313301"></a>

<a id="canonical-2133213211001221-2112223002113033-0313112230033133-1203011002212222-3301020030031232-0232211202320312-0212100123130121-2113103323332013"></a>

#### `routes.route_destination.host_rewrite` property

Type: `"string"`. Optional.

Exclusive with \[auto\_host\_rewrite\] Indicates that during forwarding, the host header will be
swapped with this value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [mirror_policy](resources--route--reference--group-003.md#canonical-1312110103201222-3211323013120213-1220232003210122-1333011320200321-0302110311030231-3120313321230312-0300310313300210-1231213001201021): complete subsection reference.

<a id="canonical-3300231230103211-3131012203302321-2231333032013023-2032201232302133-1110332030311313-2233313102012022-1133202032013231-1003031021221222"></a>

<a id="canonical-0032311120102123-0011000021010100-2112111010232031-0203032031112011-3033222211312200-1202212102131322-2322212230233121-1233131331321020"></a>

#### `routes.route_destination.prefix_rewrite` property

Type: `"string"`. Optional.

Exclusive with \[regular expression\_rewrite\] prefix\_rewrite indicates that during forwarding, the matched
prefix (or path) should be swapped with its value. When using regular expression path matching, the entire path
(not including the query string) will be swapped with this value. This option allows application
URLs to be rooted at a different path from those exposed at the reverse proxy layer.

Example : gcSpec: routes: &#8203;- match: &#8203;- headers: \[\] path: prefix : /register/
query\_params: \[\] &#8203;- headers: \[\] path: prefix: /register query\_params: \[\]
routeDestination: prefixRewrite: "/" destinations: &#8203;- cluster: &#8203;- kind: cluster.object
uid: cluster-1

Having above entries in the config, requests to /register will be stripped to /, while requests to
/register/public will be stripped to /public.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2332133011003023-2203210112332003-2220001320113301-2011311223221131-1003131023130120-1303012222331311-0112223313131223-3011301200031022"></a>

<a id="canonical-0213111001203300-2310021302320130-2012211110311032-1333023311013030-0210132320011301-2331013302033120-1331330321312232-2032100112302222"></a>

#### `routes.route_destination.priority` property

Type: `"string"`. Optional.

\[Enum: DEFAULT|HIGH\] Priority routing for each request. Different connection pools are used based
on the priority selected for the request. Also, circuit-breaker configuration at destination cluster
is chosen based on selected priority. Possible values are \`DEFAULT\`, \`HIGH\`. Defaults to
\`DEFAULT\`.

Additional upstream details:

Priority routing for each request. Default routing mechanism High-Priority routing mechanism.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEFAULT",
    "HIGH"),
}
```

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

- [query_params](resources--route--reference--group-003.md#canonical-0210013123232110-1101203121320003-0232332230203020-2203202001132332-1123010132101312-1300303103231023-2212021131213210-0023212010013031): complete subsection reference.

- [regex_rewrite](resources--route--reference--group-003.md#canonical-1002132200332031-2112010113122321-3210320213210322-3203102103022111-0020102231301312-3020023031212111-1313020132121103-1031331331113022): complete subsection reference.

- [retract_cluster](resources--route--reference--group-003.md#canonical-1302302032030012-1303121333102133-0133131022330201-0212312130311012-1211001203310220-1221201023122013-2133021022130012-0333323020112031): complete subsection reference.

- [retry_policy](resources--route--reference--group-003.md#canonical-0301103202310121-1322001133231203-3330201331230333-2121120301200322-3222320330230320-1021310030201212-1120112023110002-2203110113210212): complete subsection reference.

- [spdy_config](resources--route--reference--group-003.md#canonical-1211112100111122-3213121213133231-1231320323033110-2010301220033103-1021202000330100-0023023012000230-0022012210003120-3103323301010233): complete subsection reference.

<a id="canonical-0023211312233333-0003323002031112-0021112010003130-1033120010002230-1222210230321312-1113203201320201-1202223102012200-3130121332201223"></a>

<a id="canonical-0003030121232022-3010332003132210-2232033210230333-0301313231033013-3231030030111311-3123131002201120-3210030322131010-2113012322203313"></a>

#### `routes.route_destination.timeout` property

Type: `"number"`. Optional.

Specifies the timeout for the route in milliseconds. This timeout includes all retries. For server
side streaming, configure this field with higher value or leave it un-configured for infinite
timeout.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 1800000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1800000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "1800000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "1800000"
  }
}
```

- [web_socket_config](resources--route--reference--group-003.md#canonical-0223313132331320-3312320001233000-1122103111302032-2311223023000012-0002221013031321-2022311222102223-3033032231010300-1312132322100022): complete subsection reference.

<a id="canonical-1101330032333303-3102321110222222-0320031323312213-2030010101032331-0332221210202312-2022102132200133-3310103022110310-1230121322213333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.buffer_policy` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.buffer_policy

<a id="canonical-3332132131303110-3303021323103203-0322131323020321-0231302332022120-1022312330333011-0232101300120021-2113302103113023-1101233121302130"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
buffer_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303100113100102-3322203020002113-0233312310323012-2233032301201313-0310212113033321-3112333113211113-3232023330331321-1113121131223323"></a>

### Direct properties for `routes.route_destination.buffer_policy`

<a id="canonical-0330123133230013-1002310231131330-3110113001002321-2113303001101220-1101201022212302-1100300121113233-3011130113322220-2322310221113322"></a>

#### `routes.route_destination.buffer_policy.disabled` property

Type: `"bool"`. Optional.

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

<a id="canonical-0133300333132100-1212310332103320-2300201203201202-1312231132213131-3111112232011112-0032203230311202-3212213102201010-1123200102023032"></a>

<a id="canonical-0301222131210220-0120300020301130-1002010022013331-0321232002033010-0203103330231100-3021310321123101-0021011232203210-1012111120010131"></a>

#### `routes.route_destination.buffer_policy.max_request_bytes` property

Type: `"number"`. Optional.

The maximum request size that the filter will buffer before the connection manager will stop
buffering and return a RequestEntityTooLarge (413) response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(10485760),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3020020023122211-2131012103232233-2103022303301200-1310232123121322-2023220300101201-3013230110022110-2123210133331133-0210102331213233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.cors_policy` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.cors_policy

<a id="canonical-1002001011310010-1321122311202020-2313123222111310-3320112133233220-2230111312030313-1101313021102132-0133231303220330-3031002200232210"></a>

Type: `"object"`. single nested block, Optional.

Cross-Origin Resource Sharing requests configuration specified at Virtual-host or Route level. Route
level configuration takes precedence.

An example of an Cross origin HTTP request GET /resources/public-data/ HTTP/1.1 Host: bar.other
User-Agent: Mozilla/5.0 (Macintosh; U; Intel macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130
Minefield/3.1b3pre Accept: text/HTML,application/xhtml+XML,application/XML;q=0.9,\*/\*;q=0.8
Accept-Language: en-us,en;q=0.5 Accept-Encoding: gzip,deflate Accept-Charset:
ISO-8859-1,utf-8;q=0.7,\*;q=0.7 Connection: keep-alive Referrer:
http&#58;//foo.example/examples/access-control/simplexsinvocation.HTML Origin:
http&#58;//foo.example

HTTP/1.1 200 OK Date: Mon, 01 Dec 2008 00:23:53 GMT Server: Apache/2.0.61
Access-Control-Allow-Origin: \* Keep-Alive: timeout=2, max=100 Connection: Keep-Alive
Transfer-Encoding: chunked Content-Type: application/XML

An example for cross origin HTTP OPTIONS request with Access-Control-Request-\* header

OPTIONS /resources/POST-here/ HTTP/1.1 Host: bar.other User-Agent: Mozilla/5.0 (Macintosh; U; Intel
macOS X 10.5; en-US; rv:1.9.1b3pre) Gecko/20081130 Minefield/3.1b3pre Accept:
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

Terraform syntax:

```terraform
cors_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011202233310012-3100021022030110-3300212102032131-0202020300121212-2101320021323032-0330222000301102-2332021333223102-0301130122231323"></a>

### Direct properties for `routes.route_destination.cors_policy`

<a id="canonical-3323033113301101-0302003112302303-2131012202212330-2023310300032320-1220202323210123-2000202301120212-2322221300120022-3101211022202133"></a>

#### `routes.route_destination.cors_policy.allow_credentials` property

Type: `"bool"`. Optional.

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

<a id="canonical-3220220220033200-2122010221110132-3233031032133231-0300322132303113-0131030302201232-1133112310311110-3022031332000103-3200131331022203"></a>

<a id="canonical-3021023213322301-0212101223001022-3302101013321011-1123131001100233-1303310131312112-1120111211313320-2100122031203232-3131031112021122"></a>

#### `routes.route_destination.cors_policy.allow_headers` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0230123213310113-3311223003320010-1000303203333030-1233202232313332-1331033031322313-0300103221131030-3002030102123100-1103102111111130"></a>

<a id="canonical-0133203132311110-3123002232330201-3300013220310302-0222021322111213-0233331333111221-1020322211320320-2033203002330132-3231110300031021"></a>

#### `routes.route_destination.cors_policy.allow_methods` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1110121302020313-0112330301222121-1230203133301011-2203333301323330-0303122302302000-0001122220100130-0013203230003312-0222030020233111"></a>

<a id="canonical-2131033000011201-0231320300313313-2331001201032012-3020310123231323-3321112131003221-1230001122130023-0103333222010022-3101332120320031"></a>

#### `routes.route_destination.cors_policy.allow_origin` property

Type: `["list", "string"]`. Optional.

Specifies the origins that will be allowed to do CORS requests. An origin is allowed if either
allow\_origin or allow\_origin\_regex match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1111200100321101-0111023300010103-1003031313231313-3133312211302100-3030212200133132-3133112133011031-0131102031333022-0200200103201032"></a>

<a id="canonical-3200130113323111-3323133111210001-3001102130330011-1210222023200023-1312321101100331-3222312220030110-3033233211213333-2100322011122003"></a>

#### `routes.route_destination.cors_policy.allow_origin_regex` property

Type: `["list", "string"]`. Optional.

Specifies regular expression patterns that match allowed origins. An origin is allowed if either allow\_origin or
allow\_origin\_regex match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1120312123121223-1212222203002303-2121032100112233-0313100131003030-2330233101131213-2001200210232211-0032010010233212-0001220321323323"></a>

<a id="canonical-2212232101232322-1002002322213112-1102001011000130-0210113100333021-0223210133202020-3323003212231201-3000021112020033-0010001110212020"></a>

#### `routes.route_destination.cors_policy.disabled` property

Type: `"bool"`. Optional.

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

<a id="canonical-2223112113323322-3120132232203030-3213313322322010-0001213132102302-2123231113023320-3122300021111112-0222212020200001-3131333131202021"></a>

<a id="canonical-3013133011323330-0330020200120122-1210011011032303-0023323020002012-0311333121130111-0323101132130111-3101311232132303-2033230100223102"></a>

#### `routes.route_destination.cors_policy.expose_headers` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3331311321130311-0312233203312001-3013022312322103-1320132021201011-3231030110121121-2230133322202300-0121221033211130-1311310200220221"></a>

<a id="canonical-2023021131330301-1011122110210332-0233130210203221-0112213232323332-3321120100210030-1232003032131311-1302331113002113-0300221233000221"></a>

#### `routes.route_destination.cors_policy.maximum_age` property

Type: `"number"`. Optional.

Specifies the content for the access-control-max-age header in seconds. This indicates the maximum
number of seconds the results can be cached A value of -1 will disable caching. Maximum permitted
value is 86400 seconds (24 hours).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(-1, 86400),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1101321211132233-2202003002310130-3123211010001032-2023310312100303-3310321001200211-2001220013330302-2200133210032112-2213001112022232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.csrf_policy` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.csrf_policy

<a id="canonical-1201101113101320-2323301123110211-1200230130122233-3301301233302201-3021210021301132-1021301001313121-2100102101223223-2233302311020121"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "custom_domain_list"),
  validators.ConflictingObjectAttributes("all_load_balancer_domains",
    "disabled"),
  validators.ConflictingObjectAttributes("custom_domain_list",
    "disabled")}
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
  "x-ves-oneof-field-allowed_domains": "[\"all_load_balancer_domains\",\"custom_domain_list\",\"disabled\"]"
}
```

Terraform syntax:

```terraform
csrf_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300332213122333-1230331023221030-1310030032102001-2133231112203000-0120310030101133-0310223012201131-3122330013232313-1120012313113303"></a>

### Direct properties for `routes.route_destination.csrf_policy`

- [all_load_balancer_domains](resources--route--reference--group-003.md#canonical-2321023330303333-2110120212233333-2200212232312332-0033130231212202-2201030130300313-3200022220121331-0101003113313231-1300120000111222): complete subsection reference.

- [custom_domain_list](resources--route--reference--group-003.md#canonical-2201223100330332-2010011312100323-1230032121233000-2120323021102231-3132301222332013-1013300133301132-1033201333113101-0131233012200202): complete subsection reference.

- [disabled](resources--route--reference--group-003.md#canonical-3232121301320030-3212231113202022-2001011312211100-1022003113030123-3323222303202022-0012310202101201-1132023203313220-3331303130222011): complete subsection reference.

<a id="canonical-2321023330303333-2110120212233333-2200212232312332-0033130231212202-2201030130300313-3200022220121331-0101003113313231-1300120000111222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.csrf_policy.all_load_balancer_domains` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.csrf_policy](resources--route--reference--group-003.md#canonical-1101321211132233-2202003002310130-3123211010001032-2023310312100303-3310321001200211-2001220013330302-2200133210032112-2213001112022232)
- routes.route_destination.csrf_policy.all_load_balancer_domains

<a id="canonical-2111101213212300-0130030103322333-1332123223120002-3111103133313301-1123113333333313-1012130102030020-1113003312222022-1010012231312002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

Additional upstream details:

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

Terraform syntax:

```terraform
all_load_balancer_domains = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201223100330332-2010011312100323-1230032121233000-2120323021102231-3132301222332013-1013300133301132-1033201333113101-0131233012200202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.csrf_policy.custom_domain_list` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.csrf_policy](resources--route--reference--group-003.md#canonical-1101321211132233-2202003002310130-3123211010001032-2023310312100303-3310321001200211-2001220013330302-2200133210032112-2213001112022232)
- routes.route_destination.csrf_policy.custom_domain_list

<a id="canonical-2011200113112300-1022220223302031-1200333312000212-3333000112000233-0200131212300031-1033212130200112-2232322101200303-1010012132201200"></a>

Type: `"object"`. single nested block, Optional.

List of domain names used for Host header matching.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
custom_domain_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1200133300200012-2102220102300120-0221113010130222-3330103023012301-1311000210302031-0013330301221311-0003112230012011-3120200003123200"></a>

### Direct properties for `routes.route_destination.csrf_policy.custom_domain_list`

<a id="canonical-2113012313133103-1230112103030103-3031012103330221-0103200013312213-3202332010101333-0011232323032230-0333110031110211-2211010302131233"></a>

#### `routes.route_destination.csrf_policy.custom_domain_list.domains` property

Type: `["list", "string"]`. Optional.

A list of domain names that will be matched to loadbalancer. These domains are not used for SNI
match. Wildcard names are supported in the suffix or prefix form.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 32),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3232121301320030-3212231113202022-2001011312211100-1022003113030123-3323222303202022-0012310202101201-1132023203313220-3331303130222011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.csrf_policy.disabled` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.csrf_policy](resources--route--reference--group-003.md#canonical-1101321211132233-2202003002310130-3123211010001032-2023310312100303-3310321001200211-2001220013330302-2200133210032112-2213001112022232)
- routes.route_destination.csrf_policy.disabled

<a id="canonical-2130231331111223-1212202021021111-2310102021210323-2210221111033232-0101302130031312-1221302112222322-1333032313010221-0301310213331023"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033310032010003-2013013212023023-3331003021220120-3211312230110102-0332232123000011-2322023233233003-2010313220111220-3203133211033012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.destinations` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.destinations

<a id="canonical-1112021013231321-0113122210020133-0133302023010120-1210312112321010-0022033012131121-3201100010031021-0301121112201221-3033121132300102"></a>

Type: `"object"`. list nested block, Optional.

When requests have to distributed among multiple upstream clusters, multiple destinations are
configured, each having its own cluster and weight. Traffic is distributed among clusters based on
the weight configured.

Example: destinations: &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-1 weight: 20 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-2 weight: 30 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-3 weight: 50

This indicates that out of every 100 requests, 50 goes to cluster-3, 30 to cluster-2 and 20 to
cluster-1

When single destination is configured, weight is ignored. All the requests are sent to the cluster
specified in the destination.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("cluster")}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
destinations {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023020210123001-2022231212232033-1000301011033000-1110130313221010-2220131032313111-1323133101230212-0112323003000023-3311003311021330"></a>

### Direct properties for `routes.route_destination.destinations`

- [cluster](resources--route--reference--group-003.md#canonical-3310330202210210-1003322300311113-0020121010120302-1230323021002201-3123231121001210-3223121233110230-3220300031023011-3112302031023321): complete subsection reference.

- [endpoint_subsets](resources--route--reference--group-003.md#canonical-1331310213003320-3302000201123230-1023221202020130-3311202130331323-1213300021321332-2301122031222030-3213022122020222-2131200200220130): complete subsection reference.

<a id="canonical-0231211220302310-3022333121201331-1003013221103033-0002202231313023-0211010332033010-3022301211332031-1301122300300022-2033121310233002"></a>

<a id="canonical-3032210201031101-1213301132222230-1221223320233012-1100013121123313-1030211202200331-2033110320322202-2300330303021023-1123113102301200"></a>

#### `routes.route_destination.destinations.priority` property

Type: `"number"`. Optional.

Priority of this cluster, valid only with multiple destinations are configured. Value of 0 will make
the cluster as lowest priority upstream cluster Priority of 1 means highest priority and is
considered active. When active cluster is not available, lower priority clusters are made active as
per the increasing priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 32),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-2211122320202232-0233131203103100-0212332313122322-0132220101021033-0130223203322103-2020101001221102-0310002000023120-3303310333030233"></a>

<a id="canonical-1132233123020123-3302020102122321-3001033222301333-2100013031200021-1020023213303121-3133310132012033-1333003111120010-3203203231321133"></a>

#### `routes.route_destination.destinations.weight` property

Type: `"number"`. Optional.

When requests have to distributed among multiple upstream clusters, multiple destinations are
configured, each having its own cluster and weight. Traffic is distributed among clusters based on
the weight configured.

Example: destinations: &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-1 weight: 20 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-2 weight: 30 &#8203;- cluster: &#8203;- kind: F5 xc.vega.cfg.adc.cluster.object uid:
cluster-3 weight: 10

This indicates that out of every 60 requests, 10 goes to cluster-3, 30 to cluster-2 and 20 to
cluster-1

When single destination is configured, weight is ignored. All the requests are sent to the cluster
specified in the destination.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 65535),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3310330202210210-1003322300311113-0020121010120302-1230323021002201-3123231121001210-3223121233110230-3220300031023011-3112302031023321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.destinations.cluster` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.destinations](resources--route--reference--group-003.md#canonical-2033310032010003-2013013212023023-3331003021220120-3211312230110102-0332232123000011-2322023233233003-2010313220111220-3203133211033012)
- routes.route_destination.destinations.cluster

<a id="canonical-3112032321222223-1010021133011211-1113310223033131-3132130010120322-2003312221021003-0013003132130320-0033112213100023-1220313200230202"></a>

Type: `"object"`. list nested block, Optional.

Indicates the upstream cluster to which the request should be sent. If the cluster does not exist
ServiceUnavailable response will be sent.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312310303101000-1223102220223012-3232203333202210-1131300210202123-0102111331122112-2010031211232210-1110303012133133-3211321103201012"></a>

### Direct properties for `routes.route_destination.destinations.cluster`

<a id="canonical-2023211112200200-0100113331330221-2222103221003131-2211321032130321-1300230223123331-3230202121220200-3002221233122210-1131030121010120"></a>

#### `routes.route_destination.destinations.cluster.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3100010121222333-3010001321212223-0213030111102020-3023213200303320-1120123120212221-1211333332023103-1022102322102222-1101202031332300"></a>

<a id="canonical-2010221133201033-1210202010031022-0303223313322221-2111122032313232-0010230033113133-2200320033323321-0313122130001102-3320033202330120"></a>

#### `routes.route_destination.destinations.cluster.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1211111020320010-2311111221323230-2320033033230113-1221220110023003-2112013130310000-1221303201312100-2210203201031131-3333022323130130"></a>

<a id="canonical-0300133202110102-1221213120222002-0331303332312021-3111312112011202-0320330103030112-1100310331003023-2133200222100031-2121303103102310"></a>

#### `routes.route_destination.destinations.cluster.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
  }
}
```

<a id="canonical-2331211232213031-0302233020211102-1332301332000233-3021222302012223-2000022010032202-3210120323131330-2212112213022330-1122223011100310"></a>

<a id="canonical-1230202303120333-0001233301212023-0031333320112232-2133013301020021-0110030202113132-1321310110231012-3313203033103033-2300330231302100"></a>

#### `routes.route_destination.destinations.cluster.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3321123113222022-2101000101122233-3110020210233011-1002101313223231-1202103133203322-3030223230221110-2120330332110203-0011013032100200"></a>

<a id="canonical-0302200002122031-2120323202132013-1001202033311133-3220322022130112-1031012301330303-2301030313030202-1030322121101130-2221100102011222"></a>

#### `routes.route_destination.destinations.cluster.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1331310213003320-3302000201123230-1023221202020130-3311202130331323-1213300021321332-2301122031222030-3213022122020222-2131200200220130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.destinations.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.destinations](resources--route--reference--group-003.md#canonical-2033310032010003-2013013212023023-3331003021220120-3211312230110102-0332232123000011-2322023233233003-2010313220111220-3203133211033012)
- routes.route_destination.destinations.endpoint_subsets

<a id="canonical-3023111002013112-0331013020301003-3212323002011002-1112223230232101-3133001122203222-0100001203201100-0003120121303232-1223203221010123"></a>

Type: `"object"`. single nested block, Optional.

Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached
to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be
selected by the load balancer

Labels field of endpoint object's metadata is used for subset matching. For endpoints which are
discovered in K8s or Consul cluster, the label of the service is merged with endpoint's labels. In
case of Consul, the label is derived from the "Tag" field. For labels that are common between
configured endpoint and discovered service, labels from discovered service takes precedence.

List of key-value pairs that will be used as matching metadata. Only those endpoints of upstream
cluster which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
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

Terraform syntax:

```terraform
endpoint_subsets {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232131110032213-2031333201002023-1223031002323001-0020002002100232-1113320131001233-1033123003220021-0030130030111321-2323233220010332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.do_not_retract_cluster` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.do_not_retract_cluster

<a id="canonical-3133130213210122-1031303100213011-1111022303303330-0312012102221332-0001232112220322-0332103332122131-3133202220002222-1021130122001220"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
do_not_retract_cluster = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111031211001032-1233310121003212-2023210031222203-2221132210120000-3013102123000323-3111031112232231-1202110122322030-1221020103030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.endpoint_subsets` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.endpoint_subsets

<a id="canonical-2133132312333130-0300231121312103-1332221322003010-2111323330203031-2131323312000232-0032222122112032-2211313312103010-3211322011201233"></a>

Type: `"object"`. single nested block, Optional.

Upstream cluster may be configured to divide its endpoints into subsets based on metadata attached
to the endpoints. Routes may then specify the metadata that a endpoint must match in order to be
selected by the load balancer

Labels field of endpoint object's metadata is used for subset matching. For endpoint's which are
discovered in K8s or Consul cluster, the label of the service is merged with endpoint's labels. In
case of Consul, the label is derived from the "Tag" field. For labels that are common between
configured endpoint and discovered service, labels from discovered service takes precedence.

List of key-value pairs that will be used as matching metadata. Only those endpoints of upstream
cluster which match this metadata will be selected for load balancing.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "originalRules": {
      "ves.io.schema.rules.map.max_pairs": "16"
    }
  },
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

Terraform syntax:

```terraform
endpoint_subsets {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.hash_policy

<a id="canonical-1013123110222103-0013032232130130-0302013230222213-0022333002111232-2111033110301233-2212332330201011-0003023310311103-2313110000010202"></a>

Type: `"object"`. list nested block, Optional.

Specifies a list of hash policies to use for ring hash load balancing. Each hash policy is evaluated
individually and the combined result is used to route the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("cookie",
    "header_name"),
  validators.ConflictingListObjectAttributes("cookie",
    "source_ip"),
  validators.ConflictingListObjectAttributes("header_name",
    "source_ip")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

Terraform syntax:

```terraform
hash_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232023303300102-1201233312123223-2313331002133032-1110321212333012-2210121130111201-0132233330131221-1212103212220103-2033323211011313"></a>

### Direct properties for `routes.route_destination.hash_policy`

- [cookie](resources--route--reference--group-003.md#canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101): complete subsection reference.

<a id="canonical-3030332102111320-3020113232322212-3101233203311223-1010021321003200-3322032323321223-0033112123323221-0122321001230102-0220302031201033"></a>

<a id="canonical-0320233233001213-1030310231302031-0111010210302320-1331312132301001-3111122310232301-0202011301020121-1102100202211320-3330102330220110"></a>

#### `routes.route_destination.hash_policy.header_name` property

Type: `"string"`. Optional.

Exclusive with \[cookie source\_ip\] The name or key of the request header that will be used to
obtain the hash key.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3222200112323331-0132001111011033-1031020113001202-2030021003001013-1101222102312200-1302103300113332-0203211003123333-1310233331023132"></a>

<a id="canonical-0310311313101203-3213132002002232-2131023311013131-2300302113213202-3021030311031122-3023113200122022-3000011003013003-1133323202130203"></a>

#### `routes.route_destination.hash_policy.source_ip` property

Type: `"bool"`. Optional.

Exclusive with \[cookie header\_name\] Hash based on source IP address.

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

<a id="canonical-3000123202120111-1230023210113110-3023021023030021-2013210011302133-3220320302311300-1302311102233113-3121103003232213-2311122211000103"></a>

<a id="canonical-2011123303120120-3230221113230121-3200002210223303-1012303120001321-3220112301333110-0022321012030223-3011311322222101-1113130021112121"></a>

#### `routes.route_destination.hash_policy.terminal` property

Type: `"bool"`. Optional.

Terminal. Specify if its a terminal policy.

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

<a id="canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123)
- routes.route_destination.hash_policy.cookie

<a id="canonical-0121132100220221-0023323022210331-2200322300133131-3203100112223213-0020321121012322-0000303122301332-2113212000022320-2130001121312221"></a>

Type: `"object"`. single nested block, Optional.

Two types of cookie affinity:

&#8203;1. Passive. Takes a cookie that's present in the cookies header and hashes on its value.

&#8203;2. Generated. Generates and sets a cookie with an expiration (TTL) on the first request from
the client in its response to the client, based on the endpoint the request gets sent to. The client
then presents this on the next and all subsequent requests. The hash of this is sufficient to ensure
these requests GET sent to the same endpoint. The cookie is generated by hashing the source and
destination ports and addresses so that multiple independent HTTP2 streams on the same connection
will independently receive the same cookie, even if they arrive simultaneously.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("name"),
  validators.ConflictingObjectAttributes("add_httponly",
    "ignore_httponly"),
  validators.ConflictingObjectAttributes("add_secure",
    "ignore_secure"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_lax"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_none"),
  validators.ConflictingObjectAttributes("ignore_samesite",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_none"),
  validators.ConflictingObjectAttributes("samesite_lax",
    "samesite_strict"),
  validators.ConflictingObjectAttributes("samesite_none",
    "samesite_strict")}
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
  "x-ves-oneof-field-httponly": "[\"add_httponly\",\"ignore_httponly\"]",
  "x-ves-oneof-field-samesite": "[\"ignore_samesite\",\"samesite_lax\",\"samesite_none\",\"samesite_strict\"]",
  "x-ves-oneof-field-secure": "[\"add_secure\",\"ignore_secure\"]"
}
```

Terraform syntax:

```terraform
cookie {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223010103101203-2232332013311102-3012210200012102-2023103213311212-3132221020010212-0100232332321031-2112013213213102-2200313310230011"></a>

### Direct properties for `routes.route_destination.hash_policy.cookie`

- [add_httponly](resources--route--reference--group-003.md#canonical-2022301212111113-0000112220113010-2333130322023111-0300212113310101-0101030322320312-0003230102132000-0032201132333100-1130332323031301): complete subsection reference.

- [add_secure](resources--route--reference--group-003.md#canonical-2330230302211001-1001233312013101-2321121113122011-2222020013133302-3302321130203330-0220133023310330-3001002011213132-3020212031133031): complete subsection reference.

- [ignore_httponly](resources--route--reference--group-003.md#canonical-1213303321232121-3033113213032222-2330220032013121-1233231222133222-3112130111310203-3223203330222001-3212031020321102-1213030111301320): complete subsection reference.

- [ignore_samesite](resources--route--reference--group-003.md#canonical-3223131201231103-3112020020012100-1131110222001221-1122002001320332-3322212020002332-0232202331111333-2020301001111123-1001312110322223): complete subsection reference.

- [ignore_secure](resources--route--reference--group-003.md#canonical-0011130100332212-3300100231300302-2122123223100101-1303032001132310-3033010210113223-2331110222023032-3002032021032001-2322030130323230): complete subsection reference.

<a id="canonical-2302030300221023-2331030021033101-1312032003011013-0201112131021130-1021222032000212-2110110122322300-1312101322220013-3120112100120302"></a>

<a id="canonical-3322131232012102-0032322132123013-3111303131032013-3212233021203113-3033303110001200-3102031100330121-3023023223012200-0102333121112023"></a>

#### `routes.route_destination.hash_policy.cookie.name` property

Type: `"string"`. Optional.

The name of the cookie that will be used to obtain the hash key. If the cookie is not present and
TTL below is not set, no hash will be produced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2202212010101230-3113131123302322-0003213230130300-0133303200122310-0322012003221001-0311112111213110-3213033230103321-2023203333023210"></a>

<a id="canonical-2332312213330301-1321231000323111-2130301330212330-3011020232320203-3133122032222310-3311311021021023-1331322002211312-1231322022222222"></a>

#### `routes.route_destination.hash_policy.cookie.path` property

Type: `"string"`. Optional.

The name of the path for the cookie. If no path is specified here, no path will be set for the
cookie.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.8,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [samesite_lax](resources--route--reference--group-003.md#canonical-0330022013103032-3312123211210022-0220100001032323-1232302301033311-3322102231220223-0003033002123112-0130212200111323-0011313112122003): complete subsection reference.

- [samesite_none](resources--route--reference--group-003.md#canonical-2201011232132012-3110112110233133-3301231010203102-0303320332030033-2220301300312022-0111110011313111-2201101121202313-3133120221031201): complete subsection reference.

- [samesite_strict](resources--route--reference--group-003.md#canonical-2020131323033213-0301333210023221-1002031200003031-2001230130111230-2233132233221110-0213230023200021-1320332023011013-0312100101121001): complete subsection reference.

<a id="canonical-0333302131313121-3003220321013302-0302023310313033-3301133120302022-0013212000332001-2002310323011023-1230332031213232-1113223213333230"></a>

<a id="canonical-3101203132212111-1231123103010100-3321003021002312-1001332303001301-2200133121122321-0320330013013012-2320133222212301-0101201023331003"></a>

#### `routes.route_destination.hash_policy.cookie.ttl` property

Type: `"number"`. Optional.

If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is
present and zero, the generated cookie will be a session cookie. TTL value is in milliseconds.

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

<a id="canonical-2022301212111113-0000112220113010-2333130322023111-0300212113310101-0101030322320312-0003230102132000-0032201132333100-1130332323031301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.add_httponly` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123)
- [routes.route_destination.hash_policy.cookie](resources--route--reference--group-003.md#canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101)
- routes.route_destination.hash_policy.cookie.add_httponly

<a id="canonical-3212231201012130-1013300112120131-2112100201132112-1100303313013020-1202202102233233-0310023302032233-1031101300032220-2030121030223001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for add httponly.

Additional upstream details:

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

Terraform syntax:

```terraform
add_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330230302211001-1001233312013101-2321121113122011-2222020013133302-3302321130203330-0220133023310330-3001002011213132-3020212031133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.add_secure` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123)
- [routes.route_destination.hash_policy.cookie](resources--route--reference--group-003.md#canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101)
- routes.route_destination.hash_policy.cookie.add_secure

<a id="canonical-2312212130233131-3221331103033330-3023122000300132-2012111202331322-0223120102001003-0112331302010013-1112201011003033-0301333123322300"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
add_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213303321232121-3033113213032222-2330220032013121-1233231222133222-3112130111310203-3223203330222001-3212031020321102-1213030111301320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.ignore_httponly` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123)
- [routes.route_destination.hash_policy.cookie](resources--route--reference--group-003.md#canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101)
- routes.route_destination.hash_policy.cookie.ignore_httponly

<a id="canonical-1313012031333203-1202020032331201-0301033201322333-3231001301303313-3001011323201200-0011100313023221-3303322101101012-0320321113230203"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore httponly.

Additional upstream details:

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

Terraform syntax:

```terraform
ignore_httponly = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3223131201231103-3112020020012100-1131110222001221-1122002001320332-3322212020002332-0232202331111333-2020301001111123-1001312110322223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.ignore_samesite` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123)
- [routes.route_destination.hash_policy.cookie](resources--route--reference--group-003.md#canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101)
- routes.route_destination.hash_policy.cookie.ignore_samesite

<a id="canonical-1012302222333022-1131133020203313-2232202232122013-1000020011323330-1112331233330311-0013231230132112-3101322231300311-2003101212321101"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
ignore_samesite = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011130100332212-3300100231300302-2122123223100101-1303032001132310-3033010210113223-2331110222023032-3002032021032001-2322030130323230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.ignore_secure` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123)
- [routes.route_destination.hash_policy.cookie](resources--route--reference--group-003.md#canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101)
- routes.route_destination.hash_policy.cookie.ignore_secure

<a id="canonical-0000320013100001-3302102311111101-2111213111133302-2303200023233131-0230001001211200-2100030310112220-3311010000220332-2030201312100113"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
ignore_secure = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330022013103032-3312123211210022-0220100001032323-1232302301033311-3322102231220223-0003033002123112-0130212200111323-0011313112122003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.samesite_lax` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123)
- [routes.route_destination.hash_policy.cookie](resources--route--reference--group-003.md#canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101)
- routes.route_destination.hash_policy.cookie.samesite_lax

<a id="canonical-2220033133323022-3322202332100320-2113110323303100-3301321122303303-3321212030133232-2102332122321132-1021131012021213-3302332002123220"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
samesite_lax = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2201011232132012-3110112110233133-3301231010203102-0303320332030033-2220301300312022-0111110011313111-2201101121202313-3133120221031201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.samesite_none` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123)
- [routes.route_destination.hash_policy.cookie](resources--route--reference--group-003.md#canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101)
- routes.route_destination.hash_policy.cookie.samesite_none

<a id="canonical-3132323303320010-2000331232023100-1232330132301001-0330031100201322-3310230111121302-0003330111320100-1221301031020112-0313031212230021"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
samesite_none = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020131323033213-0301333210023221-1002031200003031-2001230130111230-2233132233221110-0213230023200021-1320332023011013-0312100101121001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.hash_policy.cookie.samesite_strict` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.hash_policy](resources--route--reference--group-003.md#canonical-3011213001211130-3002130231011332-0021321333220210-3110310201022102-0211011003113300-2330021322000203-3323130012211301-3113313111302123)
- [routes.route_destination.hash_policy.cookie](resources--route--reference--group-003.md#canonical-1312331222131301-3233320013022212-0031130213010123-2110231112323301-3020021201212310-0132000123122100-0133210330202001-0133203210110101)
- routes.route_destination.hash_policy.cookie.samesite_strict

<a id="canonical-2001031230200211-1300113322313300-0111300330102032-0022121202310311-3003200013310030-2100313223231131-2230222032022010-1322002120123202"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
samesite_strict = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312110103201222-3211323013120213-1220232003210122-1333011320200321-0302110311030231-3120313321230312-0300310313300210-1231213001201021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.mirror_policy` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.mirror_policy

<a id="canonical-3103103220332022-3301230001230003-3012303123110101-2223301012122202-2211031023332111-2001013323223200-3231310322021311-1122013113121022"></a>

Type: `"object"`. single nested block, Optional.

MirrorPolicy is used for shadowing traffic from one cluster to another. The approach used is 'fire
and forget', meaning it will not wait for the shadow cluster to respond before returning the
response from the primary cluster. All normal statistics are collected for the shadow cluster
making..

Additional upstream details:

The approach used is "fire and forget", meaning it will not wait for the shadow cluster to respond
before returning the response from the primary cluster. All normal statistics are collected for the
shadow cluster making this feature useful for testing and troubleshooting.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cluster")}
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
mirror_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3133221133130223-3121323031321200-2001303031221130-2130032130320203-1300022311203221-3220223300320000-2001020222201233-3110012113031222"></a>

### Direct properties for `routes.route_destination.mirror_policy`

- [cluster](resources--route--reference--group-003.md#canonical-1220211122332310-1233121021033213-1300133022222333-3300333321020331-1131130320311031-3333131332101303-1223130101222032-3301112330021022): complete subsection reference.

- [percent](resources--route--reference--group-003.md#canonical-3110023323023222-0101122211022033-2121203321103121-3100133012031000-2322233300021221-3200131321221020-3221323010022221-0311103301212101): complete subsection reference.

<a id="canonical-1220211122332310-1233121021033213-1300133022222333-3300333321020331-1131130320311031-3333131332101303-1223130101222032-3301112330021022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.mirror_policy.cluster` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.mirror_policy](resources--route--reference--group-003.md#canonical-1312110103201222-3211323013120213-1220232003210122-1333011320200321-0302110311030231-3120313321230312-0300310313300210-1231213001201021)
- routes.route_destination.mirror_policy.cluster

<a id="canonical-1030111212111010-3323120011231231-2133313023011112-1002310313231321-2221033103313322-1313011222202320-3133221322313321-1222101001303022"></a>

Type: `"object"`. list nested block, Optional.

Specifies the cluster to which the requests will be mirrored. The cluster object referred here must
be present.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102323213003320-3022120333320112-0011013333220330-3102220331022012-2220013332202122-3103030010232322-3021233112010212-2231030122200103"></a>

### Direct properties for `routes.route_destination.mirror_policy.cluster`

<a id="canonical-2221031211000201-1220033013223111-2302031011231120-3032132022122322-0232322232131131-0132132031332103-2301332122020330-3220112132002003"></a>

#### `routes.route_destination.mirror_policy.cluster.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0203122202231013-3133012020002100-1311330321123202-3323010023223232-2133333123032011-0203313200303213-1123203013220030-1112311010300323"></a>

<a id="canonical-1111012121100311-1121233032321012-3203313120232001-0303332102112101-0021202221012302-0302001200130231-0002121323203101-1300002203312110"></a>

#### `routes.route_destination.mirror_policy.cluster.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1021310020012002-2010233013322010-3203103112022112-0330013323120320-3132101211101130-2220110302201021-1120013110000313-0000230203122121"></a>

<a id="canonical-1333022011303110-3101120301033213-0221301200320021-1310332033310312-0201002101023021-2013333130013112-1100012033110201-0133121021300300"></a>

#### `routes.route_destination.mirror_policy.cluster.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
  }
}
```

<a id="canonical-3213321013010012-0310230311312222-0303020231223313-2130300121232001-0313232023213031-2010132203100123-3130300003033200-0322103220333201"></a>

<a id="canonical-1222012311331333-2121220321102302-3120200223221202-2220230111310023-0103122021031231-0233023231102322-1010021121103101-1020100321030211"></a>

#### `routes.route_destination.mirror_policy.cluster.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0230013212020031-1321223123332333-1031033101222221-2110030220102210-2001320220133102-1220210013100113-0201122032321313-1101130313122111"></a>

<a id="canonical-2233000313130132-3003332221133303-0332000310130201-0000210332332320-2302120132212121-2112010031033203-3123313033321032-1002000111230033"></a>

#### `routes.route_destination.mirror_policy.cluster.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3110023323023222-0101122211022033-2121203321103121-3100133012031000-2322233300021221-3200131321221020-3221323010022221-0311103301212101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.mirror_policy.percent` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.mirror_policy](resources--route--reference--group-003.md#canonical-1312110103201222-3211323013120213-1220232003210122-1333011320200321-0302110311030231-3120313321230312-0300310313300210-1231213001201021)
- routes.route_destination.mirror_policy.percent

<a id="canonical-1221011213212311-0111213320333033-0200230201121021-1331233132231010-2103301020223121-0131121322020310-3113323201321122-0122000203332030"></a>

Type: `"object"`. single nested block, Optional.

Fraction used where sampling percentages are needed. Example sampled requests.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("numerator")}
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
percent {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213101003103102-1000333221331232-0021321201122011-0232033023223032-0110330010222100-0203113131221010-2322001132332003-3302113212022230"></a>

### Direct properties for `routes.route_destination.mirror_policy.percent`

<a id="canonical-0322132122101310-3332333202133030-0000301300122213-3322300211233313-3302232213002003-0313301301320320-1320313010033320-0010221133032321"></a>

#### `routes.route_destination.mirror_policy.percent.denominator` property

Type: `"string"`. Optional.

\[Enum: HUNDRED|TEN\_THOUSAND|MILLION\] Denominator used in fraction where sampling percentages are
needed. Example sampled requests Use hundred as denominator Use ten thousand as denominator Use
million as denominator. Possible values are \`HUNDRED\`, \`TEN\_THOUSAND\`, \`MILLION\`. Defaults to
\`HUNDRED\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("HUNDRED",
    "TEN_THOUSAND",
    "MILLION"),
}
```

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

<a id="canonical-2120233110230332-0132030231313322-2002033300333200-3013302331202211-0133023223012133-0233032230132023-2133320030221022-1300010210110210"></a>

<a id="canonical-2101231013332213-0220221221222222-0202102012200032-1023003231221303-3111010002022120-0023122200032011-0121213131122000-1330212101333311"></a>

#### `routes.route_destination.mirror_policy.percent.numerator` property

Type: `"number"`. Optional.

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

<a id="canonical-0210013123232110-1101203121320003-0232332230203020-2203202001132332-1123010132101312-1300303103231023-2212021131213210-0023212010013031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.query_params` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.query_params

<a id="canonical-3322102330231331-3300311203201320-0000122101103111-0303021333331323-0110130202120013-1000122210000132-1311002111001011-2011230101312303"></a>

Type: `"object"`. single nested block, Optional.

Handling of incoming query parameters in simple route.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("remove_all_params",
    "replace_params"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "retain_all_params"),
  validators.ConflictingObjectAttributes("replace_params",
    "retain_all_params")}
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
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]"
}
```

Terraform syntax:

```terraform
query_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131002102313110-3110323112303030-3323112202320031-2000001103303231-0033113203332010-0203132200233322-0010102000211111-3121113133131111"></a>

### Direct properties for `routes.route_destination.query_params`

- [remove_all_params](resources--route--reference--group-003.md#canonical-1031023210011311-3130112303030201-2022303313203011-1022012023320123-0302220133013033-1123212010021232-3010332323301021-0120320020030113): complete subsection reference.

<a id="canonical-2310033101202202-0311102331302133-0032200113103022-1322210031100130-1031022113100101-0023233313122120-3011011001112213-0030102233031330"></a>

<a id="canonical-2013021233303013-2222010133121010-2133033110011210-1230122030210113-1130023230203102-3112231313202221-1123110322332112-3030022100023031"></a>

#### `routes.route_destination.query_params.replace_params` property

Type: `"string"`. Optional.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [retain_all_params](resources--route--reference--group-003.md#canonical-3132311113113020-3222302331123313-2313032002331323-3112332133001221-2101122330010301-3013331331223111-1231100321110132-0232222123320221): complete subsection reference.

<a id="canonical-1031023210011311-3130112303030201-2022303313203011-1022012023320123-0302220133013033-1123212010021232-3010332323301021-0120320020030113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.query_params.remove_all_params` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.query_params](resources--route--reference--group-003.md#canonical-0210013123232110-1101203121320003-0232332230203020-2203202001132332-1123010132101312-1300303103231023-2212021131213210-0023212010013031)
- routes.route_destination.query_params.remove_all_params

<a id="canonical-2322332332011010-0233233011332323-2123011002233030-0130232223102323-3322121130223020-0320133222311302-0101120212121121-0202330323113232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

Additional upstream details:

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

Terraform syntax:

```terraform
remove_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132311113113020-3222302331123313-2313032002331323-3112332133001221-2101122330010301-3013331331223111-1231100321110132-0232222123320221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.query_params.retain_all_params` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.query_params](resources--route--reference--group-003.md#canonical-0210013123232110-1101203121320003-0232332230203020-2203202001132332-1123010132101312-1300303103231023-2212021131213210-0023212010013031)
- routes.route_destination.query_params.retain_all_params

<a id="canonical-0311130221111002-3330300100223330-0211020210230110-0020221103200130-3123102032212302-0333311320201113-0033322201013121-2032330220201100"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

Additional upstream details:

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

Terraform syntax:

```terraform
retain_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002132200332031-2112010113122321-3210320213210322-3203102103022111-0020102231301312-3020023031212111-1313020132121103-1031331331113022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.regex_rewrite` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.regex_rewrite

<a id="canonical-1013121130310212-2012233002003333-0120310130032131-2002312220330201-0333322120201030-2113031003131210-0013120202131000-3232130213013232"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
regex_rewrite {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032312101311031-2121113132103213-3003023102202120-2012000113011003-3322201022213213-0313213130132001-3013001322300302-1231011221221220"></a>

### Direct properties for `routes.route_destination.regex_rewrite`

<a id="canonical-0311210110212320-0230203213210212-3310002103200221-3332232332322132-1203123301012110-0011212131023023-1223023120320113-0321331210203331"></a>

#### `routes.route_destination.regex_rewrite.pattern` property

Type: `"string"`. Optional.

The regular expression used to find portions of a string that should be replaced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3132122230211331-0031333022032033-2113100322230131-3210001100100032-2220222021132230-2100332022103211-1123302122113223-1011103302200213"></a>

<a id="canonical-2222211310120110-1201001033201201-2321111121302002-3132201120221312-2220223210203101-2303202232203311-1113323311102021-1000211032212201"></a>

#### `routes.route_destination.regex_rewrite.substitution` property

Type: `"string"`. Optional.

The string that should be substituted into matching portions of the subject string during a
substitution operation to produce a new string.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1302302032030012-1303121333102133-0133131022330201-0212312130311012-1211001203310220-1221201023122013-2133021022130012-0333323020112031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.retract_cluster` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.retract_cluster

<a id="canonical-3133033201021013-1220100131310331-1023103000323112-0322203333003220-0111303312211200-2012200320231002-2331332302211033-1110213301231210"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
retract_cluster = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301103202310121-1322001133231203-3330201331230333-2121120301200322-3222320330230320-1021310030201212-1120112023110002-2203110113210212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.retry_policy` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.retry_policy

<a id="canonical-0332120203322330-0231001201132303-3131103033201212-1120133013223103-1202022020111333-0121203233322211-1110030301201322-0303332300101330"></a>

Type: `"object"`. single nested block, Optional.

Retry policy configuration for route destination.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("retry_condition")}
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
retry_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023013321013323-1000301101331210-2203203121011310-3322312312302001-3211332301323203-1023033213222013-0132021130002233-0223033100312130"></a>

### Direct properties for `routes.route_destination.retry_policy`

- [back_off](resources--route--reference--group-003.md#canonical-3333100010113321-1102030132022101-1120023010023101-1211232223320201-0213303320303033-1213213133220120-2100010200020111-3302303223023103): complete subsection reference.

<a id="canonical-3330313310300130-3312310312010001-3130022311003221-0230220102233132-2323201321202100-2210223100300010-2230203102012202-3233211212301120"></a>

<a id="canonical-1322101111300232-0120110020032213-0021210023012312-3311303230102323-1111312021033013-0231213202133220-1211102030022012-1332300003203201"></a>

#### `routes.route_destination.retry_policy.num_retries` property

Type: `"number"`. Optional.

Specifies the allowed number of retries. Retries can be done any number of times. An exponential
back-off algorithm is used between each retry. Defaults to \`1\`.

Additional upstream details:

Defaults to 1.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "8"
  }
}
```

<a id="canonical-1001121203001131-0211103020313001-0301332131232112-1131011033213310-1001332013303300-2031301013131130-0111110211103001-0133002232303101"></a>

<a id="canonical-2011231112200002-2321210033233302-1031102021003133-3320123233103201-2031032102133321-2300113113101223-3312200232121332-0111100010321131"></a>

#### `routes.route_destination.retry_policy.per_try_timeout` property

Type: `"number"`. Optional.

Specifies a non-zero timeout per retry attempt. In milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(600000),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2203021300100032-3203111312120103-1120203200023132-3032002101313003-1112120230223212-0333313011200302-0232121123231202-3333003103222300"></a>

<a id="canonical-1330230300030213-2122321300110020-1110231122312300-2333131201121002-3302211233101000-3020033032311202-0222023022100100-3122232023133313"></a>

#### `routes.route_destination.retry_policy.retriable_status_codes` property

Type: `["list", "number"]`. Optional.

HTTP status codes that should trigger a retry in addition to those specified by retry\_on.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0212201223233101-3232113001000331-1101210313232012-3213203123232000-2221212111313322-3322311020331033-1010013330220203-1330311101111131"></a>

<a id="canonical-2333001131011300-1330122002211300-1131113131303030-1001202111102223-3001220102221311-2033322213003133-2102011031011032-0220320330221233"></a>

#### `routes.route_destination.retry_policy.retry_condition` property

Type: `["list", "string"]`. Optional.

Specifies the conditions under which retry takes place. Retries can be on different types of
condition depending on application requirements. For example, network failure, all 5xx response
codes, idempotent 4xx response codes, etc The possible values are '5xx' : Retry will be done if
the..

Additional upstream details:

For example, network failure, all 5xx response codes, idempotent 4xx response codes, etc

The possible values are

"5xx" : Retry will be done if the upstream server responds with any 5xx response code, or does not
respond at all (disconnect/reset/read timeout).

"gateway-error" : Retry will be done only if the upstream server responds with 502, 503 or 504
responses (Included in 5xx)

"connect-failure" : Retry will be done if the request fails because of a connection failure to the
upstream server (connect timeout, etc.). (Included in 5xx)

"refused-stream" : Retry is done if the upstream server resets the stream with a REFUSED\_STREAM
error code (Included in 5xx)

"retriable-4xx" : Retry is done if the upstream server responds with a retriable 4xx response code.
The only response code in this category is HTTP CONFLICT (409)

"retriable-status-codes" : Retry is done if the upstream server responds with any response code
matching one defined in retriable\_status\_codes field

"reset" : Retry is done if the upstream server does not respond at all (disconnect/reset/read
timeout.)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 7),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 7,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 7,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"5xx\\\",\\\"gateway-error\\\",\\\"connect-failure\\\",\\\"refused-stream\\\",\\\"retriable-4xx\\\",\\\"retriable-status-codes\\\",\\\"reset\\\"]",
    "ves.io.schema.rules.repeated.max_items": "7",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3333100010113321-1102030132022101-1120023010023101-1211232223320201-0213303320303033-1213213133220120-2100010200020111-3302303223023103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.retry_policy.back_off` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- [routes.route_destination.retry_policy](resources--route--reference--group-003.md#canonical-0301103202310121-1322001133231203-3330201331230333-2121120301200322-3222320330230320-1021310030201212-1120112023110002-2203110113210212)
- routes.route_destination.retry_policy.back_off

<a id="canonical-2322231011021330-2102132021303133-3132030223101012-2210121213010213-0200032120022221-0031302302230023-3103312203222103-2020322003020233"></a>

Type: `"object"`. single nested block, Optional.

Specifies parameters that control retry back off.

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
back_off {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322023111013123-3102010330301033-2122123213103133-0010202320133330-0121022311000203-2002220012310211-0210210103202310-2131021032113021"></a>

### Direct properties for `routes.route_destination.retry_policy.back_off`

<a id="canonical-2213302023003231-3230112012222031-1101221232222011-1113223022002033-0201221202333130-3010212303332231-3311102330103111-2131202232330231"></a>

#### `routes.route_destination.retry_policy.back_off.base_interval` property

Type: `"number"`. Optional.

Specifies the base interval between retries in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gt": "0"
  }
}
```

<a id="canonical-0312303303212133-0310220220200023-0310123303300013-2323231000102133-2113013320110101-0313111321112222-1222121111103131-2111301021121001"></a>

<a id="canonical-1110103332122323-2330001032332233-1332020031321013-2323023032223331-2320202131021202-0132110303310211-1322030201310302-3023123010203302"></a>

#### `routes.route_destination.retry_policy.back_off.max_interval` property

Type: `"number"`. Optional.

Specifies the maximum interval between retries in milliseconds. This parameter is optional, but must
be greater than or equal to the base\_interval if set. The times the base\_interval. Defaults to
\`10\`.

Additional upstream details:

The default is 10 times the base\_interval.

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

<a id="canonical-1211112100111122-3213121213133231-1231320323033110-2010301220033103-1021202000330100-0023023012000230-0022012210003120-3103323301010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.spdy_config` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.spdy_config

<a id="canonical-0030003212102300-1303002203222302-0012302330132121-3322003213202013-2302221310210203-3332130210330001-0312101133322231-0012213100003100"></a>

Type: `"object"`. single nested block, Optional.

Request headers of such upgrade looks like below 'connection', 'Upgrade' 'upgrade', 'SPDY/3.1'

Configuration to allow UPGRADE of connection to SPDY and any additional tuning With configuration to
allow SPDY upgrade, ADC will produce following response 'HTTP/1.1 101 Switching Protocols 'Upgrade':
'SPDY/3.1' 'Connection': 'Upgrade'

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
spdy_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221223123303112-1132311121003011-1101113331001232-1133021113320022-3013323332321333-2303133030221000-1312331032102213-2010100302303203"></a>

### Direct properties for `routes.route_destination.spdy_config`

<a id="canonical-3123311300030233-2123110021122301-2012210221223130-3012121032221111-3330210021300003-0021203222001111-3100001231322013-2303320203332323"></a>

#### `routes.route_destination.spdy_config.use_spdy` property

Type: `"bool"`. Optional.

Specifies that the HTTP client connection to this route is allowed to upgrade to a SPDY connection.

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

<a id="canonical-0223313132331320-3312320001233000-1122103111302032-2311223023000012-0002221013031321-2022311222102223-3033032231010300-1312132322100022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_destination.web_socket_config` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_destination](resources--route--reference--group-003.md#canonical-0220130131223320-0012002013321000-2333211313321101-2211101322133023-2310223133333233-3200210230002323-0333010203210122-1120013002111310)
- routes.route_destination.web_socket_config

<a id="canonical-3101002232312011-1211310201000221-2323130121302321-0210310032312112-0020010111312100-1232122231102203-2230232321313322-3312113132221333"></a>

Type: `"object"`. single nested block, Optional.

Configuration to allow Websocket Request headers of such upgrade looks like below 'connection',
'Upgrade' 'upgrade', 'websocket' With configuration to allow websocket upgrade, ADC will produce
following response 'HTTP/1.1 101 Switching Protocols 'Upgrade': 'websocket' 'Connection': 'Upgrade'.

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
web_socket_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132301023323220-2232221010202320-0301233333220221-2310113330120323-1312220222020232-2013211313300000-1023131323102303-3010222022331002"></a>

### Direct properties for `routes.route_destination.web_socket_config`

<a id="canonical-1113111201133330-1023102023302213-3112232113230133-2211301220013013-3222101020120300-3322023031001122-3203323200221220-0120103221113331"></a>

#### `routes.route_destination.web_socket_config.use_websocket` property

Type: `"bool"`. Optional.

Specifies that the HTTP client connection to this route is allowed to upgrade to a WebSocket
connection.

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

<a id="canonical-0033133232132220-1302111112032120-3000302211110332-3011333112200100-0032120221023020-1101332301231111-1032033302231223-1323030110311233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_direct_response` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.route_direct_response

<a id="canonical-2302300123131210-3213103312320230-3310001122100100-2321123102113011-1202212130102111-2103211033323100-0000112122332101-2001013102102011"></a>

Type: `"object"`. single nested block, Optional.

Send this direct response in case of route match action is direct response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_code")}
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
route_direct_response {
  # Configure direct properties listed below.
}
```

<a id="canonical-0233121221012001-2310013102110301-3322233123110220-0210300132103321-3331202112323011-1323322231310320-2202123001312100-1120221202212230"></a>

### Direct properties for `routes.route_direct_response`

<a id="canonical-0222302303020322-3312111332232133-3113311202113321-0301303122100303-2101201311310001-3223302220120231-3233222213233331-1011112302133331"></a>

#### `routes.route_direct_response.response_body_encoded` property

Type: `"string"`. Optional.

Response body to send. Currently supported URL schemes is string:/// for which message should be
encoded in base64 format. The message can be either plain text or HTML. E.g. "&lt;p&gt; Access
Denied &lt;/p&gt;". base64 encoded string URL for this is
string:///PHA+IEFjY2VzcyBEZW5pZWQgPC9wPg==.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1121110031003012-0213231211323312-2123332020010221-0310030111300302-1300120221030231-0121002300230111-2030030021333322-0311203130333202"></a>

<a id="canonical-0020023222301203-2231120311123210-1212000122131021-1002133003020230-0103000222312322-2021100013013101-3110210131101032-2230231120321232"></a>

#### `routes.route_direct_response.response_code` property

Type: `"number"`. Optional.

Response Code. Response code to send.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(100, 599),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3030203133200300-2232032232300021-2021302312232013-2012023332102223-2002132031320330-1223213213033033-3233302022003331-1303231221020021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_redirect` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.route_redirect

<a id="canonical-2133332303131212-2113132023333213-3232122202313313-0300313003133233-3011011131333210-0010102301013122-2011110201201133-2223120002333222"></a>

Type: `"object"`. single nested block, Optional.

Route redirect parameters when match action is redirect.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path_redirect",
    "prefix_rewrite"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "replace_params"),
  validators.ConflictingObjectAttributes("remove_all_params",
    "retain_all_params"),
  validators.ConflictingObjectAttributes("replace_params",
    "retain_all_params")}
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
  "x-ves-oneof-field-query_params": "[\"remove_all_params\",\"replace_params\",\"retain_all_params\"]",
  "x-ves-oneof-field-redirect_path_choice": "[\"path_redirect\",\"prefix_rewrite\"]"
}
```

Terraform syntax:

```terraform
route_redirect {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220002112122201-0031332233103023-1221320331313213-1023210213223020-2112101333123320-1032122011230321-2303220132002111-0332131201223310"></a>

### Direct properties for `routes.route_redirect`

<a id="canonical-0101010203031003-0230330123212111-0122102202320322-0030210232212021-1120120032332120-3030210302022322-3033101303213312-3031021030031331"></a>

#### `routes.route_redirect.host_redirect` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2321122203200303-2303333300122111-0011200122023233-1303101211302100-2302021111020300-3203322033120320-1100123302120223-2112200202330333"></a>

<a id="canonical-3123123103010122-3223003112220300-1020320001223022-3312221301113122-2212233022111133-1030131123002010-2302311032210010-0110211010200003"></a>

#### `routes.route_redirect.path_redirect` property

Type: `"string"`. Optional.

Exclusive with \[prefix\_rewrite\] swap path part of incoming URL in redirect URL.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2133231130102301-0030333130120332-1203322110022230-2122132321233330-3010330213011322-2303233302103310-3303323200123100-0121211201310310"></a>

<a id="canonical-0213222300302102-3200330113203122-2230032103201311-2211011322230130-3312112311000302-3300122113220220-1232031101011320-3311200232110211"></a>

#### `routes.route_redirect.prefix_rewrite` property

Type: `"string"`. Optional.

Exclusive with \[path\_redirect\] In Redirect response, the matched prefix (or path) should be
swapped with this value. This option allows redirect URLs be dynamically created based on the
request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2111212001231200-3011023330002222-3100002230103332-3303130312113302-1131013220011203-3030101113312303-2333201203000013-2031123132031023"></a>

<a id="canonical-2013311232013223-2222021232132002-2322322012000333-3321321102121222-3301323002033211-0001212301312112-3233312331310021-3321032022332101"></a>

#### `routes.route_redirect.proto_redirect` property

Type: `"string"`. Optional.

\[Enum: incoming-proto|http|https\] Swap protocol part of incoming URL in redirect URL The protocol
can be swapped with either HTTP or HTTPS When incoming-proto option is specified, swapping of
protocol is not done. Possible values are \`incoming-proto\`, \`http\`, \`https\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("incoming-proto",
    "http",
    "https"),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [remove_all_params](resources--route--reference--group-003.md#canonical-2112000102133131-3330233010323032-2330110232311113-1222013003331011-2113100211123113-0122010001101103-3012000133322130-0130312202323132): complete subsection reference.

<a id="canonical-1310012213020231-0021013011110001-0001123211220323-1210231031011330-0323002121101322-1120203232133303-0102213133121203-3011222130132001"></a>

<a id="canonical-0013203212011000-2231200331301203-0100212213022331-2221232302133313-0132212332123303-0003103100321313-0121100300312301-3221001021301001"></a>

#### `routes.route_redirect.replace_params` property

Type: `"string"`. Optional.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2130203220230202-0120301232033223-1203032010212222-1112301300010311-2012223123120031-3220330300000003-0303030302003330-0200322111102321"></a>

<a id="canonical-3333203330202322-1233131210131321-3233220120321311-1121131123232122-1100213133102011-3220211323222231-3020323120112330-3122100313303322"></a>

#### `routes.route_redirect.response_code` property

Type: `"number"`. Optional.

The HTTP status code to use in the redirect response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.AtMost(599),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [retain_all_params](resources--route--reference--group-003.md#canonical-0312223003131122-3221203023023111-1111203122233310-1232321112210333-1002123301033232-0233302300300123-0101113332310031-1011301202103223): complete subsection reference.

<a id="canonical-2112000102133131-3330233010323032-2330110232311113-1222013003331011-2113100211123113-0122010001101103-3012000133322130-0130312202323132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_redirect.remove_all_params` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_redirect](resources--route--reference--group-003.md#canonical-3030203133200300-2232032232300021-2021302312232013-2012023332102223-2002132031320330-1223213213033033-3233302022003331-1303231221020021)
- routes.route_redirect.remove_all_params

<a id="canonical-1130112110102132-0323112211230000-1232021131331331-3101130220131031-0111301330030320-0100103011021211-3213122000121013-1230222121201101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

Additional upstream details:

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

Terraform syntax:

```terraform
remove_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312223003131122-3221203023023111-1111203122233310-1232321112210333-1002123301033232-0233302300300123-0101113332310031-1011301202103223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.route_redirect.retain_all_params` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.route_redirect](resources--route--reference--group-003.md#canonical-3030203133200300-2232032232300021-2021302312232013-2012023332102223-2002132031320330-1223213213033033-3233302022003331-1303231221020021)
- routes.route_redirect.retain_all_params

<a id="canonical-3233111231320202-3103213010231200-1000333130330113-0120203100231100-0322122110033111-3021201220220321-3011313123021331-3033333032201200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

Additional upstream details:

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

Terraform syntax:

```terraform
retain_all_params = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223310311031311-1102132201322103-0030022030223013-3213023033100232-1022023332031231-1310033332022211-3110312333301203-2212302001232030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.service_policy` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.service_policy

<a id="canonical-2023323222222023-1122101000020101-0133022030032131-3110002210200313-0221223013131021-1331312110320202-3111210010012031-2010022313202003"></a>

Type: `"object"`. single nested block, Optional.

ServicePolicy configuration details at route level.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-service_policy_choice": "[\"disable\"]"
}
```

Terraform syntax:

```terraform
service_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332321232321313-1301131200303003-1002012310110033-0203002010203012-3211310213021030-3222333223210100-1211130021322230-0002122210102120"></a>

### Direct properties for `routes.service_policy`

<a id="canonical-0212233320321030-1201131112030001-3332030223133030-2022210001322213-0213221132002203-1020301233333320-2032231013032230-2120121200230223"></a>

#### `routes.service_policy.disable_spec` property

Type: `"bool"`. Optional.

Exclusive with \[\] disable service policy at route level, if it is configured at virtual-host
level.

<a id="canonical-3113313331301130-2022203212231123-3330200013112321-1303112232000301-0332001232221002-3313202122033003-1323230113203023-3011123020002001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_exclusion_policy` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.waf_exclusion_policy

<a id="canonical-2002013313001200-2211230001233100-3213002023010302-0003113113213113-1220101021221323-2310310033302131-2033311313033122-3303301021200331"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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
waf_exclusion_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313001200301230-2302032333211113-2303131022201123-2223222223103020-3310133202211020-3113110121213032-2233203203113333-3022202112303011"></a>

### Direct properties for `routes.waf_exclusion_policy`

<a id="canonical-0323333110321132-1321231002100031-0332113312313302-2121031002221313-1132233200120001-2311323000001222-0120011311300232-2131030120330001"></a>

#### `routes.waf_exclusion_policy.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 128),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3320223201311312-1230332200221132-2023211032030213-2311120001232102-1330201233101010-3312101333222000-0032101301312212-3231202330011003"></a>

<a id="canonical-3023321002330111-0232212132312212-3003013310101032-3030232032323133-0022233212031202-3310031201001000-0131231113032102-0102210123133121"></a>

#### `routes.waf_exclusion_policy.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0112303323221002-0112010132321100-2020213300131313-2313332203100003-1110211220313310-3011301330230102-0320120110211120-0220022210011311"></a>

<a id="canonical-3111003221131220-1310023313203103-2111331031313332-2303310021131000-2233123020132113-0311000101100023-0311301331112123-1133021120200203"></a>

#### `routes.waf_exclusion_policy.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- routes.waf_type

<a id="canonical-2001321113131313-1313201110111212-3122331232022222-2331202122321020-0221120200230002-3200330210012331-0322102302101033-3110033121012002"></a>

Type: `"object"`. single nested block, Optional.

WAF instance will be pointing to an app\_firewall object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("app_firewall",
    "disable_waf"),
  validators.ConflictingObjectAttributes("app_firewall",
    "inherit_waf"),
  validators.ConflictingObjectAttributes("disable_waf",
    "inherit_waf")}
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
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

Terraform syntax:

```terraform
waf_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-1211223001102212-3303201100103210-3330303002132210-1100203222113213-1221031130233202-2333321331122320-0102303122011233-2101033321321320"></a>

### Direct properties for `routes.waf_type`

- [app_firewall](resources--route--reference--group-003.md#canonical-0200132232312312-1323232312210200-0011013333320012-3210230132301211-2013022232102130-2010311100203312-1233212201021313-2302110120000122): complete subsection reference.

- [disable_waf](resources--route--reference--group-003.md#canonical-2312222012021023-0130011200022021-2111231323310220-0123320233110003-3320301100010120-3210220010033100-2100321203330010-0023002113330003): complete subsection reference.

- [inherit_waf](resources--route--reference--group-003.md#canonical-3311301022131220-1131221221133322-2320231022201202-2033021133023220-3000103220310303-2210232232001201-2121022131100322-2012102321223011): complete subsection reference.

<a id="canonical-0200132232312312-1323232312210200-0011013333320012-3210230132301211-2013022232102130-2010311100203312-1233212201021313-2302110120000122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type.app_firewall` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- routes.waf_type.app_firewall

<a id="canonical-2231021231220033-0233233122220013-1111123320222301-3022211031003332-0300100210121302-1331002312111300-3302022222231002-0021032131203131"></a>

Type: `"object"`. single nested block, Optional.

A list of references to the app\_firewall configuration objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("app_firewall")}
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
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-1333221203222030-0323031111233021-3000303232212103-3003211330133230-1300302220220101-1230321132302032-1031200111312102-1221210110133232"></a>

### Direct properties for `routes.waf_type.app_firewall`

- [app_firewall](resources--route--reference--group-003.md#canonical-2311111213330130-1003101232130330-3221320120310311-3233210300103003-3102313130023232-2302313001002221-0100030100023323-0310113023320310): complete subsection reference.

<a id="canonical-2311111213330130-1003101232130330-3221320120310311-3233210300103003-3102313130023232-2302313001002221-0100030100023323-0310113023320310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type.app_firewall.app_firewall` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- [routes.waf_type.app_firewall](resources--route--reference--group-003.md#canonical-0200132232312312-1323232312210200-0011013333320012-3210230132301211-2013022232102130-2010311100203312-1233212201021313-2302110120000122)
- routes.waf_type.app_firewall.app_firewall

<a id="canonical-2230210213222132-1012121003022303-0323000302220103-3321133300210232-3211231012303323-1310203122333131-3311113033023200-0213002021322231"></a>

Type: `"object"`. list nested block, Optional.

References to an Application Firewall configuration object.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011110123321200-1202220213311320-3021301001133110-3310132121312002-2000231313322001-3233003001012223-2011200113321223-0301311330011311"></a>

### Direct properties for `routes.waf_type.app_firewall.app_firewall`

<a id="canonical-0221200322002023-1111100102000210-0022012102232211-2131103030222311-2003300313323301-1231311300220312-0010332013321010-1212012202001003"></a>

#### `routes.waf_type.app_firewall.app_firewall.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2113332133021120-1220103233200303-2230331332001002-0023322013133132-1333132123222221-1021002333312120-2033020003322030-1333321132312220"></a>

<a id="canonical-3123021320203232-3110102011301310-3300302310323133-1213223012302022-0033212202320330-3120003303101331-1110032303221111-1302323113203331"></a>

#### `routes.waf_type.app_firewall.app_firewall.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-3023230311231311-0210103201300133-3023212300121331-2221021202312331-3302003200221011-1332131113221120-2022033100100023-0032333322132230"></a>

<a id="canonical-3323232321022001-2223131020102222-2112003121013212-3232222003110021-2210322111133203-0133231123322213-2010031030111113-3333322332313200"></a>

#### `routes.waf_type.app_firewall.app_firewall.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
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
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
  }
}
```

<a id="canonical-1002113013001210-3323003100130111-3313321303111123-2221132302330032-1323233312030212-3313311302122331-0112123023103011-1201311213223302"></a>

<a id="canonical-0002031022101030-0121030000331013-0321221220332023-1120330232223120-1132031123230321-2020110130133313-1012220121012011-0203013321031101"></a>

#### `routes.waf_type.app_firewall.app_firewall.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2121302121131022-3021111201022332-2320202231233313-0322330330202130-1202113003123103-0121310301331031-2210223120012030-2333122130310120"></a>

<a id="canonical-0111002201202321-3311123300300333-0122101231310013-0130032121312130-2210332120120211-3001121330020101-0200122233310121-0000132201103310"></a>

#### `routes.waf_type.app_firewall.app_firewall.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-2312222012021023-0130011200022021-2111231323310220-0123320233110003-3320301100010120-3210220010033100-2100321203330010-0023002113330003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type.disable_waf` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- routes.waf_type.disable_waf

<a id="canonical-3011003032123110-2333031310103322-1333213021123113-0113330312130021-1111112331131121-0333010000212002-3120112133321320-2201031223002331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

Additional upstream details:

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

Terraform syntax:

```terraform
disable_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311301022131220-1131221221133322-2320231022201202-2033021133023220-3000103220310303-2210232232001201-2121022131100322-2012102321223011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `routes.waf_type.inherit_waf` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- [routes](resources--route--reference--group-001.md#canonical-3123120021213101-3123232233023211-0220000122002101-2303300120130232-3310221113321032-1122210133233232-3320330200321212-2120210203132130)
- [routes.waf_type](resources--route--reference--group-003.md#canonical-1322122033023130-0213202101131021-1131300023012121-1013202313230211-0213001020021200-2003012230323221-0031201130003011-1032223101322301)
- routes.waf_type.inherit_waf

<a id="canonical-1231213121221303-0023310201031321-1202211302132031-0232203213331212-2012120322121332-0132103323132111-0233231003033032-2121200121120001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inherit waf.

Additional upstream details:

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

Terraform syntax:

```terraform
inherit_waf = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000122211223303-0323203133120331-3213333010022223-1120331312123010-0233002030003331-1012023100032301-3200122303322132-3110023202032230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_route](../resources/route.md#canonical-2032020033333212-1122013011031221-2111330322212302-0032132111212121-3023032322101232-1231332022020023-2021133001000111-2130020232210312)
- [Property reference](resources--route--reference--group-001.md#canonical-3222103030211311-3130310001310032-2022203231320131-3102313133133331-2221123101031120-0130311122011322-0010220330230103-1030212032020232)
- timeouts

<a id="canonical-3111321022202013-2032120201220322-1301231002330300-0031231130003330-2001303332311331-2230212213301331-2030333231301120-3130212320122332"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100233003010002-0030312120330200-3312030232010311-3323210211132213-3221203012303110-0322031322111300-0302020113103331-1122100103321210"></a>

### Direct properties for `timeouts`

<a id="canonical-1130001230123010-1221321323133311-0332223011301032-1222031132033112-0312023111332000-1030220102100201-3331011330003112-1100012311023001"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0323312333201213-3130331323123302-0301112112123302-3300000201101000-2112201012020120-0302111212132300-1013131300210210-2222020332310212"></a>

<a id="canonical-1232022323332102-1002123232311003-1001112113123122-1101301203131013-3333233202020000-3312132132212310-2033021331002121-3130323133000112"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1101103120000211-1211330033202310-1212223212233002-1111001232202201-0121320303030010-1032322300320121-3022222303233011-3223030111021110"></a>

<a id="canonical-3130220211302101-2302100311331130-1231203023322111-3033122030321111-2133121332032320-0302102201200001-3032331113201321-1233332132330221"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0123103313202120-1031213332201320-2333113113000030-1103221112203330-0103121100203123-3132223020113212-0000203132130333-1332013103312230"></a>

<a id="canonical-0031033100220110-2000131103202232-0133333313013320-3130231032203020-2302113000012230-1123021223332202-2231230310120321-1323232313113210"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
