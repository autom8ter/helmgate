<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace Helmgate;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * HistoryFilter is used to query a timeseries of releases for a specific app/release
 *
 * Generated from protobuf message <code>helmgate.HistoryFilter</code>
 */
class HistoryFilter extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.helmgate.AppRef ref = 1 [(.validator.field) = {</code>
     */
    private $ref = null;
    /**
     * Generated from protobuf field <code>uint32 limit = 2;</code>
     */
    private $limit = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Helmgate\AppRef $ref
     *     @type int $limit
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Schema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.helmgate.AppRef ref = 1 [(.validator.field) = {</code>
     * @return \Helmgate\AppRef
     */
    public function getRef()
    {
        return $this->ref;
    }

    /**
     * Generated from protobuf field <code>.helmgate.AppRef ref = 1 [(.validator.field) = {</code>
     * @param \Helmgate\AppRef $var
     * @return $this
     */
    public function setRef($var)
    {
        GPBUtil::checkMessage($var, \Helmgate\AppRef::class);
        $this->ref = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>uint32 limit = 2;</code>
     * @return int
     */
    public function getLimit()
    {
        return $this->limit;
    }

    /**
     * Generated from protobuf field <code>uint32 limit = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setLimit($var)
    {
        GPBUtil::checkUint32($var);
        $this->limit = $var;

        return $this;
    }

}
