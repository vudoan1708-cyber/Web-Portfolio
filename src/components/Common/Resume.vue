<template>
  <div id="resume">
    <h1>RESUME</h1> <br />
    <div id="embeded">
      <p>Open the CV file in another
        <a href="https://onedrive.live.com/?authkey=%21AHj%2DzKHwgq04VFU&cid=9D99B78A1DA5DAD8&id=9D99B78A1DA5DAD8%21329&parId=9D99B78A1DA5DAD8%21323&o=OneUp" target="_blank">window</a>
      </p> <br />
      <iframe :src="CV" width="100%" height="750px">
      </iframe>
    </div>
  </div>
</template>

<script>
// PDF
import CV from '@/assets/pdfs/CV.pdf';

export default {
  name: 'Resume',
  setup() {
    return {
      CV,
    };
  },
};
</script>

<style lang="scss" scoped>
#resume {
  display: block;
  position: relative;
  top: 15%;
  width: 80%;
  padding: 20px;
  color: snow;

  #embeded {
    p {
      text-align: right;
    }
  }
}
</style>
