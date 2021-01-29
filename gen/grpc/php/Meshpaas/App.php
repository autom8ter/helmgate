<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace MeshelmProxy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * App is an App created from a helm chart
 *
 * Generated from protobuf message <code>helmProxy.App</code>
 */
class App extends \Google\Protobuf\Internal\Message
{
    /**
     * name of the application
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     */
    private $name = '';
    /**
     * namespace name the app belongs to(autocreated)
     *
     * Generated from protobuf field <code>string namespace = 2 [(.validator.field) = {</code>
     */
    private $namespace = '';
    /**
     * release holds information about the currently deployed release of the application
     *
     * Generated from protobuf field <code>.helmProxy.Release release = 5;</code>
     */
    private $release = null;
    /**
     * chart is the chart used to deploy the App
     *
     * Generated from protobuf field <code>.helmProxy.Chart chart = 20;</code>
     */
    private $chart = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name
     *           name of the application
     *     @type string $namespace
     *           namespace name the app belongs to(autocreated)
     *     @type \MeshelmProxy\Release $release
     *           release holds information about the currently deployed release of the application
     *     @type \MeshelmProxy\Chart $chart
     *           chart is the chart used to deploy the App
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * name of the application
     *
     * Generated from protobuf field <code>string name = 1 [(.validator.field) = {</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * name of the application
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
     * namespace name the app belongs to(autocreated)
     *
     * Generated from protobuf field <code>string namespace = 2 [(.validator.field) = {</code>
     * @return string
     */
    public function getNamespace()
    {
        return $this->namespace;
    }

    /**
     * namespace name the app belongs to(autocreated)
     *
     * Generated from protobuf field <code>string namespace = 2 [(.validator.field) = {</code>
     * @param string $var
     * @return $this
     */
    public function setNamespace($var)
    {
        GPBUtil::checkString($var, True);
        $this->namespace = $var;

        return $this;
    }

    /**
     * release holds information about the currently deployed release of the application
     *
     * Generated from protobuf field <code>.helmProxy.Release release = 5;</code>
     * @return \MeshelmProxy\Release
     */
    public function getRelease()
    {
        return $this->release;
    }

    /**
     * release holds information about the currently deployed release of the application
     *
     * Generated from protobuf field <code>.helmProxy.Release release = 5;</code>
     * @param \MeshelmProxy\Release $var
     * @return $this
     */
    public function setRelease($var)
    {
        GPBUtil::checkMessage($var, \MeshelmProxy\Release::class);
        $this->release = $var;

        return $this;
    }

    /**
     * chart is the chart used to deploy the App
     *
     * Generated from protobuf field <code>.helmProxy.Chart chart = 20;</code>
     * @return \MeshelmProxy\Chart
     */
    public function getChart()
    {
        return $this->chart;
    }

    /**
     * chart is the chart used to deploy the App
     *
     * Generated from protobuf field <code>.helmProxy.Chart chart = 20;</code>
     * @param \MeshelmProxy\Chart $var
     * @return $this
     */
    public function setChart($var)
    {
        GPBUtil::checkMessage($var, \MeshelmProxy\Chart::class);
        $this->chart = $var;

        return $this;
    }

}

