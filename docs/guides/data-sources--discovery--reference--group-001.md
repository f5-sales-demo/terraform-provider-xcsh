---
page_title: "xcsh_discovery reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery reference."
---

# xcsh_discovery reference

<a id="canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- Property reference

<a id="canonical-0011031322112331-2212002311010212-2232221330210332-3212321213222230-3200312200110131-2222212122132232-1132133112333122-0311222231000331"></a>

### Direct properties for `xcsh_discovery`

<a id="canonical-0013323033213331-3000331103203323-1201121220302100-2231011130132230-3021023133021202-3230333200310103-1231101202003323-0000133331130220"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-1203102113122022-1221113201311023-0103311123213210-0113302110230113-2222111313120011-3020032120031223-2322233020320211-2210312120011333"></a>

<a id="canonical-2000231310100022-2120120221000212-3310132231122031-3311221313123331-2210113033302312-2212113012230012-2203210101100211-2320101110213222"></a>

#### `cluster_id` property

Type: `"string"`. Computed.

\[OneOf: cluster\_id, no\_cluster\_id; Default: no\_cluster\_id\] Exclusive with \[no\_cluster\_id\]
Specify identifier for discovery cluster. This identifier can be specified in endpoint object to
discover only from this discovery object.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

OneOf alternatives in this subsection:

- [cluster_id](data-sources--discovery--reference--group-001.md#canonical-1203102113122022-1221113201311023-0103311123213210-0113302110230113-2222111313120011-3020032120031223-2322233020320211-2210312120011333)
- [no_cluster_id](data-sources--discovery--reference--group-002.md#canonical-2330131222130101-1223121232130000-3310203111002300-3111030300122323-2032022200033233-3120020220301310-3232021301122000-1013013012212302)

Select alternatives according to the provider validators above.

<a id="canonical-0131301302223120-0212113100311233-2030310012032010-0120213012033000-0030221210031333-1220133033123033-2321031031130012-1210320103231111"></a>

<a id="canonical-2013133130010103-2213000211211302-1112231100301301-0030221010230110-2013310331110200-2010011030003131-0320013312203310-0313021210102101"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the Discovery.

Additional upstream details:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110): complete subsection reference.

- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322): complete subsection reference.

<a id="canonical-0333321230233213-0031212000100031-3302111231213110-1022123311301233-3222211330101210-3212333313000311-3311031310213013-2032212220310302"></a>

<a id="canonical-0101330232323133-2000120110223012-2120223122211121-3233120323002133-3123102230202011-1300202130303203-0010110330300002-0232301013112333"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-3223222201123202-2332112000233020-3233202330221211-3101030202222121-3030023203331030-1012001130303032-1220002230013103-0130030020033122"></a>

<a id="canonical-1223223332201312-0131011323230130-2303333333030032-3222232231100012-1011003100210311-2301021013320212-3331112020332211-1030221223212212"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

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

<a id="canonical-3233301022022013-1000030032313010-2121301132231033-1332111121212003-2300212131221013-2230320223130210-3201023313001323-3022323213130032"></a>

<a id="canonical-1112113131001111-1323131232313120-2131120233031013-3203112111311131-1032331003220200-3311211131302011-3121330001332032-1000031100120223"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Discovery.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0302231212213212-3312212232232223-2311232330102233-1332211003203321-0223011201033122-1220022021223333-3032110002302000-0302010130030231"></a>

