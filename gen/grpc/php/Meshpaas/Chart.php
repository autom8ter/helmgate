<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace MeshelmProxy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Chart is a helm chart that may be used to deploy an App
 *
 * Generated from protobuf message <code>helmProxy.Chart</code>
 */
class Chart extends \Google\Protobuf\Internal\Message
{
    /**
     * name of the app chart
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     */
    private $name = '';
    /**
     * home page of the app chart
     *
     * Generated from protobuf field <code>string home = 2;</code>
     */
    private $home = '';
    /**
     * description of the app chart
     *
     * Generated from protobuf field <code>string description = 3;</code>
     */
    private $description = '';
    /**
     * version of the app chart
     *
     * Generated from protobuf field <code>string version = 4;</code>
     */
    private $version = '';
    /**
     * Generated from protobuf field <code>repeated string sources = 5;</code>
     */
    private $sources;
    /**
     * keywords associated with the app chart
     *
     * Generated from protobuf field <code>repeated string keywords = 6;</code>
     */
    private $keywords;
    /**
     * icon is an the icon/brand associated with the chart
     *
     * Generated from protobuf field <code>string icon = 7;</code>
     */
    private $icon = '';
    /**
     * chart is not actively maintained if deprecated = true
     *
     * Generated from protobuf field <code>bool deprecated = 8;</code>
     */
    private $deprecated = false;
    /**
     * extra charts that this chart depends on
     *
     * Generated from protobuf field <code>repeated .helmProxy.Dependency dependencies = 9;</code>
     */
    private $dependencies;
    /**
     * maintainers of this chart
     *
     * Generated from protobuf field <code>repeated .helmProxy.Maintainer maintainers = 10;</code>
     */
    private $maintainers;
    /**
     * arbitrary metadata associated with the chart
     *
     * Generated from protobuf field <code>map<string, string> metadata = 11;</code>
     */
    private $metadata;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *           name of the app chart
     *     @type string $home
     *           home page of the app chart
     *     @type string $description
     *           description of the app chart
     *     @type string $version
     *           version of the app chart
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $sources
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $keywords
     *           keywords associated with the app chart
     *     @type string $icon
     *           icon is an the icon/brand associated with the chart
     *     @type bool $deprecated
     *           chart is not actively maintained if deprecated = true
     *     @type \MeshelmProxy\Dependency[]|\Google\Protobuf\Internal\RepeatedField $dependencies
     *           extra charts that this chart depends on
     *     @type \MeshelmProxy\Maintainer[]|\Google\Protobuf\Internal\RepeatedField $maintainers
     *           maintainers of this chart
     *     @type array|\Google\Protobuf\Internal\MapField $metadata
     *           arbitrary metadata associated with the chart
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * name of the app chart
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * name of the app chart
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * home page of the app chart
     *
     * Generated from protobuf field <code>string home = 2;</code>
     * @return string
     */
    public function getHome()
    {
        return $this->home;
    }

    /**
     * home page of the app chart
     *
     * Generated from protobuf field <code>string home = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setHome($var)
    {
        GPBUtil::checkString($var, True);
        $this->home = $var;

        return $this;
    }

    /**
     * description of the app chart
     *
     * Generated from protobuf field <code>string description = 3;</code>
     * @return string
     */
    public function getDescription()
    {
        return $this->description;
    }

    /**
     * description of the app chart
     *
     * Generated from protobuf field <code>string description = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setDescription($var)
    {
        GPBUtil::checkString($var, True);
        $this->description = $var;

        return $this;
    }

    /**
     * version of the app chart
     *
     * Generated from protobuf field <code>string version = 4;</code>
     * @return string
     */
    public function getVersion()
    {
        return $this->version;
    }

    /**
     * version of the app chart
     *
     * Generated from protobuf field <code>string version = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setVersion($var)
    {
        GPBUtil::checkString($var, True);
        $this->version = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string sources = 5;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getSources()
    {
        return $this->sources;
    }

    /**
     * Generated from protobuf field <code>repeated string sources = 5;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setSources($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->sources = $arr;

        return $this;
    }

    /**
     * keywords associated with the app chart
     *
     * Generated from protobuf field <code>repeated string keywords = 6;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getKeywords()
    {
        return $this->keywords;
    }

    /**
     * keywords associated with the app chart
     *
     * Generated from protobuf field <code>repeated string keywords = 6;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setKeywords($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->keywords = $arr;

        return $this;
    }

    /**
     * icon is an the icon/brand associated with the chart
     *
     * Generated from protobuf field <code>string icon = 7;</code>
     * @return string
     */
    public function getIcon()
    {
        return $this->icon;
    }

    /**
     * icon is an the icon/brand associated with the chart
     *
     * Generated from protobuf field <code>string icon = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setIcon($var)
    {
        GPBUtil::checkString($var, True);
        $this->icon = $var;

        return $this;
    }

    /**
     * chart is not actively maintained if deprecated = true
     *
     * Generated from protobuf field <code>bool deprecated = 8;</code>
     * @return bool
     */
    public function getDeprecated()
    {
        return $this->deprecated;
    }

    /**
     * chart is not actively maintained if deprecated = true
     *
     * Generated from protobuf field <code>bool deprecated = 8;</code>
     * @param bool $var
     * @return $this
     */
    public function setDeprecated($var)
    {
        GPBUtil::checkBool($var);
        $this->deprecated = $var;

        return $this;
    }

    /**
     * extra charts that this chart depends on
     *
     * Generated from protobuf field <code>repeated .helmProxy.Dependency dependencies = 9;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getDependencies()
    {
        return $this->dependencies;
    }

    /**
     * extra charts that this chart depends on
     *
     * Generated from protobuf field <code>repeated .helmProxy.Dependency dependencies = 9;</code>
     * @param \MeshelmProxy\Dependency[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setDependencies($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \MeshelmProxy\Dependency::class);
        $this->dependencies = $arr;

        return $this;
    }

    /**
     * maintainers of this chart
     *
     * Generated from protobuf field <code>repeated .helmProxy.Maintainer maintainers = 10;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getMaintainers()
    {
        return $this->maintainers;
    }

    /**
     * maintainers of this chart
     *
     * Generated from protobuf field <code>repeated .helmProxy.Maintainer maintainers = 10;</code>
     * @param \MeshelmProxy\Maintainer[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setMaintainers($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \MeshelmProxy\Maintainer::class);
        $this->maintainers = $arr;

        return $this;
    }

    /**
     * arbitrary metadata associated with the chart
     *
     * Generated from protobuf field <code>map<string, string> metadata = 11;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getMetadata()
    {
        return $this->metadata;
    }

    /**
     * arbitrary metadata associated with the chart
     *
     * Generated from protobuf field <code>map<string, string> metadata = 11;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setMetadata($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->metadata = $arr;

        return $this;
    }

}

