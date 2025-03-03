/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import React, { useMemo } from 'react';
import { Button, Intent } from '@blueprintjs/core';

import type { ColumnType } from '@/components';
import { Plugins, DataScopeList } from '@/plugins';

import type { BPConnectionItemType } from '../types';

import * as S from './styled';

interface Props {
  onDetail: (connection: BPConnectionItemType) => void;
  onDelete: (plugin: Plugins, connectionId: ID, scopeId: ID) => void;
}

export const useColumns = ({ onDetail, onDelete }: Props) => {
  return useMemo(
    () =>
      [
        {
          title: 'Data Connections',
          dataIndex: ['icon', 'name'],
          key: 'connection',
          render: ({ icon, name }: Pick<BPConnectionItemType, 'icon' | 'name'>) => (
            <S.ConnectionColumn>
              <img src={icon} alt="" />
              <span>{name}</span>
            </S.ConnectionColumn>
          ),
        },
        {
          title: 'Data Scope',
          dataIndex: ['plugin', 'id', 'scopeIds'],
          key: 'unique',
          render: ({ plugin, id, scopeIds }: Pick<BPConnectionItemType, 'plugin' | 'id' | 'scopeIds'>) => (
            <DataScopeList
              groupByTs={false}
              plugin={plugin}
              connectionId={id}
              scopeIds={scopeIds}
              onDelete={onDelete}
            />
          ),
        },
        {
          title: '',
          key: 'action',
          align: 'center',
          render: (_: any, connection: BPConnectionItemType) => (
            <Button
              small
              minimal
              intent={Intent.PRIMARY}
              icon="add"
              text="Add Scope"
              onClick={() => onDetail(connection)}
            />
          ),
        },
      ] as ColumnType<BPConnectionItemType>,
    [onDetail, onDelete],
  );
};
