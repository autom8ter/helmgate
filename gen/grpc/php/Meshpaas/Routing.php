<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace MeshelmProxy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Routing describes service mesh routing options(gateway/host bindings, route rewrites, etc) for an APIlication
 *
 * Generated from protobuf message <code>helmProxy.Routing</code>
 */
class Routing extends \Google\Protobuf\Internal\Message
{
    /**
     * gateway to bind to
     *
     * Generated from protobuf field <code>string gateway = 1;</code>
     */
    private $gateway = '';
    /**
     * host names to bind to
     *
     * Generated from protobuf field <code>repeated string hosts = 2;</code>
     */
    private $hosts;
    /**
     * http route matchers/configurations
     *
     * Generated from protobuf field <code>repeated .helmProxy.HTTPRoute http_routes = 4;</code>
     */
    private $http_routes;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $gateway
     *           gateway to bind to
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $hosts
     *           host names to bind to
     *     @type \MeshelmProxy\HTTPRoute[]|\Google\Protobuf\Internal\RepeatedField $http_routes
     *           http route matchers/configurations
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * gateway to bind to
     *
     * Generated from protobuf field <code>string gateway = 1;</code>
     * @return string
     */
    public function getGateway()
    {
        return $this->gateway;
    }

    /**
     * gateway to bind to
     *
     * Generated from protobuf field <code>string gateway = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setGateway($var)
    {
        GPBUtil::checkString($var, True);
        $this->gateway = $var;

        return $this;
    }

    /**
     * host names to bind to
     *
     * Generated from protobuf field <code>repeated string hosts = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getHosts()
    {
        return $this->hosts;
    }

    /**
     * host names to bind to
     *
     * Generated from protobuf field <code>repeated string hosts = 2;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setHosts($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->hosts = $arr;

        return $this;
    }

    /**
     * http route matchers/configurations
     *
     * Generated from protobuf field <code>repeated .helmProxy.HTTPRoute http_routes = 4;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getHttpRoutes()
    {
        return $this->http_routes;
    }

    /**
     * http route matchers/configurations
     *
     * Generated from protobuf field <code>repeated .helmProxy.HTTPRoute http_routes = 4;</code>
     * @param \MeshelmProxy\HTTPRoute[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setHttpRoutes($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \MeshelmProxy\HTTPRoute::class);
        $this->http_routes = $arr;

        return $this;
    }

}

