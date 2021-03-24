<template>
  <div id="blogs">
    <h1>BLOG</h1> <br />
    <div class="blog_content" v-for="(content, contentKey) in Blog" :key="contentKey">
      <h3 class="headings">{{ content.title }}</h3>
      <div class="public_date">
        <i>{{ content.date_of_publication }}</i>
      </div>
      <section class="contents">
        <p>{{ content.content }}</p>
      </section>
    </div>
  </div>
</template>

<script>
// JSON
import Blog from '@/components/JSON/blog.json';

export default {
  name: 'Blogs',
  setup() {
    return {
      Blog,
    };
  },
};
</script>

<style lang="scss" scoped>
#blogs {
  display: block;
  position: relative;
  top: 15%;
  width: 80%;
  padding: 20px;
  color: snow;

  .blog_content {
    background-color: black;
    color: rgb(221, 221, 221);
    border-radius: 20px;

    .headings {
      background-color: rgb(145, 100, 100);
      text-align: center;
      padding: 10px;
    }
    .public_date {
      margin: 10px;
      text-align: right;
    }

    .contents {
      margin: 10px;
      padding: 20px;
      text-align: left;
    }
  }
}
</style>
