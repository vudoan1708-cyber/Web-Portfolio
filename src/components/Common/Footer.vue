<template>
  <div id="footer">
    <footer>Made with <a href="https://vuejs.org/" target="_blank">Vue.js</a></footer>
  </div>
</template>

<script>
export default {
  name: 'Footer',
};
</script>

<style lang="scss" scoped>
#footer {
  position: absolute;
  bottom: 0;
  margin: 0 auto;
  padding: 10px;
  background-color: rgb(148, 53, 53);
  color: rgb(146, 146, 230);
  width: 100%;

  a {
    text-decoration: none;
    color: rgb(146, 146, 230);
  }
}
</style>
