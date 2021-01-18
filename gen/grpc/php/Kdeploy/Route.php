<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: kdeploy.proto

namespace Kdeploy;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>kdeploy.Route</code>
 */
class Route extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated string hosts = 1;</code>
     */
    private $hosts;
    /**
     * Generated from protobuf field <code>repeated string gateways = 2;</code>
     */
    private $gateways;
    /**
     * Generated from protobuf field <code>string path_prefix = 3;</code>
     */
    private $path_prefix = '';
    /**
     * Generated from protobuf field <code>string rewrite_uri = 4;</code>
     */
    private $rewrite_uri = '';
    /**
     * Generated from protobuf field <code>repeated string allow_origins = 5;</code>
     */
    private $allow_origins;
    /**
     * Generated from protobuf field <code>repeated string allow_methods = 6;</code>
     */
    private $allow_methods;
    /**
     * Generated from protobuf field <code>repeated string allow_headers = 7;</code>
     */
    private $allow_headers;
    /**
     * Generated from protobuf field <code>repeated string expose_headers = 8;</code>
     */
    private $expose_headers;
    /**
     * Generated from protobuf field <code>bool allow_credentials = 9;</code>
     */
    private $allow_credentials = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $hosts
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $gateways
     *     @type string $path_prefix
     *     @type string $rewrite_uri
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $allow_origins
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $allow_methods
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $allow_headers
     *     @type string[]|\Google\Protobuf\Internal\RepeatedField $expose_headers
     *     @type bool $allow_credentials
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Kdeploy::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated string hosts = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getHosts()
    {
        return $this->hosts;
    }

    /**
     * Generated from protobuf field <code>repeated string hosts = 1;</code>
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
     * Generated from protobuf field <code>repeated string gateways = 2;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getGateways()
    {
        return $this->gateways;
    }

    /**
     * Generated from protobuf field <code>repeated string gateways = 2;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setGateways($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->gateways = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string path_prefix = 3;</code>
     * @return string
     */
    public function getPathPrefix()
    {
        return $this->path_prefix;
    }

    /**
     * Generated from protobuf field <code>string path_prefix = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setPathPrefix($var)
    {
        GPBUtil::checkString($var, True);
        $this->path_prefix = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string rewrite_uri = 4;</code>
     * @return string
     */
    public function getRewriteUri()
    {
        return $this->rewrite_uri;
    }

    /**
     * Generated from protobuf field <code>string rewrite_uri = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setRewriteUri($var)
    {
        GPBUtil::checkString($var, True);
        $this->rewrite_uri = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string allow_origins = 5;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAllowOrigins()
    {
        return $this->allow_origins;
    }

    /**
     * Generated from protobuf field <code>repeated string allow_origins = 5;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAllowOrigins($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->allow_origins = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string allow_methods = 6;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAllowMethods()
    {
        return $this->allow_methods;
    }

    /**
     * Generated from protobuf field <code>repeated string allow_methods = 6;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAllowMethods($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->allow_methods = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string allow_headers = 7;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getAllowHeaders()
    {
        return $this->allow_headers;
    }

    /**
     * Generated from protobuf field <code>repeated string allow_headers = 7;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setAllowHeaders($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->allow_headers = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated string expose_headers = 8;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getExposeHeaders()
    {
        return $this->expose_headers;
    }

    /**
     * Generated from protobuf field <code>repeated string expose_headers = 8;</code>
     * @param string[]|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setExposeHeaders($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::STRING);
        $this->expose_headers = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool allow_credentials = 9;</code>
     * @return bool
     */
    public function getAllowCredentials()
    {
        return $this->allow_credentials;
    }

    /**
     * Generated from protobuf field <code>bool allow_credentials = 9;</code>
     * @param bool $var
     * @return $this
     */
    public function setAllowCredentials($var)
    {
        GPBUtil::checkBool($var);
        $this->allow_credentials = $var;

        return $this;
    }

}

