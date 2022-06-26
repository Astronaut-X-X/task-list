<template>
  <div class="flex">
    <div class="px-6 w-full">
      <div class="flex flex-col items-end">
        <i class="el-icon-circle-plus-outline text-xl hover:opacity-60 transition-opacity" @click="insert"></i>
      </div>
      <div class="flex flex-col items-end ">
        <div v-for="(item, index)  in quoteData" v-bind:item="item" v-bind:index="index" v-bind:key="item.ID"
          class="py-2 w-full flex flex-row items-center justify-center border-b">
          <div class="flex-1 flex justify-between" @dblclick="edit(index)">
            <a v-if="index !== editing" class="tracking-wide hover:text-blue-500 transition-all"
              @click="clickPlan(item.ID)">{{item.quote}}</a>
            <div v-else class="w-3/6">
              <el-input v-model="item.quote" @keyup.enter.native="finishEdit(item)" @blur="finishEdit(item)"
                size="small" />
            </div>
            <a v-if="index !== editing" class="tracking-wide hover:text-blue-500 transition-all pre-quote-author"
              @click="clickPlan(item.ID)">{{item.author}}</a>
            <div v-else class="w-1/6">
              <el-input v-model="item.author" @keyup.enter.native="finishEdit(item)" @blur="finishEdit(item)"
                size="small" />
            </div>
          </div>
          <i class="el-icon-delete text-xl opacity-20 ml-6  hover:opacity-100 transition-opacity"
            @click="deleteQuote(item)"></i>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Quote',
  data () {
    return {
      editing: -1,
      quoteData: [
        {
          ID: 1,
          user_id: 1,
          quote: '励志名言',
          author: '作者',
        }
      ]
    }
  },

  mounted () {

  },

  methods: {
    initQuoteData () {
      this.quoteData = [
        {
          ID: 1,
          user_id: 1,
          quote: '励志名言',
          author: '作者',
        }
      ];
    },
    insert () {
      this.quoteData.push({
        ID: 1,
        user_id: 1,
        quote: '励志名言',
        author: '作者',
      });
    },
    edit (index) {
      this.editing = index;
    },
    finishEdit (item) {
      this.editing = -1;
    },
    deleteQuote (item) {
      this.initQuoteData();
    }
  }
}
</script>

<style>
.pre-quote-author::before {
  content: "—— ";
}
</style>