---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-1331200003331032-3122210122323003-3002303320230031-2001122322123310-3222010100102002-2322013132100110-0312303303221100-1232303203000032"></a>

## custom_page property — js_challenge / 203221113132 / 5

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
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

<a id="canonical-2022321002033033-1222211220022031-2322312231321001-3313100203332232-1212331321111223-0023212323333213-2021122210101220-0203001032332202"></a>

<a id="canonical-2021122302313233-2312123323222000-1123333210332002-1213333313323103-0021112233113011-2120003001322302-0101210210011212-1222023131033222"></a>

## js_script_delay property — js_challenge / 203221113132 / 6

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-0331310021123010-0031022021113010-1101123121100200-1233031120331001-0003112033310221-3022330233330111-1132212001332323-2212002121030133"></a>

## Next pages — js_challenge / 203221113132 / 7

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121121221032111-0001213320322323-3313031223203231-2330021213231230-0022330130330123-0103113031311201-3112101123001032-0131202003232111"></a>

## jwt_validation — jwt_validation / 200101012312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- jwt_validation

<a id="canonical-3010000120333003-3110233200033313-1221101022203110-0331331211223011-0323110201021110-2202203132122323-3111133202330220-0211121311033102"></a>

Type: `"object"`. single nested block, Optional.

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Upstream description:

JWT Validation stops JWT replay attacks and JWT tampering by cryptographically verifying incoming
JWTs before they are passed to your API origin. JWT Validation will also stop requests with expired
tokens or tokens that are not yet valid.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("authorization_server",
    "jwks_config")}
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
  "x-ves-oneof-field-jwks_configuration": "[\"authorization_server\",\"jwks_config\"]"
}
```

Terraform syntax:

```terraform
jwt_validation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230301001111321-1003222010221122-1110002332131121-2232322000112121-0311231113033331-0022113102333220-3102302331003113-2202332111133032"></a>

## Direct properties — jwt_validation / 200101012312 / 3

