<template>
  <a-card>
    <div :class="advanced ? 'search' : null">
      <a-form layout="horizontal">
        <div :class="advanced ? null: 'fold'">
          <a-row >
          <a-col :md="8" :sm="24" >
            <a-form-item
              label="账号"
              :labelCol="{span: 5}"
              :wrapperCol="{span: 18, offset: 1}"
            >
              <a-input placeholder="请输入" />
            </a-form-item>
          </a-col>
          <a-col :md="8" :sm="24" >
            <a-form-item
              label="状态"
              :labelCol="{span: 5}"
              :wrapperCol="{span: 18, offset: 1}"
            >
              <a-select placeholder="请选择">
                <a-select-option value="0">停用</a-select-option>
                <a-select-option value="2">启用</a-select-option>
              </a-select>
            </a-form-item>
          </a-col>
          <a-col :md="8"  :sn="24">
             <div style="margin-left: 20px; margin-top: 3px;">
              <a-button type="primary">查询</a-button>
              <a-button style="margin-left: 8px">重置</a-button>
            </div>
          </a-col>
        </a-row>

        </div>

      </a-form>
    </div>
    <div>
      <a-space class="operator">
        <a-button @click="addNew" type="primary">新建</a-button>
        <a-button >批量操作</a-button>
        <a-dropdown>
          <a-menu @click="handleMenuClick" slot="overlay">
            <a-menu-item key="delete" v-auth="`deleteUser`">删除</a-menu-item>
          </a-menu>
          <a-button>
            更多操作 <a-icon type="down" />
          </a-button>
        </a-dropdown>
      </a-space>
      <standard-table
          :rowKey="rowKey"
        :columns="columns"
        :dataSource="dataSource"
        :selectedRows.sync="selectedRows"
        @clear="onClear"
        @change="onChange"
        :pagination="{...pagination, onChange: onPageChange}"
        @selectedRowChange="onSelectChange"
      >
        <div slot="roleTag" slot-scope="{text}">
          <a-tag v-for="(role, index) in text" :key="index" :color="colorArray[index % colorArray.length]">
            {{role.role_name}}
          </a-tag>
        </div>

        <div slot="statusTag" slot-scope="{text}">
          <a-tag v-if="text === 1" color="green">启用</a-tag>
          <a-tag v-else-if="text === 0">停用</a-tag>
        </div>

        <div slot="action" slot-scope="{text, record}">
          <a style="margin-right: 8px" v-auth="`editInfo`">
            <a-icon type="editInfo"/>修改资料
          </a>
          <a @click="deleteRecord(record.id)" v-auth="`editRole`">
            <a-icon type="edit" />修改角色
          </a>
          <a @click="deleteRecord(record.id)" v-auth="`deleteUser`">
            <a-icon type="delete" />删除账号
          </a>
        </div>
      </standard-table>
    </div>
  </a-card>
</template>

<script>
import StandardTable from '@/components/table/StandardTable'
import {getUserList} from '@/services/user'
import dayjs from 'dayjs';

const columns = [
  {
    title: 'ID',
    dataIndex: 'id'
  },
  {
    title: '用户名',
    dataIndex: 'username',
  },
  {
    title: 'IP地址',
    dataIndex: 'ip',
    sorter: true,
  },
  {
    title: '角色',
    dataIndex: 'role',
    sorter: false,
    scopedSlots: { customRender: 'roleTag' }
  },
  {
    title: '状态',
    dataIndex: 'status',
    needTotal: true,
    scopedSlots: { customRender: 'statusTag' }
  },
  {
    title: '最后登陆',
    dataIndex: 'login_last',
    sorter: true,
    customRender: (text) => dayjs.unix(text).format('YYYY-MM-DD HH:mm:ss')
  },
  {
    title: '操作',
    scopedSlots: { customRender: 'action' }
  }
]

let that
export default {
  name: 'UserList',
  components: {StandardTable},
  data () {
    return {
      colorArray: ['pink', 'red', 'orange', 'green', 'cyan', 'blue', 'purple'],
      advanced: true,
      columns: columns,
      dataSource: [],
      selectedRows: [],
      rowKey: 'id',
      params: {
        username: null,
        status: null
      },
      pagination: {
        current: 1,
        pageSize: 10,
        total: 0
      }
    }
  },
  authorize: {
    editInfoRecord: 'editInfo',
    editRoleRecord: 'editRole',
    deleteUserRecord: 'deleteUser'
  },
  mounted() {
    this.getData()
    console.log(this.selectedRows)
  },
  methods: {
    onPageChange(page, pageSize) {
      this.pagination.current = page
      this.pagination.pageSize = pageSize
      this.getData()

    },
    async getData() {
      that = this
      let {data} = await getUserList({
        page: that.pagination.current,
        page_size: that.pagination.pageSize,
        ...that.params
      })

      if (data.code === 0) {
        that.dataSource = data.data.list
        that.pagination = {
          current: data.data.page,
          pageSize: data.data.page_size,
          total: data.data.count
        }
      } else {
        this.$message.error(data.msg)
      }

    },
    deleteRecord(key) {
      console.log(key)
      this.dataSource = this.dataSource.filter(item => item.id !== key)
      this.selectedRows = this.selectedRows.filter(item => item.id !== key)
    },
    toggleAdvanced () {
      this.advanced = !this.advanced
    },
    remove () {
      this.dataSource = this.dataSource.filter(item => this.selectedRows.findIndex(row => row.id === item.id) === -1)
      this.selectedRows = []
    },
    onClear() {
      this.$message.info('您清空了勾选的所有行')
    },
    onStatusTitleClick() {
      this.$message.info('你点击了状态栏表头')
    },
    onChange() {
      this.$message.info('表格状态改变了')
    },
    onSelectChange(selectedRowKeys, selectedRows) {
      console.log(selectedRowKeys, selectedRows)
    },
    addNew () {
      this.dataSource.unshift({
        key: this.dataSource.length,
        no: 'NO ' + this.dataSource.length,
        description: '这是一段描述',
        callNo: Math.floor(Math.random() * 1000),
        status: Math.floor(Math.random() * 10) % 4,
        updatedAt: '2018-07-26'
      })
    },
    handleMenuClick (e) {
      if (e.key === 'delete') {
        this.remove()
      }
    }
  }
}
</script>

<style lang="less" scoped>
  .search{
    margin-bottom: 54px;
  }
  .fold{
    width: calc(100% - 216px);
    display: inline-block
  }
  .operator{
    margin-bottom: 18px;
  }
  @media screen and (max-width: 900px) {
    .fold {
      width: 100%;
    }
  }
</style>
