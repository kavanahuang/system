/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"github.com/kavanahuang/system"
	"testing"
)

func TestIpAddrStart(t *testing.T) {
	Test.Start("IpAddr")
}

func TestGetLocal(t *testing.T) {
	msg := "192.168.1.136"
	result := system.IpAddr.GetLocal()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetLocalReal(t *testing.T) {
	msg := "192.168.1.136"
	result, _ := system.IpAddr.GetLocalReal()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestIpAddrEnd(t *testing.T) {
	Test.End("IpAddr")
}
