---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2110231031112310-1130011000331010-1312132022002200-1232103213011301-2012100333100023-2000320333201213-1130030102202313-3122301032230233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- Property reference

<a id="canonical-3101333112312310-1102212132231101-3033031231020110-1012230211231222-3332131031332013-2121121012031100-3330212311232010-3210102302221113"></a>

### Direct properties for `xcsh_http_loadbalancer`

- [active_service_policies](resources--http_loadbalancer--reference--group-004.md#canonical-1011302001123301-2313002203212101-0233321031103031-0132102102022202-3111110221001123-1311000231023013-2310321131201012-3120312323220332): complete subsection reference.

<a id="canonical-2130230100111030-3302202230022023-1113230320131330-2302323011332123-1221001122102233-3001010323230311-3112011232310203-1132013131232303"></a>

<a id="canonical-2312333332112022-0102120231001322-2313320111321301-1122223130102221-0312121131001323-3132331012311331-1300130100323322-1113302301122121"></a>

#### `add_location` property

Type: `"bool"`. Optional, Computed.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses. This configuration is ignored on CE sites. Defaults to \`false\`. Server applies
default when omitted.

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

- [advertise_custom](resources--http_loadbalancer--reference--group-004.md#canonical-2020213333220332-1200313101133332-3232201222320323-3032320013323220-2103203302222323-1032010322323133-3000202032001123-1020022302012011): complete subsection reference.

- [advertise_dualstack_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-0223122302110001-1233213021312213-1303312231122322-3210130310233332-1323202101300301-0012222323223300-1111202232330131-0221301031130130): complete subsection reference.

- [advertise_on_public](resources--http_loadbalancer--reference--group-004.md#canonical-2032030213022320-1201110321010130-2223122210020021-2330313311130213-2003011213011220-2003321021213303-2303033002202100-1300231010012311): complete subsection reference.

- [advertise_on_public_default_vip](resources--http_loadbalancer--reference--group-005.md#canonical-3122032220203002-1320231213301220-1303120120111230-2231321223321013-0331311110130203-3101201330021212-3102330002010010-0011210003233122): complete subsection reference.

- [advertise_v6_on_public](resources--http_loadbalancer--reference--group-005.md#canonical-2133220122231331-1012333101303110-2323230033102213-1132133010132311-1200323020332232-0211032223320321-2020233002320131-3320110032220210): complete subsection reference.

<a id="canonical-0100222232331102-0002021002322033-2101221222032201-0103120333001033-2102102002212112-3112122011320030-1203101301030310-2310300203212003"></a>

<a id="canonical-0022011121310220-3331330120113032-1310133302313033-1301011030013311-2201323321212202-3101223021023222-3020001112210012-3233300311022111"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

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

- [api_protection_rules](resources--http_loadbalancer--reference--group-005.md#canonical-0323010320013112-1013131323103311-2220212212202133-0201213033222111-0232211100100103-3333310132333133-0223030120200030-1222221022201130): complete subsection reference.

- [api_rate_limit](resources--http_loadbalancer--reference--group-006.md#canonical-0123000123030121-0102233302330020-2123112021130221-3331312230203303-3313013120012121-2000020111201032-0333203033020101-2322013121112102): complete subsection reference.

- [api_specification](resources--http_loadbalancer--reference--group-008.md#canonical-3203213310123232-0113311031200013-1133033113100132-3300021031111113-2232100021003100-2223211022002310-0012032220303032-2131211111300100): complete subsection reference.

- [api_testing](resources--http_loadbalancer--reference--group-009.md#canonical-1221213100202332-3122311330122203-3123201111011003-2300313322003311-0010003220323322-1031002030311120-3130302333032122-3331322112221301): complete subsection reference.

- [app_firewall](resources--http_loadbalancer--reference--group-009.md#canonical-2333332121111031-2232311201113212-1330231222032103-0232233320112212-0011002003013022-0302021111130330-0013131230131031-3111322301332033): complete subsection reference.

- [blocked_clients](resources--http_loadbalancer--reference--group-010.md#canonical-0322023221021312-1122332013300331-0213123111330200-2212020301212310-1312102131121221-0111203031300302-3010332303211031-3103000311130112): complete subsection reference.

- [bot_defense](resources--http_loadbalancer--reference--group-010.md#canonical-0130323233303113-2010203332130222-3312321113303112-1330013223031011-3113112323303320-2130211112113312-1300122111102332-0233232302123030): complete subsection reference.

- [bot_defense_advanced_protection](resources--http_loadbalancer--reference--group-011.md#canonical-3021132012201110-0202200030201230-2103020232303031-2312002102323313-0312331332213000-3313003100012220-3030232323333222-2330313002120331): complete subsection reference.

- [caching_policy](resources--http_loadbalancer--reference--group-013.md#canonical-3102000111002330-2330012203002202-2013302232030202-2322011030203033-0013300220121222-2223102330220322-1011323030221103-3323302013003331): complete subsection reference.

- [captcha_challenge](resources--http_loadbalancer--reference--group-013.md#canonical-3331001122010103-1330222312203012-2102022330310300-3121202020223232-3202122212132233-0220222003113103-2301213310311323-1232121323232302): complete subsection reference.

- [client_side_defense](resources--http_loadbalancer--reference--group-013.md#canonical-0311102213112310-0312111212032223-3211013302233313-0112122213121120-1300023311013320-1323213131000131-0110123300213032-2230120330012222): complete subsection reference.

- [cookie_stickiness](resources--http_loadbalancer--reference--group-014.md#canonical-3322222220120200-1330331312120131-2323231032320212-0121111210230332-1011313101330022-0002302120301202-0110200313122233-1010303103322111): complete subsection reference.

- [cors_policy](resources--http_loadbalancer--reference--group-014.md#canonical-2120230123301020-2223032123333110-0222303131313113-1302023032311223-0231200001201300-0311101003010332-1110132103312122-1021322003001231): complete subsection reference.

- [csrf_policy](resources--http_loadbalancer--reference--group-014.md#canonical-3021222013003320-3230102012010221-2203233323210232-0111132233210023-1332310233100013-1310012310002212-3333112112231001-0212212331211311): complete subsection reference.

- [data_guard_rules](resources--http_loadbalancer--reference--group-014.md#canonical-0310100223013020-2111222133111333-2201233312231132-0102103013122132-0111220303113031-1122123223301021-1030212023023130-0331013321302221): complete subsection reference.

- [ddos_mitigation_rules](resources--http_loadbalancer--reference--group-014.md#canonical-3113021303031131-1312012222010231-0331133200303132-3332023021101222-0113133213312010-1321130002311311-0012301201200102-1130020133302102): complete subsection reference.

- [default_pool](resources--http_loadbalancer--reference--group-015.md#canonical-3011031210301023-3111303122131111-0122010032202013-2011101001022200-0221332123322110-1223213221131013-1210333203300122-3013123231031011): complete subsection reference.

- [default_pool_list](resources--http_loadbalancer--reference--group-017.md#canonical-0232023012330101-1311101132201100-1333212211111332-1320223013300103-2111300101212212-3332332203230321-3210321213221300-3030022102121311): complete subsection reference.

- [default_route_pools](resources--http_loadbalancer--reference--group-017.md#canonical-0223203012003322-2131303223200222-3233202001321210-0211213220211130-2231301101131022-2223313123013013-2330111113211022-3120310230100133): complete subsection reference.

- [default_sensitive_data_policy](resources--http_loadbalancer--reference--group-017.md#canonical-2012113030000132-3012100211231221-3130003303031233-0223013133101100-2030232213320012-3201201203311111-0130303112131112-1022213320131303): complete subsection reference.

<a id="canonical-2110120202003313-2333311220302223-3102000123110232-0131130033130220-2332301210002120-3120233132311222-3311101003210313-0230122211321321"></a>

<a id="canonical-0303020210001331-1131223212302313-0321131102133213-2332223100202222-2321121221223110-1130331232223102-0013221210032303-1232032033231311"></a>

#### `description` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3302302223013230-0320213300130032-3211002023033213-2121320021023103-2301222122121000-2023201213122330-2102211013001102-2210100213223130"></a>

<a id="canonical-3021302011332321-3011031020310223-1300230122320211-1200103023221311-3332032301133212-2222133230111011-1320310200310120-3311102331121023"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

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

- [disable_api_definition](resources--http_loadbalancer--reference--group-017.md#canonical-0123111133030011-2123003110002011-1031032130100110-1022210200133330-3303330101000301-0203011300213011-1020303111133311-2132323123000212): complete subsection reference.

- [disable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-0232001202112132-3122220002210323-2123112210120123-2111123031002213-0132120133332022-2111330310202011-2220013001013012-0203213300003333): complete subsection reference.

- [disable_api_testing](resources--http_loadbalancer--reference--group-017.md#canonical-3313202103132030-0013213023113203-2231222203010002-3132330201211211-1031002112330010-2320302022023023-1021321113132320-2221311103102213): complete subsection reference.

- [disable_bot_defense](resources--http_loadbalancer--reference--group-017.md#canonical-3123003300002010-0033323131033110-3101003330221223-1232122203203021-0032123103302012-3321133121110332-3110131110231100-2002011203003323): complete subsection reference.

- [disable_caching](resources--http_loadbalancer--reference--group-017.md#canonical-2321101032330312-2021001302001033-0303012101301113-3221010231232001-2031231033003121-0102030231321121-2103303220300112-2032023320022021): complete subsection reference.

- [disable_client_side_defense](resources--http_loadbalancer--reference--group-017.md#canonical-2122321323123120-0131330211013033-0223101332002110-1133012200030110-0300010032111010-3133032130220300-0032001323100111-3300101020203300): complete subsection reference.

- [disable_ip_reputation](resources--http_loadbalancer--reference--group-017.md#canonical-1202122330212221-3201331122203123-3033330230333010-1333110133101010-1102330300133111-1223133210101012-2101013113121223-0330312222221002): complete subsection reference.

- [disable_malicious_user_detection](resources--http_loadbalancer--reference--group-017.md#canonical-3012231223303232-2102332203003323-3102212131101310-2012202210031310-2221201202003333-1132132220302101-0321000202112303-1321313331130122): complete subsection reference.

- [disable_malware_protection](resources--http_loadbalancer--reference--group-017.md#canonical-1303333110313121-3231031113100003-0101230011321330-0011203322101330-0102021301032021-1333303102122201-3123112221102102-3010103301030101): complete subsection reference.

- [disable_rate_limit](resources--http_loadbalancer--reference--group-017.md#canonical-3010203201310323-0211110203000202-0211202211033123-0200222331221023-2010100321330113-1113001031021000-2100021313301232-0222321320030113): complete subsection reference.

- [disable_threat_mesh](resources--http_loadbalancer--reference--group-017.md#canonical-2201101022220023-1103210213332120-0012201103310011-2012220223033011-2311212133223002-2010030131033033-1301101200102210-0132000311201023): complete subsection reference.

- [disable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-017.md#canonical-2010130233211201-0102210200231211-2331021322013011-3201112123233301-1203313132231022-1130220320000222-0131212013022231-0330001200132223): complete subsection reference.

- [disable_waf](resources--http_loadbalancer--reference--group-017.md#canonical-0231033130322231-2203223313010330-1200320322113021-1333023133103333-3331123032011122-3320211321333012-0333310010101210-0312032310201021): complete subsection reference.

- [do_not_advertise](resources--http_loadbalancer--reference--group-017.md#canonical-2302213303120331-0322020022021133-2103031201103012-0013213333030332-3331330132313320-2112113133220233-3320130323210022-3120123311111331): complete subsection reference.

<a id="canonical-1303231310010021-3121222003222121-0333121331303011-3032213131100231-3321010203030321-0131122123220020-3031110223131113-1333002320202000"></a>

<a id="canonical-2133002332032131-3113011011220131-0321001011001010-1321011333012103-3322332012013120-3333113201203001-0133232313223223-0330300030003120"></a>

#### `domains` property

Type: `["list", "string"]`. Required.

A list of Domains (host/authority header) that will be matched to load balancer.

Supported Domains and search order: &#8203;1. Exact Domain names: www&#46;example.com. &#8203;2.
Domains starting with a Wildcard: \*.example.com.

Not supported Domains: &#8203;- Just a Wildcard: \* &#8203;- A Wildcard and TLD with no root Domain:
\*.com. &#8203;- A Wildcard not matching a whole DNS label. E.g. \*.example.com and
\*.bar.example.com are valid Wildcards however \*bar.example.com, \*-bar.example.com, and
bar\*.example.com are all invalid.

Additional notes: A Wildcard will not match empty string. E.g. \*.example.com will match
bar.example.com and baz-bar.example.com but not .example.com. The longest Wildcards match first.
Only a single virtual host in the entire route configuration can match on \*. Also a Domain must be
unique across all virtual hosts within an advertise policy.

Domains are also used for SNI matching if the Loadbalancer type is HTTPS. Domains also indicate the
list of names for which DNS resolution will be automatically resolved to IP addresses by the system.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [enable_api_discovery](resources--http_loadbalancer--reference--group-017.md#canonical-3021021121021003-1312211130321032-1020103201123033-0123000000320211-1120031012321021-0223212332033233-1310033103202120-1201002111130321): complete subsection reference.

- [enable_challenge](resources--http_loadbalancer--reference--group-018.md#canonical-1000022300203233-0132122023301131-1030311311113322-1331030020001010-3313300121123222-1021222331031303-0302013312001131-2001112232233232): complete subsection reference.

- [enable_ip_reputation](resources--http_loadbalancer--reference--group-018.md#canonical-2231212223321013-0123213202300103-1002202103000112-0121311132023203-1003223131331322-3021301311230232-3302012302222311-0121321022022323): complete subsection reference.

- [enable_malicious_user_detection](resources--http_loadbalancer--reference--group-018.md#canonical-2102301301321320-1331123320232110-0223023132023311-1300312321320102-3102112231332231-2233232100233210-3122230110022023-1302310210123220): complete subsection reference.

- [enable_threat_mesh](resources--http_loadbalancer--reference--group-018.md#canonical-3111322010113111-1002113001030101-3211022303110233-2133221032220321-2011101130012312-1232301001111000-1203020010201230-2331021131223233): complete subsection reference.

- [enable_trust_client_ip_headers](resources--http_loadbalancer--reference--group-018.md#canonical-2023200122231012-2002203313103003-2012123021321211-3110001220301030-2122312230032131-3200023120120031-3000230111323102-0321222120322033): complete subsection reference.

- [graphql_rules](resources--http_loadbalancer--reference--group-018.md#canonical-0103131211230021-3322322013030331-0002300200111211-3023323230210012-3210222101001300-3011023202322230-0123031212122011-2102122120303112): complete subsection reference.

- [http](resources--http_loadbalancer--reference--group-018.md#canonical-3201330110311330-2123203220332301-2010310002232302-2200333212132031-2112023132010213-3001201001222010-0222333320131100-0013020102021323): complete subsection reference.

- [https](resources--http_loadbalancer--reference--group-018.md#canonical-0321232101200210-0011302000022301-2002321223210110-0303132310312201-0222311023330031-3231011200231021-2012211321122303-0001332120000102): complete subsection reference.

- [https_auto_cert](resources--http_loadbalancer--reference--group-019.md#canonical-3303010332201213-1302123100300210-3320131100121233-3022222101133310-0310203202321313-1322002013203221-0212030100121103-2103331200222220): complete subsection reference.

<a id="canonical-3200220013003332-2003233223131021-1333032200203330-1121132300131312-0310331330232102-2212120223122320-0112001023323210-0300203120020123"></a>

<a id="canonical-3121122123323122-2131100303100203-0211013011321300-2003133013102212-3220122000122211-0120131011023330-1233110222221011-0232010113130112"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

- [js_challenge](resources--http_loadbalancer--reference--group-019.md#canonical-0122213100023133-3102301113333103-2221320301323033-3302232320022112-0012001332132331-1022120132002212-0201313110302323-3320131033130221): complete subsection reference.

- [jwt_validation](resources--http_loadbalancer--reference--group-019.md#canonical-1031330302322103-2310332322003013-1202332212322233-1121331003012130-2311231031230333-1101231002121010-2211310320222130-0021121131033333): complete subsection reference.

- [l7_ddos_action_block](resources--http_loadbalancer--reference--group-020.md#canonical-1320110132020001-0022333031222332-0332001223201230-0000302202010121-2213203130001312-3221132020131213-3301310301230123-2220221132033300): complete subsection reference.

- [l7_ddos_action_default](resources--http_loadbalancer--reference--group-020.md#canonical-3122303101123332-1003322121300112-3112102023031013-1001133230202001-1121101002133313-0121023032001313-1200303023002210-0310210303322333): complete subsection reference.

- [l7_ddos_action_js_challenge](resources--http_loadbalancer--reference--group-020.md#canonical-1130022200313200-0012221203233323-2031132210123022-0122213131223223-0033223300233010-3030023013112322-3321030210121223-0130102321222031): complete subsection reference.

- [l7_ddos_protection](resources--http_loadbalancer--reference--group-020.md#canonical-2211021201313113-1021310021013033-1302002202111010-3312000330210330-2322112331020032-2233321200122022-2021112031123310-3213230222300211): complete subsection reference.

<a id="canonical-3020010001030301-2132021001231203-3333220320322020-2031320223030231-0110033230012211-3001313111301100-1101333023200311-0213131303001033"></a>

<a id="canonical-0302322331032012-1233333032121230-3033301201303123-0233103333220113-0220020311333101-1111121020133012-2011112303323323-0012003221122023"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

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

- [least_active](resources--http_loadbalancer--reference--group-020.md#canonical-3100013201201320-1030023222021000-1020231012221032-2100130020230332-3130310121332223-2110221222102103-3201213033132002-0230230101332320): complete subsection reference.

- [malware_protection_settings](resources--http_loadbalancer--reference--group-020.md#canonical-2002001101312212-3322201121023103-1300033312332032-2332131302121221-0222023030002220-1013231333110113-1001302002032133-0122132213332322): complete subsection reference.

- [more_option](resources--http_loadbalancer--reference--group-020.md#canonical-0103232331000000-3202301220200212-2333013000233203-1310210130333000-1220011203132102-0102220112330233-2210002322102333-1111002302303232): complete subsection reference.

- [multi_lb_app](resources--http_loadbalancer--reference--group-021.md#canonical-0022320220230323-0223203133103320-3032023020320010-3000011312022000-3332213021010220-3323011033022102-3132320201321311-0310230130132011): complete subsection reference.

<a id="canonical-0022320022000223-0010210032222011-3312021020031011-3103003212200302-3023333301210020-3332223123323132-3323133203103021-1211132303023232"></a>

<a id="canonical-1321312331033110-2200013232032101-2101201121133221-2033102112313330-1333213103212101-1110330003023023-2030211001123232-3111211331110222"></a>

#### `name` property

Type: `"string"`. Required.

Name of the HTTP Load Balancer. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0032330033211321-2101033331312320-3133102230032000-0201103233012210-1200330300200120-1113022102300321-0312123322223223-0011331011121300"></a>

<a id="canonical-2112210103001013-1333211213021001-1200111300130031-0312031232101102-2130203200131231-1311012202232100-0000233300100231-3002102310312120"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the HTTP Load Balancer is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [no_challenge](resources--http_loadbalancer--reference--group-021.md#canonical-2101121220322123-0303121132222013-1232012001122310-2310001013023221-1200112133200300-0332032203023122-3310012322111310-2011133110323332): complete subsection reference.

- [no_service_policies](resources--http_loadbalancer--reference--group-022.md#canonical-1213221123121332-0130210312020112-0030010203112202-1320321321320032-1222212130023330-2011211132030330-2231013002131030-0013313003020331): complete subsection reference.

- [origin_server_subset_rule_list](resources--http_loadbalancer--reference--group-022.md#canonical-0213330010003320-3033321220120021-2320212310032100-2232122130011320-2133212222310031-1313023301001000-3333100122210230-1131233020121300): complete subsection reference.

- [policy_based_challenge](resources--http_loadbalancer--reference--group-022.md#canonical-1321211020201321-0311100120203222-3322203000322000-0222330310101000-2122021101303131-2012000320103220-3330123333130100-2210013121130111): complete subsection reference.

- [protected_cookies](resources--http_loadbalancer--reference--group-024.md#canonical-0301320313311333-3120323010313020-3332021012130022-0322212212001002-2230020110210323-2202032301310033-1331201211023033-1201101312100233): complete subsection reference.

- [random](resources--http_loadbalancer--reference--group-024.md#canonical-1312120313213323-2103003010102231-1132113202300010-0103032212233002-0220031323220302-3231002030321130-0010020301002111-1312011221320111): complete subsection reference.

- [rate_limit](resources--http_loadbalancer--reference--group-024.md#canonical-1330231231302021-0002122303001301-3122013030133233-0310303222330222-3030303330233332-2320022313131101-3222310032010331-3200322201120023): complete subsection reference.

- [ring_hash](resources--http_loadbalancer--reference--group-024.md#canonical-3303333311310122-3120022023110301-0133031123332301-3313201312111021-3012232201131020-3210032223213123-3232331032121200-3322221201100330): complete subsection reference.

- [round_robin](resources--http_loadbalancer--reference--group-025.md#canonical-3120201302000120-1101032023332332-1203301030103211-2301112131023100-3323333313021213-2102103120123013-3211232030113002-1310102202033321): complete subsection reference.

- [routes](resources--http_loadbalancer--reference--group-025.md#canonical-1110012221012033-1221210233120021-3023210312201220-1131233300133330-0221332023232023-2122130212211030-3231100102132103-2203232033231100): complete subsection reference.

- [sensitive_data_disclosure_rules](resources--http_loadbalancer--reference--group-027.md#canonical-3010130000211103-2132323303233130-3123203102231323-2302320310112320-0023100231121123-3223331112120111-3300221200232213-1133321320223332): complete subsection reference.

- [sensitive_data_policy](resources--http_loadbalancer--reference--group-027.md#canonical-3300321023010221-0210333331122211-1013103230213030-3001321133032323-0100212010332121-1322030300320113-3031022123002130-1203331012330020): complete subsection reference.

- [service_policies_from_namespace](resources--http_loadbalancer--reference--group-027.md#canonical-3320321312231220-2203023313122012-3333323031313130-0133112032311321-1013010313122112-3211000010020331-1110330001120132-0323002112221111): complete subsection reference.

- [single_lb_app](resources--http_loadbalancer--reference--group-027.md#canonical-1310003232121331-1310031211031211-0032102003301212-1322003031311002-0103320322323000-1030132321201022-1330102301233302-0231033122321203): complete subsection reference.

- [slow_ddos_mitigation](resources--http_loadbalancer--reference--group-027.md#canonical-3333311120100023-2230031010023220-2013013300321213-0302220321233313-0222113033133022-2131111320010323-2303320333133122-3002301330000322): complete subsection reference.

- [source_ip_stickiness](resources--http_loadbalancer--reference--group-027.md#canonical-0120111122231002-2333331020233201-3020032131233000-0232201223221311-1012212322212322-1113023212320102-3111221221223322-0223221231203000): complete subsection reference.

- [system_default_timeouts](resources--http_loadbalancer--reference--group-027.md#canonical-3111101023101332-0113013332131030-0031310130033201-2323032100231211-0300131012212002-1223021312230103-2033033310310331-1121322330011111): complete subsection reference.

- [timeouts](resources--http_loadbalancer--reference--group-027.md#canonical-0031000301301132-0002232123000132-3130011021001231-3113311013210310-0122300301203211-0213222010232211-0313220210221220-2120201133330322): complete subsection reference.

- [trusted_clients](resources--http_loadbalancer--reference--group-027.md#canonical-3200223300000312-1133203300322333-3030222303022110-0233122302130212-2300200012221111-3230203222102303-2311020331311330-3223231332103021): complete subsection reference.

- [user_id_client_ip](resources--http_loadbalancer--reference--group-028.md#canonical-3332013311013213-1233202210102101-2202301332332033-2130302002311222-1231331220201121-2021122303232113-1311030310321233-1201033022300000): complete subsection reference.

- [user_identification](resources--http_loadbalancer--reference--group-028.md#canonical-0032001213002301-0331120033322333-0030100320321033-2022130323111321-1033121323230311-3230223120103201-0000013032330321-1020332302220223): complete subsection reference.

- [waf_exclusion](resources--http_loadbalancer--reference--group-028.md#canonical-2012120311210212-3133332122100012-3032012321322102-0211001032100121-2032033001003003-1322303203221120-2131302203112023-0001031031001101): complete subsection reference.