<a id="canonical-0203113000013331-2033322033223101-1303020221231311-2033021022302001-2030233203331120-2101123333232133-1332010210310102-2220312131322110"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Discovery exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [no_cluster_id](data-sources--discovery--reference--group-002.md#canonical-3022232112013312-1120023000331333-0002001130200231-2023010022021302-2321320033321023-1313213310100331-3000313221232320-1132122300112130): complete subsection reference.

- [where](data-sources--discovery--reference--group-002.md#canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110): complete subsection reference.

<a id="canonical-3010011012201330-1123002000113110-1113301303102013-1321022113233032-3312033230223121-0110033031203311-2303230110030212-0021320330300332"></a>

### All schema paths for `xcsh_discovery`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--discovery--reference--group-001.md#canonical-0013323033213331-3000331103203323-1201121220302100-2231011130132230-3021023133021202-3230333200310103-1231101202003323-0000133331130220) |
| `cluster_id` | [cluster_id](data-sources--discovery--reference--group-001.md#canonical-1203102113122022-1221113201311023-0103311123213210-0113302110230113-2222111313120011-3020032120031223-2322233020320211-2210312120011333) |
| `description` | [description](data-sources--discovery--reference--group-001.md#canonical-0131301302223120-0212113100311233-2030310012032010-0120213012033000-0030221210031333-1220133033123033-2321031031130012-1210320103231111) |
| `discovery_consul` | [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-3001312111101213-0003231133132302-1112300001113330-1330222302331020-0111113012203000-1321320100020022-2302032122321031-0013211130232121) |
| `discovery_consul.access_info` | [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-2322010231310123-3222313320032322-0330332333213132-2302213300013322-2321323001303233-2322200223010032-3011000012023030-2032302210210031) |
| `discovery_consul.access_info.connection_info` | [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-1110111233011322-3323003102112222-1313110031332122-2332302031301120-3102323020320133-0212100311301112-0001100111200220-3323112032323131) |
| `discovery_consul.access_info.connection_info.api_server` | [discovery_consul.access_info.connection_info.api_server](data-sources--discovery--reference--group-001.md#canonical-1221330030033200-0233330032002022-0032021212303012-1121100103021310-2130310200103013-0020212123111103-2103233133201023-3332132231102000) |
| `discovery_consul.access_info.connection_info.tls_info` | [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-1100323112301233-2311002120213011-0020222213122302-3331221013223211-0103220202120102-2102133010230302-3320101202202130-0220113330033130) |
| `discovery_consul.access_info.connection_info.tls_info.certificate` | [discovery_consul.access_info.connection_info.tls_info.certificate](data-sources--discovery--reference--group-001.md#canonical-3232202113312130-1033102000323102-2310023231033032-3221321330123210-2011231300202023-3331222100010001-0000130102013113-3220132121013233) |
| `discovery_consul.access_info.connection_info.tls_info.key_url` | [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-3011321211303220-1312210030012021-2031311311310201-1323003221113321-1322320222311100-1212223020303232-2100030212020223-0123101102221332) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-2221010001101132-0200110131101311-0332132121011001-3113003120113323-0101310102321212-3102230220001101-3122210123313122-0301311112220332) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](data-sources--discovery--reference--group-001.md#canonical-3330330022133001-1022112020123232-2011030320330101-1330022222332301-3302211032200130-2200020321231031-1320322311310313-3101200113013010) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](data-sources--discovery--reference--group-001.md#canonical-0013003330221122-1001133120010031-0322323323030113-3313020000301310-1231203020333330-1310021202220303-0032310331233222-3320111002122001) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](data-sources--discovery--reference--group-001.md#canonical-2010321302211232-0000120132332111-2222003311230202-1201311232000030-1121332230333233-0323102033011023-3033031213312300-0033132222232321) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-0201122010323203-0233132110301010-2203000231210201-1030102113103332-3233200323323300-3012110300231100-3300312330321300-3321203021223233) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](data-sources--discovery--reference--group-001.md#canonical-1003123330313020-1100002123322131-1222013222003201-0101313330030211-2032112310123331-1120300330221311-0323231003023303-1121211121331231) |
| `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url](data-sources--discovery--reference--group-001.md#canonical-1203331213020130-3020003123031100-1312210300220331-2301233222220303-3010221211331010-0222010011203300-0202220203031202-1020113003221231) |
| `discovery_consul.access_info.connection_info.tls_info.server_name` | [discovery_consul.access_info.connection_info.tls_info.server_name](data-sources--discovery--reference--group-001.md#canonical-3000012202232023-1210333101133020-2030113002123311-0133311122021102-3311320322320122-2222302333132320-2311210010033300-2123233231331210) |
| `discovery_consul.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_consul.access_info.connection_info.tls_info.trusted_ca_url](data-sources--discovery--reference--group-001.md#canonical-0313323111213031-1011121320331331-2101013033112211-2200212100212211-0013023313303020-2312103323210110-1131213020221030-2202321000230311) |
| `discovery_consul.access_info.http_basic_auth_info` | [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-3122330021013033-1333010330330200-1331300331003023-3101012112220310-0313333313103120-0003220303232202-3321121321100130-3122313031113222) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--reference--group-001.md#canonical-1011311020020221-0303031033101333-0302132233011121-3231012130301121-1231013212012001-2322323003031130-1221121130012013-3103302100011131) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-2013003223103133-1211132230013230-2101221102031233-0210030001203033-0120103022210212-3133130322310023-0033023202300212-0203312002302200) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider](data-sources--discovery--reference--group-001.md#canonical-3022301111032120-1112300330211213-3232100322033030-0331131001211100-1103101101321310-2230322210102201-3131330102120110-3220213021310021) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location](data-sources--discovery--reference--group-001.md#canonical-2323033201120000-1201211323102323-3321022123312132-0330203220112202-3032101210123033-3030220103100103-3022013320100013-3222213221320231) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider](data-sources--discovery--reference--group-001.md#canonical-0002001232230203-2203012010022120-3233302122011132-1203021223233003-0312332322312321-1300002203031022-0011233013212201-0323300321121012) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-2011330032030233-1203302132212301-0320232300001233-3123311123321031-1231203120332231-2101122013032210-3030030212201031-0100202132213120) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref](data-sources--discovery--reference--group-001.md#canonical-2222131021001121-3031110131103202-0210120202131332-0233131322111032-3100123122013132-3232122121301132-3001011020303133-0110030230200331) |
| `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url` | [discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url](data-sources--discovery--reference--group-001.md#canonical-2101120232112123-1012020113131311-0003222033101102-1023122212111202-0003121022323010-1223121201030211-2123211311310231-3122203113303211) |
| `discovery_consul.access_info.http_basic_auth_info.user_name` | [discovery_consul.access_info.http_basic_auth_info.user_name](data-sources--discovery--reference--group-001.md#canonical-1202003001113213-2031333010013220-0030103113121100-3333231020203121-2223200100010013-2123132222131030-3230131133121113-0313102100312111) |
| `discovery_consul.publish_info` | [discovery_consul.publish_info](data-sources--discovery--reference--group-001.md#canonical-1023220111010000-2331130322103001-3232203232130313-3301212122231232-1313120011111030-3332212230202111-1300033001003323-3203220223003130) |
| `discovery_consul.publish_info.disable_spec` | [discovery_consul.publish_info.disable_spec](data-sources--discovery--reference--group-001.md#canonical-2123231033220330-0121032021013323-2011103113212303-0023011132023302-0333122000011113-0122232120321103-2322102312023112-3030230021133232) |
| `discovery_consul.publish_info.publish` | [discovery_consul.publish_info.publish](data-sources--discovery--reference--group-001.md#canonical-0323200121213000-2211312022323131-3132003131320022-1030233132200300-2311200110001201-3201103012131221-2111323231110110-0020322022301300) |
| `discovery_k8s` | [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-1032223303001200-3221200212200312-3012010132222000-3221322120331120-2101103030323020-0131333210310121-1100011210320221-1332332232203011) |
| `discovery_k8s.access_info` | [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-2210001111220033-1203132203003013-1301022020230211-1213221022123221-2023232202321313-2311332020022033-0320230222220002-2203033232001220) |
| `discovery_k8s.access_info.connection_info` | [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-2311312333331122-3012011301013110-2021201101013222-0232112103313333-2002310013300102-0223322323331132-2111133203010322-0312333310130103) |
| `discovery_k8s.access_info.connection_info.api_server` | [discovery_k8s.access_info.connection_info.api_server](data-sources--discovery--reference--group-001.md#canonical-0333023303123211-2030310032000233-2300232013233312-0212123313113010-2331220220131330-0301102302000233-1101312232330202-3002010031033221) |
| `discovery_k8s.access_info.connection_info.tls_info` | [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-1122312121320222-3322230133213331-2323302103011202-0130002331022221-1001210002112012-1321103011203012-3010203123311013-3201333020123010) |
| `discovery_k8s.access_info.connection_info.tls_info.certificate` | [discovery_k8s.access_info.connection_info.tls_info.certificate](data-sources--discovery--reference--group-001.md#canonical-0311030232100301-0021010122123203-2112002323312110-2023322301033223-1033231113111123-0130013220011121-2132112213203111-0210012230300030) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url` | [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-3121211210323100-3012100102121100-2231113312013212-2301331112223331-2311013131323322-2330313030101202-0322131020201230-0000121102223223) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-1133311221313203-2311020131130030-1221032221120312-2211302113032312-1313111113132232-2303120013202011-0133112322102301-3230311223201300) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider](data-sources--discovery--reference--group-001.md#canonical-1231023222111200-0312032030011311-0011320231101020-1120321122321310-3010200111231112-2121121332122302-0312102231112311-1322300230231031) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location](data-sources--discovery--reference--group-001.md#canonical-2023020333101020-3030100213311321-2223122333323031-1323021302302302-1303311003320312-1200033011022000-3311212230221110-1023031302021001) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider](data-sources--discovery--reference--group-001.md#canonical-3230322302310200-0003130320020130-0103312022222122-0010330113033110-2013131133132203-2210302300121132-2210221230323133-0031233231021101) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-0111221031200123-0222001023122123-1222210020103111-3323001201322030-2031310233311020-2323211023331103-2021223302033311-1210223320301122) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref](data-sources--discovery--reference--group-002.md#canonical-0330002231230302-2230033132220031-2110100212122002-1101311113000202-0320230000202203-0223121030331021-1113111232210001-1103002000022112) |
| `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url` | [discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info.url](data-sources--discovery--reference--group-002.md#canonical-3012322212212000-3333101232331313-3122110322101001-2322320130131232-1201302232220120-0313101031220020-2203020030123230-3312302302000100) |
| `discovery_k8s.access_info.connection_info.tls_info.server_name` | [discovery_k8s.access_info.connection_info.tls_info.server_name](data-sources--discovery--reference--group-001.md#canonical-0222100210010033-0330313002131111-2221103231021112-3332313322003103-3213031101202133-0013032032300322-1302321230122023-2120120313221033) |
| `discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url` | [discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url](data-sources--discovery--reference--group-001.md#canonical-1012210311133101-0112000113320111-2110213312023302-3302010113303023-0232330000122133-3013200203111230-1113201113033002-3200112233120311) |
| `discovery_k8s.access_info.isolated` | [discovery_k8s.access_info.isolated](data-sources--discovery--reference--group-002.md#canonical-3211013220013210-2330200001321322-1330221213221000-1221230133203311-1101010320302122-2111222033110231-0002130112322222-3101302221111111) |
| `discovery_k8s.access_info.kubeconfig_url` | [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--reference--group-002.md#canonical-1210333310133030-3021230123230212-3320112222212210-2300223123300123-1320101120020203-0323102033302210-2233301213121203-3132023020221323) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](data-sources--discovery--reference--group-002.md#canonical-2103322020011101-3313203020002120-3201221120300203-3122112302300230-2200211003322202-0233030112010113-1201103110030311-0202310123033202) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.decryption_provider](data-sources--discovery--reference--group-002.md#canonical-2203221012022212-1202130331031000-0021222100332331-2003231033230210-3213212330322012-0123020231101130-3013221010003102-1033130102121221) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.location](data-sources--discovery--reference--group-002.md#canonical-0010123321311301-0323231023212031-1302202131310310-1332303000122013-1033111003102001-1222201210323112-2010301033001200-1203010232113121) |
| `discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider` | [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info.store_provider](data-sources--discovery--reference--group-002.md#canonical-0112123011032013-2132213100130212-1311203132233320-3313201030022113-2220110323221123-1102203120301210-0111121311203023-1210123101123211) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](data-sources--discovery--reference--group-002.md#canonical-0231333002203201-2101132012300331-2010311023131333-3120021323222211-2103101131202033-3222012113302022-1012323222310031-1200133200131213) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.provider_ref](data-sources--discovery--reference--group-002.md#canonical-0032001230230323-3112321020310322-1022010310002021-3323121022001021-3211321100300010-2210201010311030-3012213003132222-3300300120202121) |
| `discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url` | [discovery_k8s.access_info.kubeconfig_url.clear_secret_info.url](data-sources--discovery--reference--group-002.md#canonical-1300231300130022-3233030222011120-2331321013211122-1022322023002030-0131020330030111-1022132310222111-2311003233001132-0300103032200323) |
| `discovery_k8s.access_info.reachable` | [discovery_k8s.access_info.reachable](data-sources--discovery--reference--group-002.md#canonical-3130223331201311-2313022032233003-1121212101330122-0033110121001230-2122133332111102-0102323321310103-0010133223330300-2121103123311021) |
| `discovery_k8s.default_all` | [discovery_k8s.default_all](data-sources--discovery--reference--group-002.md#canonical-3210321213011110-2311302010321030-2122310222321021-0033023130232021-3001203030020011-3203122223013103-2233212222031122-3023033123221030) |
| `discovery_k8s.namespace_mapping` | [discovery_k8s.namespace_mapping](data-sources--discovery--reference--group-002.md#canonical-2001322221023102-2130230202311020-3233201020110333-3311030322102203-3133111323220001-0000023002021113-3120203211322012-3230020231112012) |
| `discovery_k8s.namespace_mapping.items` | [discovery_k8s.namespace_mapping.items](data-sources--discovery--reference--group-002.md#canonical-2132211121203312-1232000332123322-2321011010003113-0223211010221100-1132322000322210-3023223021101123-0021010123333033-0202131111213100) |
| `discovery_k8s.namespace_mapping.items.namespace` | [discovery_k8s.namespace_mapping.items.namespace](data-sources--discovery--reference--group-002.md#canonical-0320032001030203-2303232300302021-3210203012111012-2303230001210303-2110230230213102-3333120022312133-2012312103100202-3223222020012232) |
| `discovery_k8s.namespace_mapping.items.namespace_regex` | [discovery_k8s.namespace_mapping.items.namespace_regex](data-sources--discovery--reference--group-002.md#canonical-0222322220213002-0132130023220100-3202202323203131-3200201001102201-1302012200331001-2122301300332220-3130230231201313-2031032130122311) |
| `discovery_k8s.publish_info` | [discovery_k8s.publish_info](data-sources--discovery--reference--group-002.md#canonical-0121022203221333-2232012212222101-0221331211133021-2301301002013333-0211010013100101-2022020131010031-2220113021102211-0330022332012122) |
| `discovery_k8s.publish_info.disable_spec` | [discovery_k8s.publish_info.disable_spec](data-sources--discovery--reference--group-002.md#canonical-1102232033130113-1120111201131032-3320202232120321-3213020130101131-0203300103110113-1222013203312221-3132212303033223-1033323111230300) |
| `discovery_k8s.publish_info.dns_delegation` | [discovery_k8s.publish_info.dns_delegation](data-sources--discovery--reference--group-002.md#canonical-1021331020323302-2333002120113112-2002133202131312-0201020103322210-2120122033320121-2002102100300302-2222012103230032-1303203032022232) |
| `discovery_k8s.publish_info.dns_delegation.dns_mode` | [discovery_k8s.publish_info.dns_delegation.dns_mode](data-sources--discovery--reference--group-002.md#canonical-0231112021101120-0033030231033102-1232203020121220-3223313222011223-2202200003021312-2023322011030202-0001100311122320-0220021103230121) |
| `discovery_k8s.publish_info.dns_delegation.subdomain` | [discovery_k8s.publish_info.dns_delegation.subdomain](data-sources--discovery--reference--group-002.md#canonical-0130121213233301-3133321120322023-0223202200133303-1003031123333213-0031122030232322-2120021012122120-2111230031013132-3230310213131003) |
| `discovery_k8s.publish_info.publish` | [discovery_k8s.publish_info.publish](data-sources--discovery--reference--group-002.md#canonical-0202130231222233-1321100212331103-3010120031323110-1230321133033033-3130320232310221-1312332213131320-3212303032020003-3212331001323000) |
| `discovery_k8s.publish_info.publish.namespace` | [discovery_k8s.publish_info.publish.namespace](data-sources--discovery--reference--group-002.md#canonical-3223312102132300-3103022022023301-2020030101033011-1122332231222112-3022201011213010-0323023033222211-3020020103130003-2103321320313113) |
| `discovery_k8s.publish_info.publish_fqdns` | [discovery_k8s.publish_info.publish_fqdns](data-sources--discovery--reference--group-002.md#canonical-0330202321302013-2222011222223213-1230333311230000-2303221022122201-1301000322000302-3200112332332011-1221103232210031-0123011232332202) |
| `id` | [ID](data-sources--discovery--reference--group-001.md#canonical-0333321230233213-0031212000100031-3302111231213110-1022123311301233-3222211330101210-3212333313000311-3311031310213013-2032212220310302) |
| `labels` | [labels](data-sources--discovery--reference--group-001.md#canonical-3223222201123202-2332112000233020-3233202330221211-3101030202222121-3030023203331030-1012001130303032-1220002230013103-0130030020033122) |
| `name` | [name](data-sources--discovery--reference--group-001.md#canonical-3233301022022013-1000030032313010-2121301132231033-1332111121212003-2300212131221013-2230320223130210-3201023313001323-3022323213130032) |
| `namespace` | [namespace](data-sources--discovery--reference--group-001.md#canonical-0302231212213212-3312212232232223-2311232330102233-1332211003203321-0223011201033122-1220022021223333-3032110002302000-0302010130030231) |
| `no_cluster_id` | [no_cluster_id](data-sources--discovery--reference--group-002.md#canonical-2330131222130101-1223121232130000-3310203111002300-3111030300122323-2032022200033233-3120020220301310-3232021301122000-1013013012212302) |
| `where` | [where](data-sources--discovery--reference--group-002.md#canonical-1001023320211113-3130223330020112-0313320233030103-1120010333033202-2102010130130203-2210230232101121-0112310133203302-0100332000302133) |
| `where.site` | [where.site](data-sources--discovery--reference--group-002.md#canonical-2332030123000123-1332002322102101-0303233200130021-2101101032121320-3201023103031301-0202301323132301-0310021233310111-0331123230200211) |
| `where.site.disable_internet_vip` | [where.site.disable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-2030031300113323-0010110223221003-2023031302000332-0202232312331130-3200003300003210-2213312023200100-3102003132012132-0210323030030333) |
| `where.site.enable_internet_vip` | [where.site.enable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-1002020120312333-3331230301213201-0231232130130323-0323320202000020-2231202331021120-3300103120102220-3201223022100020-3113100130213023) |
| `where.site.network_type` | [where.site.network_type](data-sources--discovery--reference--group-002.md#canonical-3011001110022122-1021013002103021-3111220011103030-3133323223130323-0001321000033021-1211300123300233-1323313130122100-0020221132001121) |
| `where.site.ref` | [where.site.ref](data-sources--discovery--reference--group-002.md#canonical-0013011012032002-2112121110233203-1122200300130232-2012333230310231-1320333102200021-0322033323202123-2103333231012031-0330103020110233) |
| `where.site.ref.kind` | [where.site.ref.kind](data-sources--discovery--reference--group-002.md#canonical-0130011331102230-3220113112221202-2333300113030221-1020201321030101-3102312302100330-0300330110222101-3302220111031132-0203213013203210) |
| `where.site.ref.name` | [where.site.ref.name](data-sources--discovery--reference--group-002.md#canonical-2232020130130030-0030222113021212-3231232032023230-0303213221011011-2301001333130122-0323313302132213-3201001112203333-3131110003103033) |
| `where.site.ref.namespace` | [where.site.ref.namespace](data-sources--discovery--reference--group-002.md#canonical-2031213121232012-3222123110002322-0211200101231210-1232110301112322-1103320011211231-2331210032122001-1022202030223313-0031320030332202) |
| `where.site.ref.tenant` | [where.site.ref.tenant](data-sources--discovery--reference--group-002.md#canonical-0110012123223223-1023210333000311-0210323110232032-1311130202213332-1310320022122221-3112213323230121-0111233233000102-3022221031030330) |
| `where.site.ref.uid` | [where.site.ref.uid](data-sources--discovery--reference--group-002.md#canonical-3312030213330313-1332031321031101-2323013000300013-3323200111023202-2323012322230033-0012222230233033-3202202220000211-2202120113211123) |
| `where.virtual_network` | [where.virtual_network](data-sources--discovery--reference--group-002.md#canonical-0223030102121002-2123230021330200-3103200211323201-0233313121200102-3030330120233330-2303002133322202-1203330011023332-3221332231311202) |
| `where.virtual_network.ref` | [where.virtual_network.ref](data-sources--discovery--reference--group-002.md#canonical-1311001320310131-0100302300002102-2320032213230102-2100003002213030-0323131120301301-1031233201321211-2122331331032011-1210210212110213) |
| `where.virtual_network.ref.kind` | [where.virtual_network.ref.kind](data-sources--discovery--reference--group-002.md#canonical-2112002132212001-2003233120203311-3032321100120121-3220031312213212-0121202212111221-2313211032311021-2012201330312103-2211232201111111) |
| `where.virtual_network.ref.name` | [where.virtual_network.ref.name](data-sources--discovery--reference--group-002.md#canonical-0102103211220203-0301333001112111-1032012320200032-1202222312310012-1311322120002311-2123222120321221-3021303023302110-3210123131013230) |
| `where.virtual_network.ref.namespace` | [where.virtual_network.ref.namespace](data-sources--discovery--reference--group-002.md#canonical-2121022210310020-2010032232022003-3003333101133323-1002203312000231-3232312022221021-1020333301101231-3110112210120212-0003210012112232) |
| `where.virtual_network.ref.tenant` | [where.virtual_network.ref.tenant](data-sources--discovery--reference--group-002.md#canonical-0122032130001322-1231202123102221-3121021003121013-1010113003330231-0231211122130013-1223130121002202-2201101230210001-3201023333311010) |
| `where.virtual_network.ref.uid` | [where.virtual_network.ref.uid](data-sources--discovery--reference--group-002.md#canonical-2303322321221301-3110012132223100-1333011120031230-1031222221000231-3212331333012313-0331033003122111-1200313121011202-0232212310312122) |
| `where.virtual_site` | [where.virtual_site](data-sources--discovery--reference--group-002.md#canonical-3121122332022100-2322022213221212-0030221202020033-0013010231331211-3232021320122012-3112321021330201-0311211221311133-0333021120321313) |
| `where.virtual_site.disable_internet_vip` | [where.virtual_site.disable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-3121010010312310-2012012330223202-0133303311112213-1300211123033101-3112301131231222-1200110102021221-3123330030130012-0323213030231012) |
| `where.virtual_site.enable_internet_vip` | [where.virtual_site.enable_internet_vip](data-sources--discovery--reference--group-002.md#canonical-3011122123133320-1303211013132313-1112123011003222-0003232113101030-3110122102333102-1322203212322021-1021003202331110-3013333001212003) |
| `where.virtual_site.network_type` | [where.virtual_site.network_type](data-sources--discovery--reference--group-002.md#canonical-3320020330211301-1102122020300322-2013331030233300-3113002012313023-0313002001233231-3100203330121111-1331013312002022-0000232011033021) |
| `where.virtual_site.ref` | [where.virtual_site.ref](data-sources--discovery--reference--group-002.md#canonical-0003002230031332-3213000322300231-0312212202233210-1002020331313320-2002103231103312-2201201210233320-0032021202013310-3001311232330013) |
| `where.virtual_site.ref.kind` | [where.virtual_site.ref.kind](data-sources--discovery--reference--group-002.md#canonical-1230000133212313-1302201130321013-0300211302313300-3103213201302210-2101020230310013-0223031133213113-3320202002211122-3131330010330301) |
| `where.virtual_site.ref.name` | [where.virtual_site.ref.name](data-sources--discovery--reference--group-002.md#canonical-2121102110133313-3103303223012200-3012220100312133-1311300020001322-3113312213223103-1301113333120331-3303113330100031-1311031232233200) |
| `where.virtual_site.ref.namespace` | [where.virtual_site.ref.namespace](data-sources--discovery--reference--group-002.md#canonical-0312223301202212-2200221213013001-0230121133112010-0313300000303022-1332232122113310-1102020233010220-2111030231100022-2211212130213320) |
| `where.virtual_site.ref.tenant` | [where.virtual_site.ref.tenant](data-sources--discovery--reference--group-002.md#canonical-2320113201012322-3011103012123211-1121033302223312-0032222302222130-3012132013310330-0020323320333332-2010230311000323-2303301212020011) |
| `where.virtual_site.ref.uid` | [where.virtual_site.ref.uid](data-sources--discovery--reference--group-002.md#canonical-3212130122302310-0120202023312212-3301200101013110-3102303122203200-3003002201111120-2001020013020220-3032223301302123-3101011122011122) |

<a id="canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- discovery_consul

<a id="canonical-3001312111101213-0003231133132302-1112300001113330-1330222302331020-0111113012203000-1321320100020022-2302032122321031-0013211130232121"></a>

Type: `"single"`. Computed.

\[OneOf: discovery\_consul, discovery\_k8s\] Discovery configuration for Hashicorp Consul.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-namespace_mapping_choice": "[]"
}
```

OneOf alternatives in this subsection:

- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-3001312111101213-0003231133132302-1112300001113330-1330222302331020-0111113012203000-1321320100020022-2302032122321031-0013211130232121)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-1032223303001200-3221200212200312-3012010132222000-3221322120331120-2101103030323020-0131333210310121-1100011210320221-1332332232203011)

Select alternatives according to the provider validators above.

<a id="canonical-2200021033021111-3211201022203333-2210022303300000-0101322003123013-1321031311031331-2122122200031202-1111301112012001-2102310320300330"></a>

### Direct properties for `discovery_consul`

- [access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201): complete subsection reference.

- [publish_info](data-sources--discovery--reference--group-001.md#canonical-3022323202130331-1300101212130033-0001310133112233-0230110021230003-0022213333130131-1230010030020111-1232120213113231-3303010023321313): complete subsection reference.

<a id="canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- discovery_consul.access_info

<a id="canonical-2322010231310123-3222313320032322-0330332333213132-2302213300013322-2321323001303233-2322200223010032-3011000012023030-2032302210210031"></a>

Type: `"single"`. Computed.

Hashicorp Consul API server information.

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

<a id="canonical-1321130030001110-1210013200311211-3202013110013200-0000220211322012-2200021003022320-1233331030030123-2233210220030321-1302133110223020"></a>

### Direct properties for `discovery_consul.access_info`

- [connection_info](data-sources--discovery--reference--group-001.md#canonical-1323022120103201-2111202220210033-2122100210313202-0213233220010233-1213002110233311-2211111031320213-0202310300101112-0002103011111002): complete subsection reference.

- [http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-3200201013223011-3322202301233100-1122211121001023-1322201313323131-0320133303212223-1112023300300321-3032030112132300-3123232213232003): complete subsection reference.

<a id="canonical-1323022120103201-2111202220210033-2122100210313202-0213233220010233-1213002110233311-2211111031320213-0202310300101112-0002103011111002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info.connection_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201)
- discovery_consul.access_info.connection_info

<a id="canonical-1110111233011322-3323003102112222-1313110031332122-2332302031301120-3102323020320133-0212100311301112-0001100111200220-3323112032323131"></a>

Type: `"single"`. Computed.

Configuration details to access discovery service REST API.

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

<a id="canonical-0220310223300323-2001323013133303-0322200311010000-2111313330313003-2223213233321312-1310223233203220-1100211100333130-2010321122001212"></a>

### Direct properties for `discovery_consul.access_info.connection_info`

<a id="canonical-1221330030033200-0233330032002022-0032021212303012-1121100103021310-2130310200103013-0020212123111103-2103233133201023-3332132231102000"></a>

#### `discovery_consul.access_info.connection_info.api_server` property

Type: `"string"`. Computed.

API server must be a fully qualified domain string and port specified as host:port pair.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

- [tls_info](data-sources--discovery--reference--group-001.md#canonical-3120002031203312-1021130102021120-2230313003313301-1301003121003302-1230021032321200-0021021022120010-2301322223032012-3120032312133220): complete subsection reference.

<a id="canonical-3120002031203312-1021130102021120-2230313003313301-1301003121003302-1230021032321200-0021021022120010-2301322223032012-3120032312133220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info.connection_info.tls_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201)
- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-1323022120103201-2111202220210033-2122100210313202-0213233220010233-1213002110233311-2211111031320213-0202310300101112-0002103011111002)
- discovery_consul.access_info.connection_info.tls_info

<a id="canonical-1100323112301233-2311002120213011-0020222213122302-3331221013223211-0103220202120102-2102133010230302-3320101202202130-0220113330033130"></a>

Type: `"single"`. Computed.

TLS config for client of discovery service.

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

<a id="canonical-2113312013103130-2021202030110031-3323011131132201-2003313111020303-2133010100232222-2020013312011301-1313120233011031-2332303212322031"></a>

### Direct properties for `discovery_consul.access_info.connection_info.tls_info`

<a id="canonical-3232202113312130-1033102000323102-2310023231033032-3221321330123210-2011231300202023-3331222100010001-0000130102013113-3220132121013233"></a>

#### `discovery_consul.access_info.connection_info.tls_info.certificate` property

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--discovery--reference--group-001.md#canonical-1311010103300011-1033223221002030-1131111200111330-0303313220230313-2103111303303110-2120230111131022-0331011110230110-1133003011121221): complete subsection reference.

<a id="canonical-3000012202232023-1210333101133020-2030113002123311-0133311122021102-3311320322320122-2222302333132320-2311210010033300-2123233231331210"></a>

<a id="canonical-2113003013200333-3012330222032221-3333130320331312-2130002121131221-2313200131203323-0311230202310231-1131302330113003-2312221120112221"></a>

#### `discovery_consul.access_info.connection_info.tls_info.server_name` property

Type: `"string"`. Computed.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0313323111213031-1011121320331331-2101013033112211-2200212100212211-0013023313303020-2312103323210110-1131213020221030-2202321000230311"></a>

<a id="canonical-2332212130030022-0012232233332120-0012213230330111-1101130000300000-2111132000121122-1020230310222312-2212010101032333-0303013303221330"></a>

#### `discovery_consul.access_info.connection_info.tls_info.trusted_ca_url` property

Type: `"string"`. Computed.

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1311010103300011-1033223221002030-1131111200111330-0303313220230313-2103111303303110-2120230111131022-0331011110230110-1133003011121221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info.connection_info.tls_info.key_url` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201)
- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-1323022120103201-2111202220210033-2122100210313202-0213233220010233-1213002110233311-2211111031320213-0202310300101112-0002103011111002)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-3120002031203312-1021130102021120-2230313003313301-1301003121003302-1230021032321200-0021021022120010-2301322223032012-3120032312133220)
- discovery_consul.access_info.connection_info.tls_info.key_url

<a id="canonical-3011321211303220-1312210030012021-2031311311310201-1323003221113321-1322320222311100-1212223020303232-2100030212020223-0123101102221332"></a>

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

<a id="canonical-3032133130032033-1232331103113131-0113303100233303-3223222202012302-2021221000032321-0133133111001222-0003023112122310-3310112003332120"></a>

### Direct properties for `discovery_consul.access_info.connection_info.tls_info.key_url`

- [blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-3211020033133213-1112011003220300-1320020003330033-2330101031022210-1302320012313301-1312331031220111-1012112121333000-1313302101210330): complete subsection reference.

- [clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-3233010311321003-1223102001030323-3233002333110102-0100321321111101-1322122302100001-0211211113111023-2001120000011100-1031010323102123): complete subsection reference.

<a id="canonical-3211020033133213-1112011003220300-1320020003330033-2330101031022210-1302320012313301-1312331031220111-1012112121333000-1313302101210330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201)
- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-1323022120103201-2111202220210033-2122100210313202-0213233220010233-1213002110233311-2211111031320213-0202310300101112-0002103011111002)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-3120002031203312-1021130102021120-2230313003313301-1301003121003302-1230021032321200-0021021022120010-2301322223032012-3120032312133220)
- [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-1311010103300011-1033223221002030-1131111200111330-0303313220230313-2103111303303110-2120230111131022-0331011110230110-1133003011121221)
- discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info

<a id="canonical-2221010001101132-0200110131101311-0332132121011001-3113003120113323-0101310102321212-3102230220001101-3122210123313122-0301311112220332"></a>

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

<a id="canonical-3123002223220331-1112301301232000-3330121103303330-3112230201011003-1221232200001123-0120013221133102-0302030011212233-3332113023022121"></a>

### Direct properties for `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info`

<a id="canonical-3330330022133001-1022112020123232-2011030320330101-1330022222332301-3302211032200130-2200020321231031-1320322311310313-3101200113013010"></a>

#### `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0013003330221122-1001133120010031-0322323323030113-3313020000301310-1231203020333330-1310021202220303-0032310331233222-3320111002122001"></a>

<a id="canonical-2031032221002303-2202022101132112-1133300100300200-3223103231310011-2300213132113013-0013102113130033-0122132112010203-2310313111103032"></a>

#### `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2010321302211232-0000120132332111-2222003311230202-1201311232000030-1121332230333233-0323102033011023-3033031213312300-0033132222232321"></a>

<a id="canonical-0302331101000210-2113210333023022-3101221031310330-3013323202200200-0313132011030023-0103333211112101-2322333330221032-1331202111022011"></a>

#### `discovery_consul.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3233010311321003-1223102001030323-3233002333110102-0100321321111101-1322122302100001-0211211113111023-2001120000011100-1031010323102123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201)
- [discovery_consul.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-1323022120103201-2111202220210033-2122100210313202-0213233220010233-1213002110233311-2211111031320213-0202310300101112-0002103011111002)
- [discovery_consul.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-3120002031203312-1021130102021120-2230313003313301-1301003121003302-1230021032321200-0021021022120010-2301322223032012-3120032312133220)
- [discovery_consul.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-1311010103300011-1033223221002030-1131111200111330-0303313220230313-2103111303303110-2120230111131022-0331011110230110-1133003011121221)
- discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info

<a id="canonical-0201122010323203-0233132110301010-2203000231210201-1030102113103332-3233200323323300-3012110300231100-3300312330321300-3321203021223233"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-1102123312102101-3122210003211231-1000021322330332-3022210201331030-1332313013223322-0110012031020323-1331221031311010-1122020230313233"></a>

### Direct properties for `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info`

<a id="canonical-1003123330313020-1100002123322131-1222013222003201-0101313330030211-2032112310123331-1120300330221311-0323231003023303-1121211121331231"></a>

#### `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1203331213020130-3020003123031100-1312210300220331-2301233222220303-3010221211331010-0222010011203300-0202220203031202-1020113003221231"></a>

<a id="canonical-1121102113111301-1211031211030333-3302212220122202-0122031110312323-2110003222112132-0310023200103321-1300113331003130-1003332220131121"></a>

#### `discovery_consul.access_info.connection_info.tls_info.key_url.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3200201013223011-3322202301233100-1122211121001023-1322201313323131-0320133303212223-1112023300300321-3032030112132300-3123232213232003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info.http_basic_auth_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201)
- discovery_consul.access_info.http_basic_auth_info

<a id="canonical-3122330021013033-1333010330330200-1331300331003023-3101012112220310-0313333313103120-0003220303232202-3321121321100130-3122313031113222"></a>

Type: `"single"`. Computed.

Authentication parameters to access Hashicorp Consul.

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

<a id="canonical-3031223332103002-2213010223010221-3313302200121301-1021321111123320-0203203101113033-2113221133321201-1223100310210312-3031323113022001"></a>

### Direct properties for `discovery_consul.access_info.http_basic_auth_info`

- [passwd_url](data-sources--discovery--reference--group-001.md#canonical-0031333100300232-0132012322233013-3130121300011033-0221122132331011-1213321332330301-0233021211233130-2121032011130212-0003212313003230): complete subsection reference.

<a id="canonical-1202003001113213-2031333010013220-0030103113121100-3333231020203121-2223200100010013-2123132222131030-3230131133121113-0313102100312111"></a>

<a id="canonical-3313303122303221-2033021331023211-3120003202330112-2010013213202320-2320310233320020-3231022311101123-2102301122230221-2331113233120032"></a>

#### `discovery_consul.access_info.http_basic_auth_info.user_name` property

Type: `"string"`. Computed.

username. Username in consul.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-0031333100300232-0132012322233013-3130121300011033-0221122132331011-1213321332330301-0233021211233130-2121032011130212-0003212313003230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info.http_basic_auth_info.passwd_url` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-3200201013223011-3322202301233100-1122211121001023-1322201313323131-0320133303212223-1112023300300321-3032030112132300-3123232213232003)
- discovery_consul.access_info.http_basic_auth_info.passwd_url

<a id="canonical-1011311020020221-0303031033101333-0302132233011121-3231012130301121-1231013212012001-2322323003031130-1221121130012013-3103302100011131"></a>

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

<a id="canonical-2131120232003022-1310002302013122-3213223321320023-3013200110013030-2202010032102120-3221211332331221-1001023312012212-0221031313021001"></a>

### Direct properties for `discovery_consul.access_info.http_basic_auth_info.passwd_url`

- [blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-2223132131113321-1203220323323223-3202310011203021-3321301110221221-1323022212012312-2122103020003102-1000212113010100-0332211130002221): complete subsection reference.

- [clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-3322000300130232-2030020220200321-1112230233023223-0021233203100302-0320313212212022-0033312022102210-0213303300320202-1111021311100310): complete subsection reference.

<a id="canonical-2223132131113321-1203220323323223-3202310011203021-3321301110221221-1323022212012312-2122103020003102-1000212113010100-0332211130002221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-3200201013223011-3322202301233100-1122211121001023-1322201313323131-0320133303212223-1112023300300321-3032030112132300-3123232213232003)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--reference--group-001.md#canonical-0031333100300232-0132012322233013-3130121300011033-0221122132331011-1213321332330301-0233021211233130-2121032011130212-0003212313003230)
- discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info

<a id="canonical-2013003223103133-1211132230013230-2101221102031233-0210030001203033-0120103022210212-3133130322310023-0033023202300212-0203312002302200"></a>

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

<a id="canonical-3312301030220123-3003330030112121-2321023022233023-0030321312131300-3312002313200302-3332030030201313-1231130221032301-2002302201103202"></a>

### Direct properties for `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info`

<a id="canonical-3022301111032120-1112300330211213-3232100322033030-0331131001211100-1103101101321310-2230322210102201-3131330102120110-3220213021310021"></a>

#### `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2323033201120000-1201211323102323-3321022123312132-0330203220112202-3032101210123033-3030220103100103-3022013320100013-3222213221320231"></a>

<a id="canonical-2110130212011032-1301223323033330-3030231132032120-1232210230132021-1031320231100312-3231003010203333-0332012113011020-0131131001230131"></a>

#### `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0002001232230203-2203012010022120-3233302122011132-1203021223233003-0312332322312321-1300002203031022-0011233013212201-0323300321121012"></a>

<a id="canonical-3123233130003022-2222113222123211-1102100021010120-0030021321003232-3031011111311230-1011232122231011-3330321221221012-0032021201332012"></a>

#### `discovery_consul.access_info.http_basic_auth_info.passwd_url.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3322000300130232-2030020220200321-1112230233023223-0021233203100302-0320313212212022-0033312022102210-0213303300320202-1111021311100310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.access_info](data-sources--discovery--reference--group-001.md#canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201)
- [discovery_consul.access_info.http_basic_auth_info](data-sources--discovery--reference--group-001.md#canonical-3200201013223011-3322202301233100-1122211121001023-1322201313323131-0320133303212223-1112023300300321-3032030112132300-3123232213232003)
- [discovery_consul.access_info.http_basic_auth_info.passwd_url](data-sources--discovery--reference--group-001.md#canonical-0031333100300232-0132012322233013-3130121300011033-0221122132331011-1213321332330301-0233021211233130-2121032011130212-0003212313003230)
- discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info

<a id="canonical-2011330032030233-1203302132212301-0320232300001233-3123311123321031-1231203120332231-2101122013032210-3030030212201031-0100202132213120"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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

<a id="canonical-3311232120203111-2112021103312122-2332231302113001-3331020220321312-2201010132202230-0332320113222012-0303332033112232-0302210330231100"></a>

### Direct properties for `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info`

<a id="canonical-2222131021001121-3031110131103202-0210120202131332-0233131322111032-3100123122013132-3232122121301132-3001011020303133-0110030230200331"></a>

#### `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-2101120232112123-1012020113131311-0003222033101102-1023122212111202-0003121022323010-1223121201030211-2123211311310231-3122203113303211"></a>

<a id="canonical-2133332202220003-2202131022231213-0300222101333013-2003132220023021-0222101333133013-2320022121000011-0300320000121023-0300330102021230"></a>

#### `discovery_consul.access_info.http_basic_auth_info.passwd_url.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3022323202130331-1300101212130033-0001310133112233-0230110021230003-0022213333130131-1230010030020111-1232120213113231-3303010023321313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.publish_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- discovery_consul.publish_info

<a id="canonical-1023220111010000-2331130322103001-3232203232130313-3301212122231232-1313120011111030-3332212230202111-1300033001003323-3203220223003130"></a>

Type: `"single"`. Computed.

Configuration parameter for publish info.

Additional upstream details:

Consul Configuration to publish VIPs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"publish\"]"
}
```

<a id="canonical-0302023300302313-2222330312020332-2122033113012331-0000201131321331-0303330113030012-0222320123301230-2131132301100232-0103112121010102"></a>

### Direct properties for `discovery_consul.publish_info`

- [disable_spec](data-sources--discovery--reference--group-001.md#canonical-0220131321330311-0113032323213012-3030223102203020-1333320211202020-1123231010212113-1012121202322301-0123310302313300-3023133010031323): complete subsection reference.

- [publish](data-sources--discovery--reference--group-001.md#canonical-3113200132221202-3320211000033310-1000121301210133-1221302023030133-2231033311213332-1230101032020002-0113033011132313-2301130030001101): complete subsection reference.

<a id="canonical-0220131321330311-0113032323213012-3030223102203020-1333320211202020-1123231010212113-1012121202322301-0123310302313300-3023133010031323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.publish_info.disable_spec` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.publish_info](data-sources--discovery--reference--group-001.md#canonical-3022323202130331-1300101212130033-0001310133112233-0230110021230003-0022213333130131-1230010030020111-1232120213113231-3303010023321313)
- discovery_consul.publish_info.disable_spec

<a id="canonical-2123231033220330-0121032021013323-2011103113212303-0023011132023302-0333122000011113-0122232120321103-2322102312023112-3030230021133232"></a>

Type: `["object", {}]`. Computed.

Enable this option

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113200132221202-3320211000033310-1000121301210133-1221302023030133-2231033311213332-1230101032020002-0113033011132313-2301130030001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_consul.publish_info.publish` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_consul](data-sources--discovery--reference--group-001.md#canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110)
- [discovery_consul.publish_info](data-sources--discovery--reference--group-001.md#canonical-3022323202130331-1300101212130033-0001310133112233-0230110021230003-0022213333130131-1230010030020111-1232120213113231-3303010023321313)
- discovery_consul.publish_info.publish

<a id="canonical-0323200121213000-2211312022323131-3132003131320022-1030233132200300-2311200110001201-3201103012131221-2111323231110110-0020322022301300"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- discovery_k8s

<a id="canonical-1032223303001200-3221200212200312-3012010132222000-3221322120331120-2101103030323020-0131333210310121-1100011210320221-1332332232203011"></a>

Type: `"single"`. Computed.

Configuration parameter for discovery k8s.

Additional upstream details:

Discovery configuration for K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-namespace_mapping_choice": "[\"default_all\",\"namespace_mapping\"]"
}
```

<a id="canonical-3232131130002021-3213323133202000-0233020312011223-0122100310213022-2010323120021020-0111310032202233-3120013211312133-1220232031320331"></a>

### Direct properties for `discovery_k8s`

- [access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233): complete subsection reference.

- [default_all](data-sources--discovery--reference--group-002.md#canonical-0111030212233202-0132113112330210-3102303023102000-2311213102011302-2231201111012332-2233001123200131-3010310220333110-2310213013031030): complete subsection reference.

- [namespace_mapping](data-sources--discovery--reference--group-002.md#canonical-3303013301213222-3332332202321221-2311112031001231-0323332330033132-2230203122010120-1131213122133310-1033023232121000-2320222012200021): complete subsection reference.

- [publish_info](data-sources--discovery--reference--group-002.md#canonical-2000131012030001-1310223011223322-3222122000322101-3130330332001002-2203022320301110-1121220212302102-1202210220223323-1101112222203112): complete subsection reference.

<a id="canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- discovery_k8s.access_info

<a id="canonical-2210001111220033-1203132203003013-1301022020230211-1213221022123221-2023232202321313-2311332020022033-0320230222220002-2203033232001220"></a>

Type: `"single"`. Computed.

Configuration parameter for access info.

Additional upstream details:

K8s API server access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-config_type": "[\"connection_info\",\"kubeconfig_url\"]",
  "x-ves-oneof-field-k8s_pod_network_choice": "[\"isolated\",\"reachable\"]"
}
```

<a id="canonical-1011121110300030-0112323320311011-3331131133002110-3033022122320203-0303131310123320-2201300212312330-1223310022031012-1203020322320203"></a>

### Direct properties for `discovery_k8s.access_info`

- [connection_info](data-sources--discovery--reference--group-001.md#canonical-3033110220233312-1331012131312110-1301011002212212-0003032332020020-2223302001311301-2012112010022110-0023111321011330-3311111002111211): complete subsection reference.

- [isolated](data-sources--discovery--reference--group-002.md#canonical-0000221330012012-2302022102003100-1020030310232110-0021302131123120-0223233223020101-3101330233200313-3233000300212213-2100230312210011): complete subsection reference.

- [kubeconfig_url](data-sources--discovery--reference--group-002.md#canonical-1033022012013103-3002223202110122-0023123111320030-3303111222312023-3212132023330113-2310233213012123-2013212323130130-3100102223221311): complete subsection reference.

- [reachable](data-sources--discovery--reference--group-002.md#canonical-1312013231130203-2031112120032320-1112122010012202-1211033002212132-2303102311110103-0201313021012021-0000203103133233-1201230330203320): complete subsection reference.

<a id="canonical-3033110220233312-1331012131312110-1301011002212212-0003032332020020-2223302001311301-2012112010022110-0023111321011330-3311111002111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.connection_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- discovery_k8s.access_info.connection_info

<a id="canonical-2311312333331122-3012011301013110-2021201101013222-0232112103313333-2002310013300102-0223322323331132-2111133203010322-0312333310130103"></a>

Type: `"single"`. Computed.

Configuration details to access discovery service REST API.

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

<a id="canonical-1211332211231121-0220311120211112-1132022322232232-1121202302233032-3332300303223030-2322300210100302-1220130313232110-3332330130002031"></a>

### Direct properties for `discovery_k8s.access_info.connection_info`

<a id="canonical-0333023303123211-2030310032000233-2300232013233312-0212123313113010-2331220220131330-0301102302000233-1101312232330202-3002010031033221"></a>

#### `discovery_k8s.access_info.connection_info.api_server` property

Type: `"string"`. Computed.

API server must be a fully qualified domain string and port specified as host:port pair.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 262,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 262,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostport": "true",
    "ves.io.schema.rules.string.max_len": "262"
  }
}
```

- [tls_info](data-sources--discovery--reference--group-001.md#canonical-3213112033331231-0002002323300313-3100230200032300-3302120331312303-3131200202121132-3321213301222231-3121330213302123-1131012112231132): complete subsection reference.

<a id="canonical-3213112033331231-0002002323300313-3100230200032300-3302120331312303-3131200202121132-3321213301222231-3121330213302123-1131012112231132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.connection_info.tls_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-3033110220233312-1331012131312110-1301011002212212-0003032332020020-2223302001311301-2012112010022110-0023111321011330-3311111002111211)
- discovery_k8s.access_info.connection_info.tls_info

<a id="canonical-1122312121320222-3322230133213331-2323302103011202-0130002331022221-1001210002112012-1321103011203012-3010203123311013-3201333020123010"></a>

Type: `"single"`. Computed.

TLS config for client of discovery service.

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

<a id="canonical-3213013233222032-3201011003302001-2233330200330222-0120130130332211-3202002320120112-3033000211120220-1003022003200103-1321201130302233"></a>

### Direct properties for `discovery_k8s.access_info.connection_info.tls_info`

<a id="canonical-0311030232100301-0021010122123203-2112002323312110-2023322301033223-1033231113111123-0130013220011121-2132112213203111-0210012230300030"></a>

#### `discovery_k8s.access_info.connection_info.tls_info.certificate` property

Type: `"string"`. Computed.

Client certificate is PEM-encoded certificate or certificate-chain.

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
    "formatDescription": "PEM-encoded X.509 certificate, max 5MB",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 100,
    "pattern": "^-----BEGIN CERTIFICATE-----\\n.*\\n-----END CERTIFICATE-----$",
    "validation": {
      "standard": "PEM"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [key_url](data-sources--discovery--reference--group-001.md#canonical-0011320033010321-0001130232332102-1301132001122023-0022211233131130-0131202000103300-1032132130202331-1002110302222223-3313020213200132): complete subsection reference.

<a id="canonical-0222100210010033-0330313002131111-2221103231021112-3332313322003103-3213031101202133-0013032032300322-1302321230122023-2120120313221033"></a>

<a id="canonical-3123133221100103-2000221112113300-1211203310201101-3213210201133113-0121010232103200-1003230121010330-1113321112333120-2021032100131322"></a>

#### `discovery_k8s.access_info.connection_info.tls_info.server_name` property

Type: `"string"`. Computed.

ServerName is passed to the server for SNI and is used in the client to check server certificates
against. If ServerName is empty, the hostname used to contact the server is used.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1012210311133101-0112000113320111-2110213312023302-3302010113303023-0232330000122133-3013200203111230-1113201113033002-3200112233120311"></a>

<a id="canonical-0021321121331032-3312123110312130-3032023302113300-1102123100320213-3230313031213120-0330013101233211-3200322202333301-3102033211320201"></a>

#### `discovery_k8s.access_info.connection_info.tls_info.trusted_ca_url` property

Type: `"string"`. Computed.

The URL or value for trusted Server CA certificate or certificate chain Certificates in PEM format
including the PEM headers.

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
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0011320033010321-0001130232332102-1301132001122023-0022211233131130-0131202000103300-1032132130202331-1002110302222223-3313020213200132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.connection_info.tls_info.key_url` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-3033110220233312-1331012131312110-1301011002212212-0003032332020020-2223302001311301-2012112010022110-0023111321011330-3311111002111211)
- [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-3213112033331231-0002002323300313-3100230200032300-3302120331312303-3131200202121132-3321213301222231-3121330213302123-1131012112231132)
- discovery_k8s.access_info.connection_info.tls_info.key_url

<a id="canonical-3121211210323100-3012100102121100-2231113312013212-2301331112223331-2311013131323322-2330313030101202-0322131020201230-0000121102223223"></a>

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

<a id="canonical-1110111000022021-2102130200220213-2011201000211303-2321021032300321-3000131300102200-3332200230210222-2223321330330113-3231103133003320"></a>

### Direct properties for `discovery_k8s.access_info.connection_info.tls_info.key_url`

- [blindfold_secret_info](data-sources--discovery--reference--group-001.md#canonical-2130121022010001-2220333021312212-1101333303333310-0010333322302310-1311131002203102-3133113303102313-0001112213223002-0312121232000112): complete subsection reference.

- [clear_secret_info](data-sources--discovery--reference--group-001.md#canonical-2301032222102203-3210023322323102-0310021110233223-0121003331022210-1102302200023022-2301013133031233-3030103012330121-1213313222320313): complete subsection reference.

<a id="canonical-2130121022010001-2220333021312212-1101333303333310-0010333322302310-1311131002203102-3133113303102313-0001112213223002-0312121232000112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-3033110220233312-1331012131312110-1301011002212212-0003032332020020-2223302001311301-2012112010022110-0023111321011330-3311111002111211)
- [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-3213112033331231-0002002323300313-3100230200032300-3302120331312303-3131200202121132-3321213301222231-3121330213302123-1131012112231132)
- [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-0011320033010321-0001130232332102-1301132001122023-0022211233131130-0131202000103300-1032132130202331-1002110302222223-3313020213200132)
- discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info

<a id="canonical-1133311221313203-2311020131130030-1221032221120312-2211302113032312-1313111113132232-2303120013202011-0133112322102301-3230311223201300"></a>

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

<a id="canonical-0233113133131302-0022031302033211-0211302023131030-2201233321221101-2030131233302310-0130111231032010-2223321122021131-2301012111131300"></a>

### Direct properties for `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info`

<a id="canonical-1231023222111200-0312032030011311-0011320231101020-1120321122321310-3010200111231112-2121121332122302-0312102231112311-1322300230231031"></a>

#### `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.decryption_provider` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2023020333101020-3030100213311321-2223122333323031-1323021302302302-1303311003320312-1200033011022000-3311212230221110-1023031302021001"></a>

<a id="canonical-1331100311033313-3223131310230230-2003332232120012-3132103020020332-2212112332010211-3132033131201311-2310320201021032-3100110002203130"></a>

#### `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3230322302310200-0003130320020130-0103312022222122-0010330113033110-2013131133132203-2210302300121132-2210221230323133-0031233231021101"></a>

<a id="canonical-3003002103120032-0221002203011213-0020002030200100-3313233122332312-3323002110110001-0200133030102201-0321313221302203-0012021211210010"></a>

#### `discovery_k8s.access_info.connection_info.tls_info.key_url.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2301032222102203-3210023322323102-0310021110233223-0121003331022210-1102302200023022-2301013133031233-3030103012330121-1213313222320313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info` properties

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md#canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222)
- [Property reference](data-sources--discovery--reference--group-001.md#canonical-2021230332002103-0110323311333230-2003133021000310-3312203112232312-1200102023013220-3201021330020123-3223332331131101-2211133321121220)
- [discovery_k8s](data-sources--discovery--reference--group-001.md#canonical-0330013311111202-1322311031122022-3320203133110210-1213102010333012-0022110212111130-2223102233100201-0022000121230323-1321310123011322)
- [discovery_k8s.access_info](data-sources--discovery--reference--group-001.md#canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233)
- [discovery_k8s.access_info.connection_info](data-sources--discovery--reference--group-001.md#canonical-3033110220233312-1331012131312110-1301011002212212-0003032332020020-2223302001311301-2012112010022110-0023111321011330-3311111002111211)
- [discovery_k8s.access_info.connection_info.tls_info](data-sources--discovery--reference--group-001.md#canonical-3213112033331231-0002002323300313-3100230200032300-3302120331312303-3131200202121132-3321213301222231-3121330213302123-1131012112231132)
- [discovery_k8s.access_info.connection_info.tls_info.key_url](data-sources--discovery--reference--group-001.md#canonical-0011320033010321-0001130232332102-1301132001122023-0022211233131130-0131202000103300-1032132130202331-1002110302222223-3313020213200132)
- discovery_k8s.access_info.connection_info.tls_info.key_url.clear_secret_info

<a id="canonical-0111221031200123-0222001023122123-1222210020103111-3323001201322030-2031310233311020-2323211023331103-2021223302033311-1210223320301122"></a>

Type: `"single"`. Computed.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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
