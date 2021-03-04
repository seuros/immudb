/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package stream

import "io"

type zStreamSender struct {
	s MsgSender
}

func NewZStreamSender(s MsgSender) *zStreamSender {
	return &zStreamSender{
		s: s,
	}
}

func (st *zStreamSender) Send(ze *ZEntry) error {
	err := st.s.Send(ze.Set.Content, ze.Set.Size)
	if err != nil {
		if err == io.EOF {
			return st.s.RecvMsg(nil)
		}
		return err
	}

	err = st.s.Send(ze.Key.Content, ze.Key.Size)
	if err != nil {
		if err == io.EOF {
			return st.s.RecvMsg(nil)
		}
		return err
	}

	err = st.s.Send(ze.Score.Content, ze.Score.Size)
	if err != nil {
		if err == io.EOF {
			return st.s.RecvMsg(nil)
		}
		return err
	}

	err = st.s.Send(ze.Value.Content, ze.Value.Size)
	if err != nil {
		if err == io.EOF {
			return st.s.RecvMsg(nil)
		}
		return err
	}

	return nil
}