- [action](resources--cdn_loadbalancer--reference--group-012.md#canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022): complete subsection reference.

- [authorization_server](resources--cdn_loadbalancer--reference--group-012.md#canonical-0110223110031011-0102132203310210-0013213232003212-0300211202011312-3013233130322201-1203130022123313-1022310011112300-3113130303223121): complete subsection reference.

- [jwks_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-3021130213331113-3212213232102020-2003001113221112-1322313333022211-2202133110100233-2211202102121021-2330203102211021-1230100211231210): complete subsection reference.

- [mandatory_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-2313103132331201-1302213302203123-3300032330210320-3332032111310013-2320320003110320-2120230132023221-2312130313220303-1223233010233300): complete subsection reference.

- [reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131): complete subsection reference.

- [target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323): complete subsection reference.

- [token_location](resources--cdn_loadbalancer--reference--group-012.md#canonical-2302000320002133-3332213120223033-1013110220110333-0233320103332020-1321213311012223-0130112021222222-0023302032032102-1202103021203002): complete subsection reference.

<a id="canonical-3031302200331221-3220003213201011-0113222103110231-2101333001120023-1132102102233301-2021122112010112-3122211310001102-1231020131233331"></a>

## Next pages — jwt_validation / 200101012312 / 4

- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-012.md#canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022)
- [jwt_validation.authorization_server](resources--cdn_loadbalancer--reference--group-012.md#canonical-0110223110031011-0102132203310210-0013213232003212-0300211202011312-3013233130322201-1203130022123313-1022310011112300-3113130303223121)
- [jwt_validation.jwks_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-3021130213331113-3212213232102020-2003001113221112-1322313333022211-2202133110100233-2211202102121021-2330203102211021-1230100211231210)
- [jwt_validation.mandatory_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-2313103132331201-1302213302203123-3300032330210320-3332032111310013-2320320003110320-2120230132023221-2312130313220303-1223233010233300)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- [jwt_validation.token_location](resources--cdn_loadbalancer--reference--group-012.md#canonical-2302000320002133-3332213120223033-1013110220110333-0233320103332020-1321213311012223-0130112021222222-0023302032032102-1202103021203002)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221110232013323-1011301132310031-3103102023322231-0303202312101120-3102131233300110-3212132020132332-2221332223321211-1222332022002302"></a>

## jwt_validation.action — action / 303111202202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.action

<a id="canonical-1323120013011333-2220201103302333-3210232213210113-0121022002301320-0203213312201121-2103031230220110-3333032030002111-2201231233211130"></a>

Type: `"object"`. single nested block, Optional.

Action

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "report")}
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
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

Terraform syntax:

```terraform
action {
  # Configure direct properties listed below.
}
```

<a id="canonical-0221000121230113-1320013312022221-3311120131013330-0201300301121100-3323223203133001-3202110230201010-0011120220132333-2330003313331001"></a>

## Direct properties — action / 303111202202 / 3

- [block](resources--cdn_loadbalancer--reference--group-012.md#canonical-2213210033112131-3323111210023330-2000100231101110-2022222000323321-1101102311333112-2231001313211111-0211223113331110-1101132113220023): complete subsection reference.

- [report](resources--cdn_loadbalancer--reference--group-012.md#canonical-1201112212022020-2033223003230220-3301230021022333-3031213113002113-1313121302333233-3113203121112312-0120303112102212-2123233223223130): complete subsection reference.

<a id="canonical-1310020320300101-0323321200333220-3003201031001311-0103200321120031-3311221100010000-1010132212030100-0010313103022000-2003010113111231"></a>

## Next pages — action / 303111202202 / 4

- [jwt_validation.action.block](resources--cdn_loadbalancer--reference--group-012.md#canonical-2213210033112131-3323111210023330-2000100231101110-2022222000323321-1101102311333112-2231001313211111-0211223113331110-1101132113220023)
- [jwt_validation.action.report](resources--cdn_loadbalancer--reference--group-012.md#canonical-1201112212022020-2033223003230220-3301230021022333-3031213113002113-1313121302333233-3113203121112312-0120303112102212-2123233223223130)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2213210033112131-3323111210023330-2000100231101110-2022222000323321-1101102311333112-2231001313211111-0211223113331110-1101132113220023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002130031113221-2021131131103203-3331011113331123-3200320002002032-1130110233322220-3212303222032120-3101021002210123-3312121333100332"></a>

## jwt_validation.action.block — block / 012113323021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-012.md#canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022)
- jwt_validation.action.block

<a id="canonical-1300213012130111-2302002101000212-1231220231232232-1300023223132331-3011201002122010-1022200203332230-3223321210201311-1131330213121123"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
block = {}
```

<a id="canonical-0203230302320003-2222013231320220-0013232312122132-2302001223102333-2200202002000221-2002020031212130-0223002313023230-0210130332213213"></a>

## Direct properties — block / 012113323021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1313210121310013-1323221231223323-2201022200312013-2303331203311231-2322032032103330-2320020302201021-3313033301330030-1311110221200332"></a>

## Next pages — block / 012113323021 / 4

- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-012.md#canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1201112212022020-2033223003230220-3301230021022333-3031213113002113-1313121302333233-3113203121112312-0120303112102212-2123233223223130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0110110330123223-2313122010132110-2120221300002231-2003312213331002-0220121323123223-0131312212230102-1031030220110111-1203131331310220"></a>

## jwt_validation.action.report — report / 023110232320 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-012.md#canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022)
- jwt_validation.action.report

<a id="canonical-1103320213203001-0031310020200121-0331001003230312-1230013211231103-3313003113330333-0003013022011320-3331231303323322-2221332111211032"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
report = {}
```

<a id="canonical-3010111010022311-1222103101332213-2323110323303230-3001030030000323-3101300333223301-2123312203012122-3022213132011332-2222231320110221"></a>

## Direct properties — report / 023110232320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000121222333021-1110001220003121-3032301101323303-1111011120310113-0031021011130222-3021121330032132-2203222332201112-2003233312112021"></a>

## Next pages — report / 023110232320 / 4

- [jwt_validation.action](resources--cdn_loadbalancer--reference--group-012.md#canonical-0223230203222113-0121230030102011-3221021112332302-0301133130311010-2330332230233100-2310321110001000-1233133023210223-1021020212103022)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0110223110031011-0102132203310210-0013213232003212-0300211202011312-3013233130322201-1203130022123313-1022310011112300-3113130303223121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021300113033112-1221332020212222-1310021233023022-1223003023103303-1131303223102103-0213012000333210-0132332212000212-2220132232331121"></a>

## jwt_validation.authorization_server — authorization_server / 330203033323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.authorization_server

<a id="canonical-1301101223212301-3031321230202031-3010202033010311-0332103231001231-0030131122103310-2022002320010320-1213011120132312-3300113223033011"></a>

Type: `"object"`. single nested block, Optional.

Reference to Authorization Server object.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("authorization_servers")}
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
authorization_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102123023000303-1300112220102020-3100132212330313-1122022220202320-1012130213003121-2031111230031231-3112003220111111-2021313011232220"></a>

## Direct properties — authorization_server / 330203033323 / 3

- [authorization_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-2110120333003020-2233231003102223-0002322100233100-0302222022313331-2223332322031020-3233113122212222-0020231012121211-0202000100031201): complete subsection reference.

<a id="canonical-2123203023031131-3301303003333120-1220200010000122-2310010300100232-3111002202303331-3120321223131231-2022303102232202-3120320113313001"></a>

## Next pages — authorization_server / 330203033323 / 4

- [jwt_validation.authorization_server.authorization_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-2110120333003020-2233231003102223-0002322100233100-0302222022313331-2223332322031020-3233113122212222-0020231012121211-0202000100031201)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2110120333003020-2233231003102223-0002322100233100-0302222022313331-2223332322031020-3233113122212222-0020231012121211-0202000100031201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212131033101002-0030221301122032-1132001300311220-1330133100013130-3011233231211011-2113313022020132-3022201313000123-1311322233313203"></a>

## jwt_validation.authorization_server.authorization_servers — authorization_servers / 022301132213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.authorization_server](resources--cdn_loadbalancer--reference--group-012.md#canonical-0110223110031011-0102132203310210-0013213232003212-0300211202011312-3013233130322201-1203130022123313-1022310011112300-3113130303223121)
- jwt_validation.authorization_server.authorization_servers

<a id="canonical-2303130133022010-2220212010022333-2111131301133022-2001203031101013-3001220030013323-2123021020200031-1022232020020223-0111230332101312"></a>

Type: `"object"`. list nested block, Optional.

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Upstream description:

Authorization Servers are configured separately in the 'Shared Objects' section of the Web App &amp;
API Protection workspace and used to fetch JWKS for JWT validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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

Terraform syntax:

```terraform
authorization_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-3230332332000011-2021230010232013-3210120301031211-3202102113231120-1121312133202213-1011031300302200-1322022000302112-1113130200201023"></a>

## Direct properties — authorization_servers / 022301132213 / 3

<a id="canonical-2030332012113211-2011023201000222-0303022200231101-2010202020313312-2303203320220301-3320112103110313-2320002132133012-2310321301223022"></a>

<a id="canonical-2300233313310002-0301031031330113-0033222031323320-3123101322000222-1203320223012120-3220212223323212-3321200123303012-3320020010213201"></a>

## name property — authorization_servers / 022301132213 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-3333231222133331-0100223112121000-2110012013330130-1102230103300311-3100231012313103-1023233332310330-3031231302012232-2001013230113021"></a>

<a id="canonical-3013020211122222-1000332201012221-0301230022312301-1001023312100212-3230301103011112-2233331010003331-3312213311200102-3111132231031311"></a>

## namespace property — authorization_servers / 022301132213 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3302132101030320-3301312333111001-3022300232301112-0331103332121302-1011123023021100-0300210302332012-3202033022122000-0311032120213212"></a>

<a id="canonical-0103103230020001-2322121320333103-3332312322013021-2230233320011302-0330313101020332-0023302300312100-3310111330301110-1200231312123321"></a>

## tenant property — authorization_servers / 022301132213 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1231003010312113-3203120012011021-1120002020210331-3020202311232212-2203121320102211-2302323322102002-2323013320222213-3010102022120222"></a>

## Next pages — authorization_servers / 022301132213 / 7

- [jwt_validation.authorization_server](resources--cdn_loadbalancer--reference--group-012.md#canonical-0110223110031011-0102132203310210-0013213232003212-0300211202011312-3013233130322201-1203130022123313-1022310011112300-3113130303223121)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3021130213331113-3212213232102020-2003001113221112-1322313333022211-2202133110100233-2211202102121021-2330203102211021-1230100211231210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020330023010223-3320220223000311-3223010312232101-2110011133333330-2231303002133321-1220333220222220-0301303232000011-1310320023232212"></a>

## jwt_validation.jwks_config — jwks_config / 101130130100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.jwks_config

<a id="canonical-0220302331321313-1021303322302132-3111222311013312-0213023112300030-0210231211011210-1320130012203003-1110221030101232-0213233221022212"></a>

Type: `"object"`. single nested block, Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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
jwks_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300323121020310-3011302100331311-2201232012003033-3203200132302200-2120333110313202-3231001003100101-0102213201210101-2033032120223311"></a>

## Direct properties — jwks_config / 101130130100 / 3

<a id="canonical-1000021132311230-3122301123332220-0022221032222012-1003333102101301-0000103112133102-1101131000010103-1302212301330101-3010320302311322"></a>

<a id="canonical-0100111103121332-3321230103321211-2212323203332211-1132210101012123-3132230230000231-1033112200132322-1113213111011112-3002223030022333"></a>

## cleartext property — jwks_config / 101130130100 / 4

Type: `"string"`. Optional.

The JSON Web Key Set (JWKS) is a set of keys used to verify JSON Web Token (JWT) issued by the
Authorization Server. See RFC 7517 for more details.

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

<a id="canonical-1212203012103132-2312320101123001-2030300032221020-0123231122002321-2133312211031333-3300230213103323-3002302211121301-1113021031120120"></a>

## Next pages — jwks_config / 101130130100 / 5

- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2313103132331201-1302213302203123-3300032330210320-3332032111310013-2320320003110320-2120230132023221-2312130313220303-1223233010233300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333311322103310-2312200220121231-0020233033001231-1000132313103003-0313033330320132-3012122032131032-2221110032212031-1131221022132231"></a>

## jwt_validation.mandatory_claims — mandatory_claims / 222310213001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.mandatory_claims

<a id="canonical-2333203332013032-3133013113210012-2023130021310200-2212013021031222-1012021102002132-1232023332033131-1212232122221212-2120320101030133"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of mandatory Claims.

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
mandatory_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-1331232130011232-0023120320302322-1321013122103202-2102022230001203-2101311220020333-3022032310312302-1331023012112322-2121130232110202"></a>

## Direct properties — mandatory_claims / 222310213001 / 3

<a id="canonical-3102233101031012-2110001222130201-1010112222131103-1102112203002130-0112032200320113-2220211020122223-1330002222132203-1103313101220330"></a>

<a id="canonical-3000303202101301-2233300231211303-2002033331311330-1113001210231011-0312010012232232-1112000110231312-0010212132320220-0133033233010132"></a>

## claim_names property — mandatory_claims / 222310213001 / 4

Type: `["list", "string"]`. Optional.

Claim Names. Human-readable name for the resource

Upstream description:

Human-readable name for the resource

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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3201210010102211-3320100021233322-2103010332331103-0131323221010301-1110012212313210-3012000220113131-3333202312222212-1020202113200130"></a>

## Next pages — mandatory_claims / 222310213001 / 5

- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113122311320110-1300300323321122-2332313120222013-1203220031210000-2230022113101021-0032102301311101-0232100301303002-2011333003210212"></a>

## jwt_validation.reserved_claims — reserved_claims / 020000000110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.reserved_claims

<a id="canonical-3212112200320031-2330221301133131-1130100101222321-3330113120121203-1030103110202033-1030220122031003-1213310023233001-1313120322230210"></a>

Type: `"object"`. single nested block, Optional.

Configurable Validation of reserved Claims.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("audience",
    "audience_disable"),
  validators.ConflictingObjectAttributes("issuer",
    "issuer_disable"),
  validators.ConflictingObjectAttributes("validate_period_disable",
    "validate_period_enable")}
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
  "x-ves-oneof-field-audience_validation": "[\"audience\",\"audience_disable\"]",
  "x-ves-oneof-field-issuer_validation": "[\"issuer\",\"issuer_disable\"]",
  "x-ves-oneof-field-validate_period": "[\"validate_period_disable\",\"validate_period_enable\"]"
}
```

Terraform syntax:

```terraform
reserved_claims {
  # Configure direct properties listed below.
}
```

<a id="canonical-3330311220211020-3320000030023000-0333112121002001-0212222201300032-2313222112122303-3213002113222202-2013022312332323-2110303322233322"></a>

## Direct properties — reserved_claims / 020000000110 / 3

- [audience](resources--cdn_loadbalancer--reference--group-012.md#canonical-3020000102201001-1310310230021102-2212222030203033-3311112032232210-3321021200031101-3102332200012310-3313101332211331-1100330321303100): complete subsection reference.

- [audience_disable](resources--cdn_loadbalancer--reference--group-012.md#canonical-1213133021031020-1001310301233300-0201202323232203-0310130100002333-3311220123223122-1002211313133332-2203012101100123-2131133012121323): complete subsection reference.

<a id="canonical-2221103212000313-1333311201322302-1332221123230230-1122022312002103-0130300323111231-2102300300330032-2100313313132031-3300221201013303"></a>

<a id="canonical-3132332133321332-2212121030223231-0320033331031333-1311322210030133-2201203000020023-0331100320103201-1303212113113011-0013311322331330"></a>

## issuer property — reserved_claims / 020000000110 / 4

Type: `"string"`. Optional.

Exact Match. Exclusive with \[issuer\_disable\]

Upstream description:

Exclusive with \[issuer\_disable\]

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

- [issuer_disable](resources--cdn_loadbalancer--reference--group-012.md#canonical-0321023311132211-3113003212232022-0020212230220103-1121320022303333-2101300230003012-3011321131221023-0011313331332012-2132133331131231): complete subsection reference.

- [validate_period_disable](resources--cdn_loadbalancer--reference--group-012.md#canonical-1111232312303222-0032011222322301-1231332201023300-3300301202110211-1131213211130213-2012023033233123-1003121002331323-2310110002331131): complete subsection reference.

- [validate_period_enable](resources--cdn_loadbalancer--reference--group-012.md#canonical-0112011011200030-1002002021103021-0023122323131033-3331303133200000-0233212311122232-1130013133032100-0300110311310223-1123210303212102): complete subsection reference.

<a id="canonical-1303232121310020-0033122202202312-3022300031220032-1013031200221130-1003233200311130-3232200313201211-0020231330303101-1020023010131202"></a>

## Next pages — reserved_claims / 020000000110 / 5

- [jwt_validation.reserved_claims.audience](resources--cdn_loadbalancer--reference--group-012.md#canonical-3020000102201001-1310310230021102-2212222030203033-3311112032232210-3321021200031101-3102332200012310-3313101332211331-1100330321303100)
- [jwt_validation.reserved_claims.audience_disable](resources--cdn_loadbalancer--reference--group-012.md#canonical-1213133021031020-1001310301233300-0201202323232203-0310130100002333-3311220123223122-1002211313133332-2203012101100123-2131133012121323)
- [jwt_validation.reserved_claims.issuer_disable](resources--cdn_loadbalancer--reference--group-012.md#canonical-0321023311132211-3113003212232022-0020212230220103-1121320022303333-2101300230003012-3011321131221023-0011313331332012-2132133331131231)
- [jwt_validation.reserved_claims.validate_period_disable](resources--cdn_loadbalancer--reference--group-012.md#canonical-1111232312303222-0032011222322301-1231332201023300-3300301202110211-1131213211130213-2012023033233123-1003121002331323-2310110002331131)
- [jwt_validation.reserved_claims.validate_period_enable](resources--cdn_loadbalancer--reference--group-012.md#canonical-0112011011200030-1002002021103021-0023122323131033-3331303133200000-0233212311122232-1130013133032100-0300110311310223-1123210303212102)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3020000102201001-1310310230021102-2212222030203033-3311112032232210-3321021200031101-3102332200012310-3313101332211331-1100330321303100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1112121021333031-2132301303203301-1202120033031012-1310203002030313-0112332033330021-0111123000010102-2213323000010021-3321231201213111"></a>

## jwt_validation.reserved_claims.audience — audience / 301202102000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.audience

<a id="canonical-0300301132322231-2230203222131233-2222323000023201-1111010310313331-1030331102020023-1231302222302330-2210300203322031-1313321010011102"></a>

Type: `"object"`. single nested block, Optional.

Audiences

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("audiences")}
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
audience {
  # Configure direct properties listed below.
}
```

<a id="canonical-0333111132201332-0030300230201321-1031123330221331-1120110302132202-3303301030202331-3203110330023313-2133332213021110-1323322201301331"></a>

## Direct properties — audience / 301202102000 / 3

<a id="canonical-1200301323211332-0310021131020323-1332211203312102-1223321201322200-1132031232203131-0210031302131300-2131310013210220-0223311102232322"></a>

<a id="canonical-3000202203313311-1312022011221113-1032123300300230-2002202333103300-1011032211222002-3330010211201332-2022130023003010-3112010120322330"></a>

## audiences property — audience / 301202102000 / 4

Type: `["list", "string"]`. Optional.

Values. Configuration parameter for audiences

Upstream description:

Configuration parameter for audiences

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1200223121211300-3102311200220103-2010102320320010-0123103323223311-3003131020113330-1101213030320030-1002202123200320-3132231323203301"></a>

## Next pages — audience / 301202102000 / 5

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1213133021031020-1001310301233300-0201202323232203-0310130100002333-3311220123223122-1002211313133332-2203012101100123-2131133012121323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310311110333311-0013313110021131-3301003103233311-1111332310003112-0232320222212201-0020133022120321-2333020002331022-1100031232313012"></a>

## jwt_validation.reserved_claims.audience_disable — audience_disable / 133211231223 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.audience_disable

<a id="canonical-2013213021201033-1130101000020020-3211030031203202-3233301032000201-2211322313321133-0101322100331300-2030230032031100-3102120111023300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for audience disable.

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

Terraform syntax:

```terraform
audience_disable = {}
```

<a id="canonical-2022213230101233-1033003303302101-3103222333122320-3110331331212212-3103200222231131-1232102022332211-1003120213323122-0322123330212212"></a>

## Direct properties — audience_disable / 133211231223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301320302333013-0110322120113232-1222322320302211-2201131020312113-3123223331003221-2031113112313330-3331313003212312-3312003002031203"></a>

## Next pages — audience_disable / 133211231223 / 4

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0321023311132211-3113003212232022-0020212230220103-1121320022303333-2101300230003012-3011321131221023-0011313331332012-2132133331131231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021130301323002-1131200222020033-1031233232130111-2212021100220231-2332033011011000-0300311012032011-1002100011101233-2201320202310020"></a>

## jwt_validation.reserved_claims.issuer_disable — issuer_disable / 031213322211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.issuer_disable

<a id="canonical-1122300310321001-2110220222230333-2210201211331003-3222102321100031-2231211122000121-1310033231312332-1002223113212002-2233232202212103"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for issuer disable.

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

Terraform syntax:

```terraform
issuer_disable = {}
```

<a id="canonical-2102112312201220-1101303100010003-1131321130013021-1123213101202231-0130003023100313-3332223221313222-2030311111320033-2131322012220003"></a>

## Direct properties — issuer_disable / 031213322211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302112132213332-1032302001333103-0232121020122322-1123012321301031-1123321231122211-0201332012332210-3120012301310222-0110210021100223"></a>

## Next pages — issuer_disable / 031213322211 / 4

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1111232312303222-0032011222322301-1231332201023300-3300301202110211-1131213211130213-2012023033233123-1003121002331323-2310110002331131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132010323231220-1203230033001202-2012013302223122-2010221232302002-0333003102012032-0303133021331101-2330032310233333-0320130231123322"></a>

## jwt_validation.reserved_claims.validate_period_disable — validate_period_disable / 313313012223 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.validate_period_disable

<a id="canonical-0123313021103223-3321120103102331-3301213120202110-3210003033211110-3111003221121221-1313113332221121-1131211120330133-2112331131131211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period disable.

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

Terraform syntax:

```terraform
validate_period_disable = {}
```

<a id="canonical-1110020233230031-0312330032120213-2210211300322001-3310112003020010-0320331102203313-3302322300230232-3131112220313221-0002230221020023"></a>

## Direct properties — validate_period_disable / 313313012223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301130301312111-2031132000020100-3330323232022001-2333211012030221-1312320321133131-3012112322113002-0212310010111021-0101333310012132"></a>

## Next pages — validate_period_disable / 313313012223 / 4

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0112011011200030-1002002021103021-0023122323131033-3331303133200000-0233212311122232-1130013133032100-0300110311310223-1123210303212102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303312212330302-1330002121101131-3000023102031101-2212100122211022-2110303123020201-0332211103120000-3333131131001103-1233323332331221"></a>

## jwt_validation.reserved_claims.validate_period_enable — validate_period_enable / 202313212213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- jwt_validation.reserved_claims.validate_period_enable

<a id="canonical-3102112110201103-1131231020102030-2100032030213210-1301002231200123-2210120222323333-0211221121020310-0210330113112200-1320102332133300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for validate period enable.

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

Terraform syntax:

```terraform
validate_period_enable = {}
```

<a id="canonical-3332300311222330-1002211323311022-1210000120303131-2322313213230012-0310222223103111-3022302132212123-0230121003120032-2233232122010221"></a>

## Direct properties — validate_period_enable / 202313212213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000302313332020-1002021230003131-2230330103220113-0202020031200002-2023120201201332-1231002032320021-1000220010222222-0300333312213221"></a>

## Next pages — validate_period_enable / 202313212213 / 4

- [jwt_validation.reserved_claims](resources--cdn_loadbalancer--reference--group-012.md#canonical-0221130232003101-3201110231233233-3112310232323213-3110211100213232-3333010021223221-3311310010100133-3200233012300201-0312130210111131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022331003133211-1021333221233320-1203003010113113-0203332120233213-1020313211030131-0301013212122321-3120203021332203-3222102232020132"></a>

## jwt_validation.target — target / 130122210103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.target

<a id="canonical-2111203333133231-0102221210211213-1001311213213320-1331133320211012-3122321001013000-1001201133133203-2210230132302012-2330030300223013"></a>

Type: `"object"`. single nested block, Optional.

Define endpoints for which JWT token validation will be performed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("all_endpoint",
    "api_groups"),
  validators.ConflictingObjectAttributes("all_endpoint",
    "base_paths"),
  validators.ConflictingObjectAttributes("api_groups",
    "base_paths")}
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
  "x-ves-oneof-field-target": "[\"all_endpoint\",\"api_groups\",\"base_paths\"]"
}
```

Terraform syntax:

```terraform
target {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010132232132103-0302100221100013-0200303013133111-0321323211103233-0310031101010321-2303300210312321-0103300300331302-0322022031213130"></a>

## Direct properties — target / 130122210103 / 3

- [all_endpoint](resources--cdn_loadbalancer--reference--group-012.md#canonical-0030331312230121-3211212230231021-2323301133300032-3223131222122023-2130213301221001-0210031113111310-2300111333003032-3300212301030330): complete subsection reference.

- [api_groups](resources--cdn_loadbalancer--reference--group-012.md#canonical-3131230110321221-3002023032100200-0013122221223303-2133331321332211-1313130012021222-3001303132131030-1303330223100100-0121010000020200): complete subsection reference.

- [base_paths](resources--cdn_loadbalancer--reference--group-012.md#canonical-3202113120131021-0320333000010332-0032112121222220-2203030233320222-3021322010321222-3123302302310302-1310121322300300-2023103213131203): complete subsection reference.

<a id="canonical-1201112332110020-2113232121220323-2212033032112212-1333202023102302-2100021001113230-0312120030003002-3212302213103023-0233232311233020"></a>

## Next pages — target / 130122210103 / 4

- [jwt_validation.target.all_endpoint](resources--cdn_loadbalancer--reference--group-012.md#canonical-0030331312230121-3211212230231021-2323301133300032-3223131222122023-2130213301221001-0210031113111310-2300111333003032-3300212301030330)
- [jwt_validation.target.api_groups](resources--cdn_loadbalancer--reference--group-012.md#canonical-3131230110321221-3002023032100200-0013122221223303-2133331321332211-1313130012021222-3001303132131030-1303330223100100-0121010000020200)
- [jwt_validation.target.base_paths](resources--cdn_loadbalancer--reference--group-012.md#canonical-3202113120131021-0320333000010332-0032112121222220-2203030233320222-3021322010321222-3123302302310302-1310121322300300-2023103213131203)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0030331312230121-3211212230231021-2323301133300032-3223131222122023-2130213301221001-0210031113111310-2300111333003032-3300212301030330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3130203310310022-2101110221111101-1102030103332321-0113222022000031-1330321010132011-0311311330230232-3212011221122122-1300111033031130"></a>

## jwt_validation.target.all_endpoint — all_endpoint / 302203133311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- jwt_validation.target.all_endpoint

<a id="canonical-3002330221301310-1023102100110000-1221110323321031-1220202323033230-0012031322013320-3033131331232110-2112031330303302-0213302001003322"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
all_endpoint = {}
```

<a id="canonical-2322011310110232-3030201002233113-3210031133021131-1003211103132123-0310010003313020-2221102213010130-2213033121111230-1131332122330011"></a>

## Direct properties — all_endpoint / 302203133311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313301001023323-3003233031323302-2030230221032221-3112321212212103-1231031312211020-0301232011222320-3120321223022302-2113011001121203"></a>

## Next pages — all_endpoint / 302203133311 / 4

- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3131230110321221-3002023032100200-0013122221223303-2133331321332211-1313130012021222-3001303132131030-1303330223100100-0121010000020200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220302200022001-3010321033110122-1333332221032233-0231230330220001-0002012100002113-3032333020001221-0121111101303133-2101001312030320"></a>

## jwt_validation.target.api_groups — api_groups / 330321333332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- jwt_validation.target.api_groups

<a id="canonical-1003333003022021-0123320210213200-0231030011303132-2021100303201323-2320113103221000-1100320112303230-1201123000213201-0032123100223112"></a>

Type: `"object"`. single nested block, Optional.

API Groups.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_groups")}
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
api_groups {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132011001021213-1013032323300013-0202222300313110-0233121132110031-0333200301133031-3033332321011133-3313031211133023-3231033310132020"></a>

## Direct properties — api_groups / 330321333332 / 3

<a id="canonical-2011010113312321-1232232300221232-3011100312223123-1101010332330213-0123211201111331-3110133232112310-2011232032012321-0212232232331332"></a>

<a id="canonical-2003220013300030-0230011313101002-2202111003210322-0032111130130312-0320001012130232-3001300013310203-0301133223311123-0210030022320232"></a>

## api_groups property — api_groups / 330321333332 / 4

Type: `["list", "string"]`. Optional.

API Groups. Group or collection configuration

Upstream description:

Group or collection configuration

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3323302013202233-2131030321333110-2110130213232010-0123122102011113-0202101312301120-1222313321023213-1103000032310313-1120221303000101"></a>

## Next pages — api_groups / 330321333332 / 5

- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3202113120131021-0320333000010332-0032112121222220-2203030233320222-3021322010321222-3123302302310302-1310121322300300-2023103213131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012211122103301-0023033021332131-2320322120321302-3032010132310010-1331202212221201-0311020012030123-0000331122123131-2030333313322002"></a>

## jwt_validation.target.base_paths — base_paths / 312130120030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- jwt_validation.target.base_paths

<a id="canonical-1002012233312203-2010132302013220-3110220302223123-0102330101233123-3011311321231011-0012323032331023-1230101330313111-1002111130200332"></a>

Type: `"object"`. single nested block, Optional.

Base Paths.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("base_paths")}
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
base_paths {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010130331332121-3233001031312202-3121022211322202-2320111322233322-3221311120033213-2320001322300211-3331101101013023-3022210010012031"></a>

## Direct properties — base_paths / 312130120030 / 3

<a id="canonical-1221103003030123-1102223231312223-0130123232021011-3021231210223021-1103210202101322-1101113303013303-2112111121103331-1023100211223010"></a>

<a id="canonical-2123302320123301-3011221321022322-3221232023030311-1201301002201031-3103120222103110-3112121202020112-2102211231231202-0031220210123022"></a>

## base_paths property — base_paths / 312130120030 / 4

Type: `["list", "string"]`. Optional.

Prefix Values. File system or URL path

Upstream description:

File system or URL path

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.http_path": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2112010123022123-1133013211110020-0120203020213033-0023332303321032-1010332021130112-2011030031333013-1203312230210112-0012112030220111"></a>

## Next pages — base_paths / 312130120030 / 5

- [jwt_validation.target](resources--cdn_loadbalancer--reference--group-012.md#canonical-2120213333333321-0012122231331331-3302233122311210-1321300003200223-2030111332003023-1200123031001332-1033230121231132-0013332133311323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2302000320002133-3332213120223033-1013110220110333-0233320103332020-1321213311012223-0130112021222222-0023302032032102-1202103021203002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002113311110313-2230213023031013-0020312221122333-2200031220200310-0132332130301301-1033121230220210-2133000330001121-2212233103010301"></a>

## jwt_validation.token_location — token_location / 233230303231 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- jwt_validation.token_location

<a id="canonical-3222122003223131-3300113130232013-1021013133211032-2220222012232313-0233132133230130-2230113120303112-2210333302201301-1333331330321023"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for token location.

Upstream description:

Location of JWT in HTTP request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-token_location": "[\"bearer_token\"]"
}
```

Terraform syntax:

```terraform
token_location {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231213011230232-1013002010111032-0212203132000100-0211022332222302-0002233233110220-3121311223221232-0310302003033300-0312302030100020"></a>

## Direct properties — token_location / 233230303231 / 3

- [bearer_token](resources--cdn_loadbalancer--reference--group-012.md#canonical-2313130223203023-3212313023211320-2030123032320301-2200113333021001-3001003000201002-0020210321111101-1211213033311112-1113332103231330): complete subsection reference.

<a id="canonical-3131010130323230-2003002113213321-0323332132010231-1220231022310022-0200303030023220-0013220122100120-0301300221212322-3110200010113313"></a>

## Next pages — token_location / 233230303231 / 4

- [jwt_validation.token_location.bearer_token](resources--cdn_loadbalancer--reference--group-012.md#canonical-2313130223203023-3212313023211320-2030123032320301-2200113333021001-3001003000201002-0020210321111101-1211213033311112-1113332103231330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2313130223203023-3212313023211320-2030123032320301-2200113333021001-3001003000201002-0020210321111101-1211213033311112-1113332103231330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323231003112102-0011022103031222-0133222100231303-2000023300211311-3112333202202120-2132220031230220-3212201221020231-2200222323031032"></a>

## jwt_validation.token_location.bearer_token — bearer_token / 031231223301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [jwt_validation](resources--cdn_loadbalancer--reference--group-012.md#canonical-3231211001230011-2120300311230023-2120101132320213-0303030323233113-3131333213011222-2310202013201013-0200320010001222-3022100012232031)
- [jwt_validation.token_location](resources--cdn_loadbalancer--reference--group-012.md#canonical-2302000320002133-3332213120223033-1013110220110333-0233320103332020-1321213311012223-0130112021222222-0023302032032102-1202103021203002)
- jwt_validation.token_location.bearer_token

<a id="canonical-1122001322110203-3110301233333101-2203312111331323-3131233011211321-3010113321113202-3030232013002122-3120101210012133-1132333130101223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bearer token.

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

Terraform syntax:

```terraform
bearer_token {}
```

<a id="canonical-0121220330021113-0101233002122223-0022003231303333-1100233330031311-1030301032322112-1101020130202311-3003120032223212-0003100010300300"></a>

## Direct properties — bearer_token / 031231223301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030022320313130-0100013211133331-0010203130113330-2023321101023030-3132102210233023-2323331011323023-3120031120223020-0011322230221020"></a>

## Next pages — bearer_token / 031231223301 / 4

- [jwt_validation.token_location](resources--cdn_loadbalancer--reference--group-012.md#canonical-2302000320002133-3332213120223033-1013110220110333-0233320103332020-1321213311012223-0130112021222222-0023302032032102-1202103021203002)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1302023000110302-3120111302312101-1232211223311220-1031332223001000-2302333000033013-1001023131000303-3232033331020302-2320323022012121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233220300120323-1131031101102233-0331200210111232-2112100213100230-0211130111330323-3120221330032013-1221023000301220-3230232010100233"></a>

## l7_ddos_action_block — l7_ddos_action_block / 130000132001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- l7_ddos_action_block

<a id="canonical-0100201001123100-2323102022302103-2201020102221003-0300000001122200-2003030213312113-0010303131320223-2011133100020323-0102111133321110"></a>

Type: `["object", {}]`. Optional.

\[OneOf: l7\_ddos\_action\_block, l7\_ddos\_action\_default, l7\_ddos\_action\_js\_challenge;
Default: l7\_ddos\_action\_default\] Enable this option

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

OneOf alternatives in this subsection:

- [l7_ddos_action_block](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201001123100-2323102022302103-2201020102221003-0300000001122200-2003030213312113-0010303131320223-2011133100020323-0102111133321110)
- [l7_ddos_action_default](resources--cdn_loadbalancer--reference--group-012.md#canonical-3031030021303312-3321123103123321-1201032120110223-1301323023112110-1300131320013230-1110313213213133-2323001312232023-3233331200030031)
- [l7_ddos_action_js_challenge](resources--cdn_loadbalancer--reference--group-012.md#canonical-2023312331022300-1320031101102101-2232303210212002-3310102232320022-3213121210230333-1311223111301323-2221121220030321-2331213011201131)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
l7_ddos_action_block = {}
```

<a id="canonical-3120131011233111-0221131002203023-1302032003320102-3011220202132130-1331132131331123-1301203123221311-0203030120113122-3101010012211201"></a>

## Direct properties — l7_ddos_action_block / 130000132001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113033130133012-0022133022211202-2231122323322223-0113330012013323-1202310101120020-3022002023331310-0023312202103131-0120011100100030"></a>

## Next pages — l7_ddos_action_block / 130000132001 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2230330203013032-1133333031113100-0021112120003221-1230203101320213-0202130233200301-3000332131012213-3203322211203031-1131222233212002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300232002000120-3311322032002201-3130100031202033-3020120200003132-3021101313003310-2301011312112112-1133012303223113-3023310332312312"></a>

## l7_ddos_action_default — l7_ddos_action_default / 101121131321 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- l7_ddos_action_default

<a id="canonical-3031030021303312-3321123103123321-1201032120110223-1301323023112110-1300131320013230-1110313213213133-2323001312232023-3233331200030031"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
l7_ddos_action_default = {}
```

<a id="canonical-1030200033311323-3232101102000012-3221332322333110-3323220320310230-2112003002101330-2032131330010000-2221202013000110-3330223233200110"></a>

## Direct properties — l7_ddos_action_default / 101121131321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131210121110330-3011001233100322-0332222023002123-1122133132112210-1313203211322232-0301331121211212-2211202030221032-3322211202133003"></a>

## Next pages — l7_ddos_action_default / 101121131321 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3230011220011130-3203232003333030-0210211220332212-1300220223001301-2310021232213132-0112010313031231-3000331122302003-2133202101030211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300330313230010-2003311202121011-1111211300102330-2321021303110321-0133213010010230-1020101231332103-0202110033032203-2321000331230100"></a>

## l7_ddos_action_js_challenge — l7_ddos_action_js_challenge / 201113220303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- l7_ddos_action_js_challenge

<a id="canonical-2023312331022300-1320031101102101-2232303210212002-3310102232320022-3213121210230333-1311223111301323-2221121220030321-2331213011201131"></a>

Type: `"object"`. single nested block, Optional.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
l7_ddos_action_js_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-3001202123100011-3201021132003310-1233221322200121-3121312310312123-0030010211010021-3030023122312232-0023101231031010-3222000212131020"></a>

## Direct properties — l7_ddos_action_js_challenge / 201113220303 / 3

<a id="canonical-2301032201201232-3301200030001211-1322221010210101-3013220220013110-1132312132330130-1111100130002132-2003130003331012-1201133333132322"></a>

<a id="canonical-3323303011103300-2302003003133313-2102330101322100-2310101232122122-1103120321023012-0022120303102131-3113001130332323-1200201202310220"></a>

## cookie_expiry property — l7_ddos_action_js_challenge / 201113220303 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
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
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0313311313220003-1013310212321020-2203011303210033-0213110011012201-0030022202002003-0103312312110211-3130301323113230-3213020232311332"></a>

<a id="canonical-3130012100002312-3122111232311032-1013233331212021-3332113132310321-2131302110132312-0202203101001223-3330010301232003-3100033332030301"></a>

## custom_page property — l7_ddos_action_js_challenge / 201113220303 / 5

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
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

<a id="canonical-0202131023303313-2103100021130301-1323121013100033-1200212103313230-2322122220333013-3003120120321030-0100220303233022-1001203030222301"></a>

<a id="canonical-3031120012200012-3310222130222223-0101102313322211-1100020021013230-0213222330223322-3331123332123100-0101332001010311-3010300330220100"></a>

## js_script_delay property — l7_ddos_action_js_challenge / 201113220303 / 6

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-2200100031121021-3223021120010111-1102300012222021-0002303211231101-2333232233233203-1332210132022022-3122300011330120-2310330313200011"></a>

## Next pages — l7_ddos_action_js_challenge / 201113220303 / 7

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3312000011031203-2231300221302120-2103312210100030-2210320110231313-2210213211211221-2111110102013130-3032201012110003-2023110221100031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203313330022120-1110333130110002-0101113020130302-3130133331023310-2203012021210232-2112311211131133-3020230211323222-3003200222323231"></a>

## no_challenge — no_challenge / 202021203130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- no_challenge

<a id="canonical-3122322231130001-0110033223312111-0211313102301220-1122203023011032-2300330122000333-2313023000331020-1121211003020210-1230333230203311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no challenge.

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

Terraform syntax:

```terraform
no_challenge = {}
```

<a id="canonical-2333130331320002-2132112331103133-1132020212130000-3330330300013132-2012101011103322-3300203300210203-1330211130110020-2111022120013031"></a>

## Direct properties — no_challenge / 202021203130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133002021210102-2003211313220022-0230103210002113-1331103023310002-0213111022111220-2102112300032323-0122231121301203-1313130331033112"></a>

## Next pages — no_challenge / 202021203130 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3123112331011010-0132220300333311-0001300132330131-1132331012130223-3311232302020312-2200131102201331-2233330000312303-0132221122110323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1121022320313301-2201333111001221-1002112321201232-2100031130132103-1032333032302220-3010101313122133-0133220121102010-0233221103033311"></a>

## no_service_policies — no_service_policies / 312001003312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- no_service_policies

<a id="canonical-1013310300231013-2321212333203010-0332022031203123-3112002213301121-3010211320310130-3233202333011223-1131111100001030-2310100131120023"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no service policies.

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

Terraform syntax:

```terraform
no_service_policies = {}
```

<a id="canonical-3103020132003333-3331003210020212-0102323103202220-3232313322300310-3213021220020331-1310210132120123-0222322111320131-0331120002313302"></a>

## Direct properties — no_service_policies / 312001003312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030310011000102-2212010012032201-2101201333030211-1200332103333202-0120300201103011-1323301331302033-3200031230333121-1010101123313132"></a>

## Next pages — no_service_policies / 312001003312 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213100313032122-2332102200122112-1312110310120331-0121331312231011-1310313232103111-3310302132212331-1100033313033133-3311223111320003"></a>

## origin_pool — origin_pool / 330133032131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- origin_pool

<a id="canonical-1333023300232232-3020120010223111-2210032333323302-3202332123100233-2202001011030210-3120233023331030-0011213003323111-2130022000131010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin pool.

Upstream description:

Origin Pool for the CDN distribution.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("origin_servers"),
  validators.ConflictingObjectAttributes("no_tls",
    "use_tls")}
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
  "x-ves-oneof-field-tls_choice": "[\"no_tls\",\"use_tls\"]"
}
```

Terraform syntax:

```terraform
origin_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-2300011113110010-1033012302100221-0013230102033201-3311120231201100-0200011020113011-2301213220131331-1131010231331032-2030301212030103"></a>

## Direct properties — origin_pool / 330133032131 / 3

- [more_origin_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0213200101331231-2120013110121131-1210303023133011-1030133313113032-3211300011331111-3213130011023120-2021320133300332-2231122333022102): complete subsection reference.

- [no_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0031320223020112-3322011231321300-2133213030302010-3003303230121300-2012120113121303-3110323013020211-2002003232313303-3223222032032013): complete subsection reference.

<a id="canonical-3121320022332231-0330103022121111-3103320130101311-1321320232331332-2303131132013201-2201010220211332-2221332111321201-2200012313213033"></a>

<a id="canonical-1200130013021011-2312331310212211-3211131112222101-0311302231020230-2320020330333200-2122103001033332-3222102132112031-2210211101322210"></a>

## origin_request_timeout property — origin_pool / 330133032131 / 4

Type: `"string"`. Optional.

Configures the time after which a request to the origin will time out waiting for a response.

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
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_time_interval": "10m",
    "ves.io.schema.rules.string.min_time_interval": "10s",
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

- [origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033): complete subsection reference.

- [public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-1002100232220132-2323022211233331-1220132030312302-2222333312210131-3203022333311220-1311221201332302-1221221333110130-1020210013322232): complete subsection reference.

- [use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200): complete subsection reference.

<a id="canonical-2011120302333011-2111102311100110-3322132231113310-2013020012121311-2122303230332200-2011202231123232-0310220023203232-0031331110031013"></a>

## Next pages — origin_pool / 330133032131 / 5

- [origin_pool.more_origin_options](resources--cdn_loadbalancer--reference--group-012.md#canonical-0213200101331231-2120013110121131-1210303023133011-1030133313113032-3211300011331111-3213130011023120-2021320133300332-2231122333022102)
- [origin_pool.no_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0031320223020112-3322011231321300-2133213030302010-3003303230121300-2012120113121303-3110323013020211-2002003232313303-3223222032032013)
- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033)
- [origin_pool.public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-1002100232220132-2323022211233331-1220132030312302-2222333312210131-3203022333311220-1311221201332302-1221221333110130-1020210013322232)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0213200101331231-2120013110121131-1210303023133011-1030133313113032-3211300011331111-3213130011023120-2021320133300332-2231122333022102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100032232121320-0322202221312302-1213113110231203-0200020120211101-2130130211133010-2031230303132013-0221200331323310-1220102000000221"></a>

## origin_pool.more_origin_options — more_origin_options / 212131311220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.more_origin_options

<a id="canonical-3302202302203103-1033021301213130-3302230113222131-0033223301201022-1303001133313213-2010300301030201-0321310022221231-0103023322011133"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for more origin options.

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
more_origin_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0212211012001131-3223122310220033-1200120232022002-1022230331333301-0113112001301222-0113221231203021-1303231021230300-1100111321310113"></a>

## Direct properties — more_origin_options / 212131311220 / 3

<a id="canonical-2221000300122130-2121103112003221-1221332002221213-1332000023111032-0302331003113013-0121121132200112-1002320311120123-0000030031312132"></a>

<a id="canonical-1301132013110032-1302003230301331-0112130101223330-3310011023310310-1132123220010100-1312222022232311-0320133122032111-2322232301020101"></a>

## enable_byte_range_request property — more_origin_options / 212131311220 / 4

Type: `"bool"`. Optional.

Choice to enable/disable byte range requests towards origin.

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

<a id="canonical-1312220230113033-0103032332110020-2012223323133320-2201031211223023-0102111220021221-2021131211102002-3131111001322201-3312123330330233"></a>

<a id="canonical-2103033233130200-1322012302022320-2330131013320122-0211203011010301-2023212310010220-0200330332321032-3002221123111322-0320211123221030"></a>

## websocket_proxy property — more_origin_options / 212131311220 / 5

Type: `"bool"`. Optional.

Option to enable proxying of websocket connections to the origin server.

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

<a id="canonical-2311321010322113-2030323233323101-0303010010300010-1102110020302322-0303023200001232-1321130201313012-3201212113301121-2130003132003003"></a>

## Next pages — more_origin_options / 212131311220 / 6

- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0031320223020112-3322011231321300-2133213030302010-3003303230121300-2012120113121303-3110323013020211-2002003232313303-3223222032032013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012020212103113-1213300130113001-3131230311331130-3322313301123103-3003331103321022-1321132330301122-2320203033021311-1312323323201133"></a>

## origin_pool.no_tls — no_tls / 213213220312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.no_tls

<a id="canonical-1203001110313320-3220003320211313-0312022231002030-2200322111111013-3133321221023222-3022002203200312-0003011313033303-0020330230103222"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
no_tls = {}
```

<a id="canonical-1231112331210020-2330022322330303-3232000122120012-3221121120011023-1020311102220022-0220301230023100-3311031320020111-2222230011123321"></a>

## Direct properties — no_tls / 213213220312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221012000122223-0320020222322211-1330201113033011-2212232323121203-3333002102123012-3130320123000113-1111022310112232-2010213313120213"></a>

## Next pages — no_tls / 213213220312 / 4

- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111310013023320-0321112130211320-3233331021111003-0300111311212211-0131321003332021-3002020130323020-2120331320213012-0022102332132302"></a>

## origin_pool.origin_servers — origin_servers / 133222303320 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.origin_servers

<a id="canonical-2211120211003300-2321001131221201-3101131113113011-0233112112030003-3031011010231112-3120133123232213-1020310022000002-2213202033111130"></a>

Type: `"object"`. list nested block, Optional.

List Of Origin Servers. List of original servers.

Upstream description:

List of original servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("public_ip",
    "public_name")}
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
origin_servers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232202003331003-0011202133302203-0131020013120221-3112312032133202-2221122113020030-3230132003332033-3001123132030022-3002211333332312"></a>

## Direct properties — origin_servers / 133222303320 / 3

<a id="canonical-2321022210203333-1310212000331103-3313231200030223-1231322102231330-0203301233103221-2011213300200212-2233231020001013-1302110010001000"></a>

<a id="canonical-1203101030323222-0011112030131031-0113133210120102-2010200020131100-1032132311131111-1233102110132313-1330132102313233-0010222131020232"></a>

## port property — origin_servers / 133222303320 / 4

Type: `"number"`. Optional.

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

Upstream description:

Port the workload can be reached on Enter a custom port only if your origin server uses a
non-default port. Leave the value as 0 to automatically use 443 (TLS) or 80 (non-TLS).

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [public_ip](resources--cdn_loadbalancer--reference--group-012.md#canonical-3331223110233121-1002130131301031-3213012113300333-1110202110221311-3231013323001202-0130301030331302-0121001202011220-3002213210031303): complete subsection reference.

- [public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-3000211111131131-0021133131203022-2211200013220222-3101213003033323-3331012011322300-3320121322302033-3213312313302023-0313132230123213): complete subsection reference.

<a id="canonical-0203200302100312-1111300310310013-2231312110112220-1211213210012031-3331123120211302-1031200013102202-0133232231200000-3022321011220232"></a>

## Next pages — origin_servers / 133222303320 / 5

- [origin_pool.origin_servers.public_ip](resources--cdn_loadbalancer--reference--group-012.md#canonical-3331223110233121-1002130131301031-3213012113300333-1110202110221311-3231013323001202-0130301030331302-0121001202011220-3002213210031303)
- [origin_pool.origin_servers.public_name](resources--cdn_loadbalancer--reference--group-012.md#canonical-3000211111131131-0021133131203022-2211200013220222-3101213003033323-3331012011322300-3320121322302033-3213312313302023-0313132230123213)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3331223110233121-1002130131301031-3213012113300333-1110202110221311-3231013323001202-0130301030331302-0121001202011220-3002213210031303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012132133112101-0000320301302011-0321331122122302-2133012332212221-0103131202020233-2132101113112000-1020100032100101-0333220012231010"></a>

## origin_pool.origin_servers.public_ip — public_ip / 301322131103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033)
- origin_pool.origin_servers.public_ip

<a id="canonical-3200132003133110-1011032300320213-0303112220121030-0112013201112302-1310310021211202-3010023110132031-0032023313310131-3003232311201220"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2322312131211201-3123300332130121-1212020022230233-3312301120222302-1221301202331211-3010023000112222-2201000203303230-0011102022302022"></a>

## Direct properties — public_ip / 301322131103 / 3

<a id="canonical-2233033111012310-2303103213300232-3030013003220222-1100023201021130-1231221231002311-3103210310031000-2301002020300300-0333030332102333"></a>

<a id="canonical-2010320003321103-1020211012210033-3310003310230301-0321230112311223-1332230001212103-0131202213030120-3101000123320102-1223223313100030"></a>

## ip property — public_ip / 301322131103 / 4

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3010013202132302-0212320310213102-3203012320203233-1011201231032300-1210011200220211-3310112003022221-2203320133002212-1231301013120000"></a>

## Next pages — public_ip / 301322131103 / 5

- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3000211111131131-0021133131203022-2211200013220222-3101213003033323-3331012011322300-3320121322302033-3213312313302023-0313132230123213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012131312020303-3200030002301302-1030102210133202-2010311232033011-1022131101331323-3313331000202121-1113223231102031-1020020213333113"></a>

## origin_pool.origin_servers.public_name — public_name / 100320133313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033)
- origin_pool.origin_servers.public_name

<a id="canonical-1033012323010132-0020013200300212-1000233122311022-1122332310102333-3111012322230212-2321130323031023-1221323101003202-3121121330102003"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-3332202232130231-3223130112003311-2323201232013221-1310012122311111-3113100333302302-0211112212121212-2011321003112013-0030132330013310"></a>

## Direct properties — public_name / 100320133313 / 3

<a id="canonical-3333300330303300-1320021323010101-2131200311133003-0231132230223301-2000133010322202-0122110103001230-1003222311121130-0123300000013111"></a>

<a id="canonical-2010011210332021-0022313201031330-3020102323320213-3113300132131001-3220212121113203-1312001230012322-0000302330120103-1211133100212001"></a>

## dns_name property — public_name / 100320133313 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

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
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3132301011020133-2023031023112123-0111310122001310-0121200001201220-1130102113323220-3313301112031100-1322330033312233-1122032110011000"></a>

<a id="canonical-1233211303332333-3001033300132201-3001220132233013-0203112312303010-3311203013001231-2220131023202011-0302331021220102-1021303022100123"></a>

## refresh_interval property — public_name / 100320133313 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-1220112131223301-2201303013113323-3011113103222231-3233011121030301-0122321203020022-0103331201331222-0331200231122233-0212332301303221"></a>

## Next pages — public_name / 100320133313 / 6

- [origin_pool.origin_servers](resources--cdn_loadbalancer--reference--group-012.md#canonical-0222220003021200-1032332323111001-2020230223011121-2000031222321133-2000210331301301-1031223323223213-0201322023303000-2102200100010033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1002100232220132-2323022211233331-1220132030312302-2222333312210131-3203022333311220-1311221201332302-1221221333110130-1020210013322232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223310332033300-2330031111330133-2112233201033130-0013102030120030-0202031223222333-1112300211130013-2312120002211232-0102131020110131"></a>

## origin_pool.public_name — public_name / 003311201130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.public_name

<a id="canonical-0303202020030112-0203311230230001-1110330132311010-0233202000131121-2010001002333220-0020320310232100-2333112333000013-2133310230000213"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_name")}
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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201122012130013-2310301332321112-0112310020303122-2223103233200202-2130231313010111-1212132102023001-0313221011012132-3122033023030202"></a>

## Direct properties — public_name / 003311201130 / 3

<a id="canonical-0223300300303210-2023312333120132-1003113331102232-3222121330301001-3233212221022201-2302023201311031-2200300121222323-1002322221321022"></a>

<a id="canonical-0130233002131212-0110203102112031-0102111302132330-0031321222023210-3213021100200203-3032002211110301-2103033322211232-2211110211230212"></a>

## dns_name property — public_name / 003311201130 / 4

Type: `"string"`. Optional.

DNS Name. DNS Name

Upstream description:

DNS Name

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
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1210021301003202-3101010022320211-0101033203322133-3221313332200211-2030112103321313-2001100011322031-1023310133022220-1301110322031210"></a>

<a id="canonical-2302323320301000-2001231020011223-0133013322230203-2303102111331203-1210300012123101-0320301313020032-3330301213321022-0101312013303013"></a>

## refresh_interval property — public_name / 003311201130 / 5

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 10, Maximum: 604800},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-2213210320202221-2012202311000012-1233223312111013-1033100223113002-2101200101320013-0203012100023221-0332200132132121-2322002003100212"></a>

## Next pages — public_name / 003311201130 / 6

- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332310312012312-0100310233222021-1211002122131220-2223022323322311-3331032312222311-3332202000023321-0030031113010131-3312322201223200"></a>

## origin_pool.use_tls — use_tls / 130300012021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- origin_pool.use_tls

<a id="canonical-1031302012301011-3023111332123113-1303121212021210-1332102313013212-1130333010110110-2210222231231211-2332330323210030-0200023310323320"></a>

Type: `"object"`. single nested block, Optional.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_session_key_caching",
    "disable_session_key_caching"),
  validators.ConflictingObjectAttributes("default_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_session_key_caching",
    "max_session_keys"),
  validators.ConflictingObjectAttributes("disable_sni",
    "sni"),
  validators.ConflictingObjectAttributes("disable_sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "use_server_verification"),
  validators.ConflictingObjectAttributes("skip_server_verification",
    "volterra_trusted_ca"),
  validators.ConflictingObjectAttributes("sni",
    "use_host_header_as_sni"),
  validators.ConflictingObjectAttributes("use_mtls",
    "use_mtls_obj"),
  validators.ConflictingObjectAttributes("use_server_verification",
    "volterra_trusted_ca")}
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
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

Terraform syntax:

```terraform
use_tls {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123032011112112-3303131010113033-1323321330233201-2203120102032011-0212103201113112-1223021202230312-0313023312323212-1103130211013100"></a>

## Direct properties — use_tls / 130300012021 / 3

- [default_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-0302320330132121-2000211122231110-0233002332102313-0202233120013211-3231303120200033-3312032033120003-1320111203022003-2002321023302011): complete subsection reference.

- [disable_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-1122220121331121-3031112213221221-0221331232102133-0122100330030331-2133220122321030-1302320222123332-3123123132103002-3203320023030012): complete subsection reference.

- [disable_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-1002301103100330-2321023201300003-0320202012221111-2222033332010301-0123023001120110-2333020322010211-0211221333230303-2313203222223103): complete subsection reference.

<a id="canonical-2122000022123202-1321111232003120-3010213212130002-2331231321313322-2212122212211310-0230223333300030-1103120033112322-2023202302023021"></a>

<a id="canonical-1111302223222231-3220210233003200-2113120222110123-3023213000212321-2120331302022012-2203322031201323-0021031321010232-1202211131230203"></a>

## max_session_keys property — use_tls / 130300012021 / 4

Type: `"number"`. Optional.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(2, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0303111300201320-3230001233023023-3331123310130210-0122201131011003-2000202333110003-2300331000111120-1033201112021023-2020010313130103): complete subsection reference.

- [skip_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231030211300212-1000300232012123-2103303232131013-0111201013200331-0200133221231133-0311012120210312-3200300300011023-3232131131321320): complete subsection reference.

<a id="canonical-1201121023312011-2222203223203102-3130222221000033-2311113110013133-1030003103020312-2002200101033232-2222031031132110-2130210131303323"></a>

<a id="canonical-2130100211311111-3000323303221222-2020210210303222-1303100202123023-1222230331322122-3100112220132001-0012112212102321-1210231120300002"></a>

## sni property — use_tls / 130300012021 / 5

Type: `"string"`. Optional.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

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
    "format": "hostname",
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230): complete subsection reference.

- [use_host_header_as_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-3123303302303123-3202333112001312-3112222203202103-3033010031032020-2300210110111002-1031113222122101-1103113202002302-3331312230112131): complete subsection reference.

- [use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220): complete subsection reference.

- [use_mtls_obj](resources--cdn_loadbalancer--reference--group-012.md#canonical-2031020210010023-1201331221331233-3030210310323100-3110000032123022-2313130031331210-2233320022202102-2211203032032301-1102211332110031): complete subsection reference.

- [use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023): complete subsection reference.

- [volterra_trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-3132323021311131-3110113002100113-3203001123320003-0101112112132233-2030100310011022-2033133112113021-2220301213001231-2322101002133031): complete subsection reference.

<a id="canonical-1201201130232300-1303123010003230-0133230123003110-1121003012001022-0333100220221212-3000233002113110-3102223120032231-0030130111132202"></a>

## Next pages — use_tls / 130300012021 / 6

- [origin_pool.use_tls.default_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-0302320330132121-2000211122231110-0233002332102313-0202233120013211-3231303120200033-3312032033120003-1320111203022003-2002321023302011)
- [origin_pool.use_tls.disable_session_key_caching](resources--cdn_loadbalancer--reference--group-012.md#canonical-1122220121331121-3031112213221221-0221331232102133-0122100330030331-2133220122321030-1302320222123332-3123123132103002-3203320023030012)
- [origin_pool.use_tls.disable_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-1002301103100330-2321023201300003-0320202012221111-2222033332010301-0123023001120110-2333020322010211-0211221333230303-2313203222223103)
- [origin_pool.use_tls.no_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0303111300201320-3230001233023023-3331123310130210-0122201131011003-2000202333110003-2300331000111120-1033201112021023-2020010313130103)
- [origin_pool.use_tls.skip_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231030211300212-1000300232012123-2103303232131013-0111201013200331-0200133221231133-0311012120210312-3200300300011023-3232131131321320)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- [origin_pool.use_tls.use_host_header_as_sni](resources--cdn_loadbalancer--reference--group-012.md#canonical-3123303302303123-3202333112001312-3112222203202103-3033010031032020-2300210110111002-1031113222122101-1103113202002302-3331312230112131)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls_obj](resources--cdn_loadbalancer--reference--group-012.md#canonical-2031020210010023-1201331221331233-3030210310323100-3110000032123022-2313130031331210-2233320022202102-2211203032032301-1102211332110031)
- [origin_pool.use_tls.use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023)
- [origin_pool.use_tls.volterra_trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-3132323021311131-3110113002100113-3203001123320003-0101112112132233-2030100310011022-2033133112113021-2220301213001231-2322101002133031)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0302320330132121-2000211122231110-0233002332102313-0202233120013211-3231303120200033-3312032033120003-1320111203022003-2002321023302011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300100301021030-2112011100232102-3232102323001012-1322323323011311-0211202031031212-2012303213201310-1133020203333002-1313310331311320"></a>

## origin_pool.use_tls.default_session_key_caching — default_session_key_caching / 311012222030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.default_session_key_caching

<a id="canonical-3023201323332222-3231221130013103-3313312000011330-1233023102200133-1203232123200310-0233203202322033-1113023333222201-1130330231301333"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

Terraform syntax:

```terraform
default_session_key_caching = {}
```

<a id="canonical-3332112111332112-1002221021002021-3230010231102120-3232112231322113-2111202221302233-0220220310212020-3202001313310101-3001302023001222"></a>

## Direct properties — default_session_key_caching / 311012222030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001310101332103-1200301000311112-0102320021321200-1231102322332010-0020230023201000-0211323201303222-0020312221022122-2013101112320232"></a>

## Next pages — default_session_key_caching / 311012222030 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1122220121331121-3031112213221221-0221331232102133-0122100330030331-2133220122321030-1302320222123332-3123123132103002-3203320023030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331013022113333-1010233320213210-2320113100201033-1132221321231222-3032010232032211-1303011020123232-3220231220013023-3013320103230233"></a>

## origin_pool.use_tls.disable_session_key_caching — disable_session_key_caching / 101323200010 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.disable_session_key_caching

<a id="canonical-1110001133011330-3100221020030101-0011023211230030-0121010121122123-3230312110012112-1132313121031311-3110121122030111-3312011022101222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

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

Terraform syntax:

```terraform
disable_session_key_caching = {}
```

<a id="canonical-2310210030310302-1313003303302021-0130322110311000-0212230230102032-3122131102332022-0131013220303010-2111131100030103-3313201033113120"></a>

## Direct properties — disable_session_key_caching / 101323200010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320300013102202-2023310333323332-1331033200112132-3223012100123300-0212010311130111-1211010113023332-2313233202222003-2000322032321132"></a>

## Next pages — disable_session_key_caching / 101323200010 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1002301103100330-2321023201300003-0320202012221111-2222033332010301-0123023001120110-2333020322010211-0211221333230303-2313203222223103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332000003323312-2111232113031331-1021103311212321-1321300311131000-0321221220302321-0213233210133032-1121331210321110-2333111303000101"></a>

## origin_pool.use_tls.disable_sni — disable_sni / 323011012131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.disable_sni

<a id="canonical-0222101131313000-3012302320321113-1201221102302030-2222022000013302-2013021132013220-0233213033002322-3003221122221203-0222002221130131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

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

Terraform syntax:

```terraform
disable_sni = {}
```

<a id="canonical-0113112113332320-1220233110031112-2020002130010101-2230131303100320-1012330331011300-2101322230010001-2212211300131113-0101302233101013"></a>

## Direct properties — disable_sni / 323011012131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2032003101310033-2120111312020031-1003013222222320-1303312012100111-2021012201133231-0001332103213021-0311111221020323-3122302300001013"></a>

## Next pages — disable_sni / 323011012131 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0303111300201320-3230001233023023-3331123310130210-0122201131011003-2000202333110003-2300331000111120-1033201112021023-2020010313130103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332302102103112-0301032020121312-2310020321012102-2210323003312221-2102200322112133-0203100033313323-3012300210301002-0222302131003222"></a>

## origin_pool.use_tls.no_mtls — no_mtls / 213200103213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.no_mtls

<a id="canonical-0222232101212200-0131011321333311-0211030210033330-2121103101122231-2232231103003013-3111320133021011-3321101113330122-2320000132331230"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

Terraform syntax:

```terraform
no_mtls = {}
```

<a id="canonical-3001303301133300-0102310123312203-3000312210313212-2220312332032001-2021233201010021-2003022003233101-1112130211001230-3113133222033322"></a>

## Direct properties — no_mtls / 213200103213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031102311210120-2011130130310100-1111122222131130-2003313333100002-2231311301102323-0222213131103033-3032231120200020-2011133120221200"></a>

## Next pages — no_mtls / 213200103213 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0231030211300212-1000300232012123-2103303232131013-0111201013200331-0200133221231133-0311012120210312-3200300300011023-3232131131321320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300101321130200-2022330321030102-1331012012133130-0122013010001233-0231233321203300-0001210021300330-3210201100021023-1211222311332000"></a>

## origin_pool.use_tls.skip_server_verification — skip_server_verification / 320202310131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.skip_server_verification

<a id="canonical-0012012332030002-0112333313030103-2300102222012323-1331300001232100-2120301300121212-0302223312332130-1313110233122330-1122203200223021"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
skip_server_verification = {}
```

<a id="canonical-0120002122002310-3021203122321131-1200003033302010-3132110013113330-2203011303030111-2133023333112002-3232120320213122-3132023332330001"></a>

## Direct properties — skip_server_verification / 320202310131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3213211201033311-0132111221021003-3233101321300220-1102333022122321-2311231113231230-0200010321003213-0302302010320201-1232320300023132"></a>

## Next pages — skip_server_verification / 320202310131 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203301000023331-1233310011210300-0020111323222033-1313001300221001-2003233220231113-3333230213113310-3300300313103110-1311023111000203"></a>

## origin_pool.use_tls.tls_config — tls_config / 100211113020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.tls_config

<a id="canonical-2010031003200021-3000332113323202-1000031233322211-2113212223303132-3332222022033200-0322320111203222-0331122030202032-3110030101012311"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303013233232022-3230310132131300-2013130221330110-0000320100222211-2201232133100112-0131032312302013-2121310232022011-1232230331122301"></a>

## Direct properties — tls_config / 100211113020 / 3

- [custom_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231112012220003-2322021233233212-2232322232013110-3013300211332211-0231000111131313-2101110200020213-1120133331232301-3312221110130122): complete subsection reference.

- [default_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-2131312123312001-3112012002301233-1322130102031332-2232020023022201-3210013231210231-0132023010112132-3212133120110112-3302010133231023): complete subsection reference.

- [low_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-3020300200001212-2220223302331012-2101011113033301-3002031232333231-0001001301131332-0303300111211310-3330303303031013-3202313313203201): complete subsection reference.

- [medium_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-0333113000022310-0232221001030222-3020002200123203-2001313101112132-3203303110010030-1322300331321313-3011021302333320-0102101222312111): complete subsection reference.

<a id="canonical-0102103021220323-0123003221100100-3311213220233220-1230223223121202-1320023012223023-3222011000303131-0031032213303122-1310311331321022"></a>

## Next pages — tls_config / 100211113020 / 4

- [origin_pool.use_tls.tls_config.custom_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-0231112012220003-2322021233233212-2232322232013110-3013300211332211-0231000111131313-2101110200020213-1120133331232301-3312221110130122)
- [origin_pool.use_tls.tls_config.default_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-2131312123312001-3112012002301233-1322130102031332-2232020023022201-3210013231210231-0132023010112132-3212133120110112-3302010133231023)
- [origin_pool.use_tls.tls_config.low_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-3020300200001212-2220223302331012-2101011113033301-3002031232333231-0001001301131332-0303300111211310-3330303303031013-3202313313203201)
- [origin_pool.use_tls.tls_config.medium_security](resources--cdn_loadbalancer--reference--group-012.md#canonical-0333113000022310-0232221001030222-3020002200123203-2001313101112132-3203303110010030-1322300331321313-3011021302333320-0102101222312111)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0231112012220003-2322021233233212-2232322232013110-3013300211332211-0231000111131313-2101110200020213-1120133331232301-3312221110130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133110331103332-0003202212210010-0103022130323331-2010130133112012-0133022000202131-2311120003032111-2111010333031133-2233310002312323"></a>

## origin_pool.use_tls.tls_config.custom_security — custom_security / 202312222101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- origin_pool.use_tls.tls_config.custom_security

<a id="canonical-1200020030320103-1330121231321310-2303323203103013-3101121031112010-3000031332022203-0331322023302122-3113021323013323-0233203203132000"></a>

Type: `"object"`. single nested block, Optional.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("cipher_suites")}
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
custom_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202300121331021-1232111220021003-2131013010202101-1113200210301001-0023011022022322-2112102231211112-3131322211311000-2331101332322321"></a>

## Direct properties — custom_security / 202312222101 / 3

<a id="canonical-0003001210322330-1120012102322313-1001222322000313-2033333202200232-1002210210111112-2123101131131222-0102002121123010-0101322301202301"></a>

<a id="canonical-2010122010233020-2102200121313232-0210032301103010-0323333001232310-0220103312131112-2112001331202012-1110111322023011-2211113300120013"></a>

## cipher_suites property — custom_security / 202312222101 / 4

Type: `["list", "string"]`. Optional.

The TLS listener will only support the specified cipher list.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2002012333201002-0031001323111320-1331310111120020-1130333200112302-3033213322302311-1312321310333230-2123301330112211-0211013001200130"></a>

<a id="canonical-2021313033121130-0232021123113132-3133301121020000-2102232122001232-2112322001333220-1313100113112222-1310103323100131-0300312331203311"></a>

## max_version property — custom_security / 202312222101 / 5

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1122331132303012-3300222303232133-0223333121300212-1213111220030220-0223023320021021-0331333003211223-3120130300030013-3012020222313212"></a>

<a id="canonical-3230323220000332-3313201100130020-0003310133311123-1020130123200223-1002310030230023-0123321120112233-0010231211001102-1200313012221012"></a>

## min_version property — custom_security / 202312222101 / 6

Type: `"string"`. Optional.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1021333223132233-2030202121223131-1100120200301021-0031223203211303-3120321310313302-3101002002032021-0031320131013021-1300001312101102"></a>

## Next pages — custom_security / 202312222101 / 7

- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2131312123312001-3112012002301233-1322130102031332-2232020023022201-3210013231210231-0132023010112132-3212133120110112-3302010133231023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003212213000213-2231010230022110-1102323020011030-2023211100120123-3203003220103231-1121120013011130-2322110310122130-1331112023221322"></a>

## origin_pool.use_tls.tls_config.default_security — default_security / 233322312220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- origin_pool.use_tls.tls_config.default_security

<a id="canonical-0133102230011213-3132310020000202-0200030231112211-1310332331020002-2221002323323232-3112323100031020-1331030201331111-2033210132300110"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_security = {}
```

<a id="canonical-1012102103131002-1001220012123113-2112032202212301-3211031233313222-3120101002310021-3000123321200132-0322000032122032-0232231200233201"></a>

## Direct properties — default_security / 233322312220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113220222002201-1201132020312110-0012100213112312-1201132320100111-3030002001033300-2101202221032010-1303000321102100-1001120002210331"></a>

## Next pages — default_security / 233322312220 / 4

- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3020300200001212-2220223302331012-2101011113033301-3002031232333231-0001001301131332-0303300111211310-3330303303031013-3202313313203201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322013010132102-1000230031012132-3332231302202220-1101003001123201-2013120222322133-1133333212303003-3011322233003102-0021032322100121"></a>

## origin_pool.use_tls.tls_config.low_security — low_security / 123210233202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- origin_pool.use_tls.tls_config.low_security

<a id="canonical-3130001032321102-3001021322030130-1132132220232300-0123112023233103-0022322023000220-1123010130100313-2220200322301130-3130211130013210"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
low_security = {}
```

<a id="canonical-2310003311022030-1221020231100022-0333321010210322-0323231131103032-2023232202001130-1132121223001333-2200012002022033-0321030200132330"></a>

## Direct properties — low_security / 123210233202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210003021233332-1211132102310202-3011321120010131-3231311103002303-2133210233120222-3323203321302002-1002300300130302-1120202310111312"></a>

## Next pages — low_security / 123210233202 / 4

- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0333113000022310-0232221001030222-3020002200123203-2001313101112132-3203303110010030-1322300331321313-3011021302333320-0102101222312111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333320132231233-2233003123203212-0202122302012003-0113103032331032-2220100031232020-0101300111130312-1100331001332311-1233333123032330"></a>

## origin_pool.use_tls.tls_config.medium_security — medium_security / 202303003202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- origin_pool.use_tls.tls_config.medium_security

<a id="canonical-0032002022100211-3011020022221031-1313211322013003-0300233120223012-0003232111313032-1032110130331010-3001201212122000-3111331121211010"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
medium_security = {}
```

<a id="canonical-3333302111110110-3111332033102322-3200031033303310-1112121023320123-2001110100130120-3300320211233113-3030210321012010-0223233210132122"></a>

## Direct properties — medium_security / 202303003202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103301230321003-0333212133321010-0000032221111131-2230003110102022-1010012202220011-1323332321311001-0101311322131031-2013332133201212"></a>

## Next pages — medium_security / 202303003202 / 4

- [origin_pool.use_tls.tls_config](resources--cdn_loadbalancer--reference--group-012.md#canonical-1133233302132311-1012222223313010-3302223302302330-0333310112132201-1012200011320113-2012300301111313-1303000121210032-0302211210120230)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3123303302303123-3202333112001312-3112222203202103-3033010031032020-2300210110111002-1031113222122101-1103113202002302-3331312230112131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001332102013301-0032031221322123-2111231323122202-1301023303021130-2212013031012320-0011022002102313-1321222322223020-0303210003012301"></a>

## origin_pool.use_tls.use_host_header_as_sni — use_host_header_as_sni / 200320333132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.use_host_header_as_sni

<a id="canonical-3333103210230101-0131313302121200-1333330203210022-0132011222100001-1021010232010111-0031211000301313-2101133122313100-2001313031120123"></a>

Type: `["object", {}]`. Optional, Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

Terraform syntax:

```terraform
use_host_header_as_sni = {}
```

<a id="canonical-1031120003011130-3303003103302132-2311012332332031-1221013103200033-0320033331300021-2320302130313012-0331031021213303-1201223120302320"></a>

## Direct properties — use_host_header_as_sni / 200320333132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203023332102023-0330000101230212-2210310013300010-2231202032333322-1311001012301300-3023020313021202-0232122121032221-2013132001130333"></a>

## Next pages — use_host_header_as_sni / 200320333132 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213301202213132-0133113322211111-0332011122210010-0033021302322130-1322033321002020-2001110211100203-1010113111121030-1110303032233331"></a>

## origin_pool.use_tls.use_mtls — use_mtls / 332011212121 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.use_mtls

<a id="canonical-0311031330111112-1010002223111221-2233002030001110-0333003232333203-2203302030300101-1322203130301121-2112200032222133-3031023131312303"></a>

Type: `"object"`. single nested block, Optional.

MTLS Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates")}
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
use_mtls {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311100321320120-0232110132110310-0311003011121321-0233100322021301-1203033223010320-2323200322302022-2101230130021303-0110022012220233"></a>

## Direct properties — use_mtls / 332011212121 / 3

- [tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020): complete subsection reference.

<a id="canonical-0211333323320232-3003021001232313-2320332221231111-0233101100323013-2331220200113313-1220023113223330-1222303233020003-3013313321132331"></a>

## Next pages — use_mtls / 332011212121 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202331212332112-1121302111102120-1312222010031122-3131102303211123-0132130031232020-0212131220200223-3021021113032120-3013033110103212"></a>

## origin_pool.use_tls.use_mtls.tls_certificates — tls_certificates / 201110221203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- origin_pool.use_tls.use_mtls.tls_certificates

<a id="canonical-1321030022322100-2222213021121020-1113320103000223-2021312132133301-1012210300010031-3000230100031320-2121311223310003-0032231323131013"></a>

Type: `"object"`. list nested block, Optional.

MTLS Client Certificate. MTLS Client Certificate.

Upstream description:

MTLS Client Certificate.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("certificate_url"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "disable_ocsp_stapling"),
  validators.ConflictingListObjectAttributes("custom_hash_algorithms",
    "use_system_defaults"),
  validators.ConflictingListObjectAttributes("disable_ocsp_stapling",
    "use_system_defaults")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-0203100331210131-3201002302321223-1030311302011012-2210210320210100-0212033002111220-1001011000021032-1110331001012022-2313110313230021"></a>

## Direct properties — tls_certificates / 201110221203 / 3

<a id="canonical-3131123230300102-2303310123021231-0222013212213331-3101001202211223-0200202333322331-0302311211220331-2320230321331301-3032301333321102"></a>

<a id="canonical-3131310320101201-3113023300012120-0033111013221130-1132033310131112-3203320003112310-2011031120031112-1133301301130101-3100232131003313"></a>

## certificate_url property — tls_certificates / 201110221203 / 4

Type: `"string"`. Optional.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Upstream description:

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--cdn_loadbalancer--reference--group-012.md#canonical-2012310320122332-0012302201101121-1032312111313220-1113220311013313-0100313221310121-1031230003022000-2200023202330130-2230131201302112): complete subsection reference.

<a id="canonical-2002002010323003-3323213311232112-2033020330003313-3223021311131111-1233121211003320-0320220011122013-0333312020310203-3000102121233300"></a>

<a id="canonical-2110201120113110-0230102301032102-2232031013322100-3001011101203000-2013203030311323-1202303323013323-2001022330121021-3111330202212011"></a>

## description_spec property — tls_certificates / 201110221203 / 5

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--cdn_loadbalancer--reference--group-012.md#canonical-0030000012212300-1203012233022132-3211102131102220-3312301020230120-1130301312112300-2123323020031311-1101031022003322-3131131211031002): complete subsection reference.

- [private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311): complete subsection reference.

- [use_system_defaults](resources--cdn_loadbalancer--reference--group-012.md#canonical-0321211001212221-2201311223311330-0203120023100113-3102023123032112-3132003033033032-1313231223221223-0311320211032220-1102210003202231): complete subsection reference.

<a id="canonical-2003202200131012-1022223102201220-0301123121033003-0320123330132023-2033232312100010-3300022310202330-2002003023013323-3232020213203112"></a>

## Next pages — tls_certificates / 201110221203 / 6

- [origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms](resources--cdn_loadbalancer--reference--group-012.md#canonical-2012310320122332-0012302201101121-1032312111313220-1113220311013313-0100313221310121-1031230003022000-2200023202330130-2230131201302112)
- [origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling](resources--cdn_loadbalancer--reference--group-012.md#canonical-0030000012212300-1203012233022132-3211102131102220-3312301020230120-1130301312112300-2123323020031311-1101031022003322-3131131211031002)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311)
- [origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults](resources--cdn_loadbalancer--reference--group-012.md#canonical-0321211001212221-2201311223311330-0203120023100113-3102023123032112-3132003033033032-1313231223221223-0311320211032220-1102210003202231)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2012310320122332-0012302201101121-1032312111313220-1113220311013313-0100313221310121-1031230003022000-2200023202330130-2230131201302112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321321221213113-1020212011101231-0120013032231021-1121101310010222-3323031123312110-1311221203213103-1011130011123130-3300302203112100"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms — custom_hash_algorithms / 203202220012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- origin_pool.use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-3323203333100332-1322332222201202-1131302031000213-0110333302133323-3012100332232120-3111121023031111-2312213211030010-2223000000303123"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_algorithms")}
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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103010303112110-2023310221302210-1331012121102130-3223323223133311-2003233133123311-0222230330303302-2332310220010002-1200222322300220"></a>

## Direct properties — custom_hash_algorithms / 203202220012 / 3

<a id="canonical-2010220211220301-2332312122112303-3223120000033312-2021003300002032-0012111222002132-1110301010011311-3303332021312011-0332232212131303"></a>

<a id="canonical-1111233000133122-3001333310021311-2010032022023202-1231300010001232-2233331232312002-0130131232330101-1212323112131222-2211320003230131"></a>

## hash_algorithms property — custom_hash_algorithms / 203202220012 / 4

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Upstream description:

Ordered list of hash algorithms to be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0210232210301103-2132001010132300-1301100210310312-2020011202302110-0331333200222312-2220203112030323-3220111113302122-1210110320032321"></a>

## Next pages — custom_hash_algorithms / 203202220012 / 5

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0030000012212300-1203012233022132-3211102131102220-3312301020230120-1130301312112300-2123323020031311-1101031022003322-3131131211031002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323022102130331-0030310331202021-1302322111001022-3113103002031311-1131010213300100-3010321113121211-2120310322201202-1203221210220231"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling — disable_ocsp_stapling / 023232321220 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- origin_pool.use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-1000321223320222-2000312331113232-0312211332033021-1110100001010033-1001111133001303-1212012123013001-0102132300010210-2200222302132303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

<a id="canonical-2113330131131301-2233122213100023-0312330202020211-1032222010101033-3200313203032011-3032110022310122-0312322203233112-3110333223002012"></a>

## Direct properties — disable_ocsp_stapling / 023232321220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211333220113311-2222301311002323-3210333320333232-3112122112221101-3301131222103202-0020231122130113-0330332302312030-2112311202203112"></a>

## Next pages — disable_ocsp_stapling / 023232321220 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310022120033101-2200213330020203-1230301323323000-3213222332110311-0011023212220032-1123311131033311-3322232211310323-1120031310011223"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.private_key — private_key / 333112323302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-3333131020323100-0110303201002010-0333022132323301-1122333202020131-1123311321223013-0213122101233301-1101320312012102-2022233220130012"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0330202033131032-2323012332233300-2221211220321022-3311000221000101-0101122012230031-0201322222103000-2312001013321121-1311121033300201"></a>

## Direct properties — private_key / 333112323302 / 3

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-3212223001330012-3221100332133323-3210032013321303-1203001200102211-3212212013121212-0313221102232330-2232002100101321-1232013222211102): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-3010101223303031-2320230222331211-1100320210230202-3321110030031210-3102211011003001-2303331301312101-2022332212120033-1122212000123130): complete subsection reference.

<a id="canonical-1002011212023202-3221302032122120-1210022022210321-2000323121230221-2100323022312313-3210100301200332-1103233300133223-2323212032132321"></a>

## Next pages — private_key / 333112323302 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-3212223001330012-3221100332133323-3210032013321303-1203001200102211-3212212013121212-0313221102232330-2232002100101321-1232013222211102)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info](resources--cdn_loadbalancer--reference--group-012.md#canonical-3010101223303031-2320230222331211-1100320210230202-3321110030031210-3102211011003001-2303331301312101-2022332212120033-1122212000123130)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3212223001330012-3221100332133323-3210032013321303-1203001200102211-3212212013121212-0313221102232330-2232002100101321-1232013222211102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232300332131102-2332333033031312-3311011130100211-2310013023010000-3333201300121211-2233312301121030-1033213133121212-3031301133202033"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info — blindfold_secret_info / 211233302302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-2331131303101201-2232012000203022-2313312210031022-2303030003303012-0311030012301231-1133233022303210-3232003120121210-3230303023322200"></a>

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

<a id="canonical-1112232332121132-2223331002220001-2323211100321330-1312201002221112-3203220102302300-0102122211312332-0332213133322303-3221302133101133"></a>

## Direct properties — blindfold_secret_info / 211233302302 / 3

<a id="canonical-0002233003331020-0300323231333131-3321231212212202-3100023102001002-0000233201210321-1033132302213033-1102202223311220-0331110301030032"></a>

<a id="canonical-0201213313232330-2233212330001232-2132210111010320-0020230122202022-1200033313113312-0330202102012223-1112130231130112-0311210232323211"></a>

## decryption_provider property — blindfold_secret_info / 211233302302 / 4

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

<a id="canonical-0012122322121213-1103100320122022-1311233023211211-0220003232103333-3221310303113101-0332133220202330-1310033012113021-3201201302113033"></a>

<a id="canonical-1232113132221113-2313010210313030-3121303210312123-0100331322021130-3323021110100003-3031312001003122-3131111023120121-0330223133133322"></a>

## location property — blindfold_secret_info / 211233302302 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

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

<a id="canonical-2301100211103023-0310131022111003-0221033310210211-1333003332110013-1022312012300102-0211103022131222-0213322202022002-3033310223313031"></a>

<a id="canonical-2122221033232121-3000121223110121-2202011002201232-3302310003013330-1123031103300132-0201213032021301-3021133230122101-1121133312131330"></a>

## store_provider property — blindfold_secret_info / 211233302302 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

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

<a id="canonical-1302312233212320-1201100201222003-3103312030000333-1112113313320032-1322000223010120-0130201232100030-2022311300101032-1132202131012131"></a>

## Next pages — blindfold_secret_info / 211233302302 / 7

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3010101223303031-2320230222331211-1100320210230202-3321110030031210-3102211011003001-2303331301312101-2022332212120033-1122212000123130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133302232123130-1321012322123310-0010110013131213-3223321123220211-3033003122212331-3132133230232030-3013233321223230-2023132030231223"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info — clear_secret_info / 332001321001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311)
- origin_pool.use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-0223231001120000-2220203332332221-0012000213303032-0022023030203310-0013321022320111-0031111130002023-3033302131333202-1300233320330220"></a>

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

<a id="canonical-3313223220323300-1300210330301103-3212300003231221-2320032302302232-2122202213313102-1222233310233313-3221322202221230-3231323012031020"></a>

## Direct properties — clear_secret_info / 332001321001 / 3

<a id="canonical-2032111223032301-0111330313010311-0220023123130323-1130110212231010-3002113111121310-1320321333011322-3321121231123202-0111302101203211"></a>

<a id="canonical-1010203322233012-1133320222202332-3323213011130232-3223233102121121-0102032133002213-3022102233003120-0321122330301012-2330000012001021"></a>

## provider_ref property — clear_secret_info / 332001321001 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1022211131012012-2000233202323231-3000221013011201-2102332130120101-3032121121110023-2133320331312303-0113010320120320-3213213220233021"></a>

<a id="canonical-2222031102312133-1202033132021232-2330220002333321-2231031000102020-3030312010000300-3112002101023030-0322132322130133-3120101332320023"></a>

## URL property — clear_secret_info / 332001321001 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1023133103230030-3200022222010213-3222333300031031-0121031323021002-1011203321221331-0123202113101103-3210133231113012-1200300032131210"></a>

## Next pages — clear_secret_info / 332001321001 / 6

- [origin_pool.use_tls.use_mtls.tls_certificates.private_key](resources--cdn_loadbalancer--reference--group-012.md#canonical-2112130302003010-1130030023133020-0311013233011003-3312330223120133-1300231200222310-1231303312112132-1223133032222023-2222012202020311)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0321211001212221-2201311223311330-0203120023100113-3102023123032112-3132003033033032-1313231223221223-0311320211032220-1102210003202231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302123133012111-1023103111210210-0330302212002123-0010111023310000-3110312121321322-3020230102100211-2211132333231013-1232132233323333"></a>

## origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults — use_system_defaults / 302303113120 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_mtls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100010311320131-0032032122322202-0311100020012221-0320030012002301-1312010223222223-3321232031013211-1133132121111022-3323031301033220)
- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- origin_pool.use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-1211300023330002-0331112112113031-0003210101323130-1222232323310333-0110303302001103-2131331233202030-0103011113331300-1020102212322012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

<a id="canonical-2020031012223111-1123330023311323-3010123121131110-2120103311010030-0112103121200221-0300002122110101-3320110001300333-2220020132231120"></a>

## Direct properties — use_system_defaults / 302303113120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222200221132031-1023201223023332-0300313102001222-2201020233021223-1300203023122120-3120013022101132-0212112100330022-2333211331313220"></a>

## Next pages — use_system_defaults / 302303113120 / 4

- [origin_pool.use_tls.use_mtls.tls_certificates](resources--cdn_loadbalancer--reference--group-012.md#canonical-2002223012102211-1302201110132120-3212230022121011-2230322302322011-0230030020120120-1232123021102023-2101133322202322-0302113200003020)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2031020210010023-1201331221331233-3030210310323100-3110000032123022-2313130031331210-2233320022202102-2211203032032301-1102211332110031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313312000321331-2113331002000120-2232211033301212-2100110031333003-3200333133133021-2100331000212010-3231221210011212-2302300103232233"></a>

## origin_pool.use_tls.use_mtls_obj — use_mtls_obj / 232223233211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.use_mtls_obj

<a id="canonical-1210331023213212-0010002012301121-0013301313330320-3012011213022320-2331322321312331-0223112022133322-0023012302222300-3312132331203213"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
use_mtls_obj {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000303110012033-1002211001330201-2222222100313032-3112113331103231-3010321232030231-2201110023233302-2121332123333002-2003130300301330"></a>

## Direct properties — use_mtls_obj / 232223233211 / 3

<a id="canonical-1133012123113013-2121323023101021-1013310101222200-3220021011331210-1230130131211131-1103331000010003-0123213202001303-2222130102302232"></a>

<a id="canonical-2122231200023221-2201313030210102-2332013113020000-3230221110203010-0330100111302111-3311203231112031-2323102121312222-1011212133202311"></a>

## name property — use_mtls_obj / 232223233211 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-0132321313033030-0111211332300022-3320303032002212-2322331213122001-3023213031321220-2222331103213211-2121233303332232-2223130000203133"></a>

<a id="canonical-3120223220210311-3002103120100311-1020300233013031-0222321101333230-0313230201202202-0013211221012000-1211111311232230-1001102320113132"></a>

## namespace property — use_mtls_obj / 232223233211 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1201331222011300-2311211121323033-2103101321101021-0331001011120133-2331332222230133-3030030211120032-0203100321020211-0200030300331221"></a>

<a id="canonical-2221020211311233-3131112130021013-3030212001333012-2310022010010032-1322233030001021-3030303033200332-0203332300213001-2121101320231200"></a>

## tenant property — use_mtls_obj / 232223233211 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3332312221323022-0321133120101101-2023031131202323-1120301122122213-3000311013323022-3020003223131223-0122001313120133-1331133202100331"></a>

## Next pages — use_mtls_obj / 232223233211 / 7

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310231212322113-0220122301031302-1222131010132201-1311130201302320-1203032330312020-0011330330120233-0233303031200001-1131323113120032"></a>

## origin_pool.use_tls.use_server_verification — use_server_verification / 333301201300 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.use_server_verification

<a id="canonical-2220130033132213-0011021310023330-1102010203300320-2123111000023020-3032130011203121-0102210310033012-2333032132133221-0111100323132103"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302200312001023-1210122322331133-3303211132030003-1102202111132113-0313222013323100-0113231000212312-2000022321230331-3033003322010013"></a>

## Direct properties — use_server_verification / 333301201300 / 3

- [trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-3002213023021322-0133011031123120-1210313002112221-2202330030331233-2100222232331003-1023130312302130-0033300311121201-3021022123321300): complete subsection reference.

<a id="canonical-3012133003230110-3122032221312322-1201031021010123-1202223303033323-1012002131300311-2233220201012103-3232230202322233-0310300302333023"></a>

<a id="canonical-1330323310033201-2333300020220211-0212103101021301-3023111010320122-2011112332320202-0020311212100223-3213223321210121-1312100031120202"></a>

## trusted_ca_url property — use_server_verification / 333301201300 / 4

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-1332123013112001-0331021133020110-1112201113231311-0113232313323113-0020101003001320-1002012221310033-1001001303210101-0103001132010131"></a>

## Next pages — use_server_verification / 333301201300 / 5

- [origin_pool.use_tls.use_server_verification.trusted_ca](resources--cdn_loadbalancer--reference--group-012.md#canonical-3002213023021322-0133011031123120-1210313002112221-2202330030331233-2100222232331003-1023130312302130-0033300311121201-3021022123321300)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3002213023021322-0133011031123120-1210313002112221-2202330030331233-2100222232331003-1023130312302130-0033300311121201-3021022123321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210212023002231-2321210002112023-2100330303113023-2331032200323100-0321120120023213-2133303233002201-3012013000312031-0323110300303013"></a>

## origin_pool.use_tls.use_server_verification.trusted_ca — trusted_ca / 111220003221 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023)
- origin_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-2303233022001300-2311022331230210-3111001002332233-3330113303203101-0213231122203020-1012332002020133-0112102000002220-1223010221311102"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123300223303230-0322230021202002-3203031100200020-3022330102333203-0323221210101120-0021100213012310-1002101121301032-1031312120333321"></a>

## Direct properties — trusted_ca / 111220003221 / 3

<a id="canonical-0031000020110300-2132200001212233-2033132003332201-1021121313123231-3112130210331312-1310310203001301-1301011101311211-0022010203001200"></a>

<a id="canonical-2200311131221330-1101121333333201-1023020330001332-3221323002101333-0121303011300003-1132213121032223-3320031222023333-2113102323112231"></a>

## name property — trusted_ca / 111220003221 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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

<a id="canonical-2313133021321132-2223122331210320-3311110013020013-0323221300220122-2131100131021200-1131102002201022-1231212122233001-1112001333112301"></a>

<a id="canonical-1210333231230022-2322213023002020-1012231121101223-1102020231100322-2130233200033000-2031311033332311-2020002100312130-2323200332211012"></a>

## namespace property — trusted_ca / 111220003221 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2213330020231010-2112030003303333-2313001320031132-1102331220232123-0313002213032320-2122220101103210-2131303002311300-3112303311222012"></a>

<a id="canonical-3112200211202132-3120321302233013-2023103333133032-0013012123100321-3122100220010310-1120031113333021-0003303102101131-3300102001212313"></a>

## tenant property — trusted_ca / 111220003221 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2130220003310111-2120223313213011-2112033012321023-3110013303103121-0231212020130001-1123322200230222-1002000113330021-1021232322132032"></a>

## Next pages — trusted_ca / 111220003221 / 7

- [origin_pool.use_tls.use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3132323021311131-3110113002100113-3203001123320003-0101112112132233-2030100310011022-2033133112113021-2220301213001231-2322101002133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
