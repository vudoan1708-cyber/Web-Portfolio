<template>
  <div id="further_detail">
    <!-- Close Button -->
    <div id="close_btn" @click="closeProjectDetail">X</div>
    <div id="details">
      <!-- Title -->
      <section id="project_title" v-if="detail.title !== ''">
        <h1>{{ detail.title }}</h1>
      </section>
      <!-- Description -->
      <section v-if="detail.desc !== ''">
        <p>{{ detail.desc }}</p>
      </section>

      <section id="links_group" >
        <!-- Videos -->
        <section class="detail" v-if="detail.videos !== []">
          <div class="videos_container" v-for="(video, videoKey) in detail.videos" :key="videoKey">
            <h3>{{ video.name }}</h3>
            <div class="img_container">
              <a :href="video.link" target="_blank" alt="demo_video">
                <img src="@/assets/logos/youtube.png" />
              </a>
            </div>
          </div>
        </section>
        <!-- Project Log -->
        <section class="detail" v-if="detail.project_log !== ''">
          <a :href="detail.project_log" target="_blank" alt="report_link">
            <p>Project Log</p>
          </a>
        </section>
        <!-- Reports -->
        <section class="detail" v-if="detail.reports !== []">
          <div class="report_container" v-for="(report, reportKey) in detail.reports"
            :key="reportKey">
            <a :href="report.link" target="_blank" alt="report_link">
              <p>{{ report.name }}</p>
            </a>
          </div>
        </section>
        <!-- Github -->
        <section class="detail" v-if="detail.github !== ''">
          <div class="img_container">
            <a :href="detail.github" target="_blank" alt="github_url">
              <img src="@/assets/logos/github.png" />
            </a>
          </div>
        </section>
        <!-- Live URL -->
        <section class="detail" v-if="detail.live_url !== ''">
          <a :href="detail.live_url" target="_blank" alt="github_url">
            <p>Live URL</p>
          </a>
        </section>
      </section>
    </div>
  </div>
</template>

<script>
import { reactive } from 'vue';

export default {
  name: 'FurtherDetail',
  props: {
    emitter: {
      type: Object,
    },
    furtherDetail: {
      type: Object,
    },
  },
  setup(props) {
    const detail = reactive({
      title: props.furtherDetail.title,
      desc: props.furtherDetail.description,
      videos: props.furtherDetail.extra_links.videos,
      presentations: props.furtherDetail.extra_links.presentations,
      project_log: props.furtherDetail.extra_links.project_log,
      reports: props.furtherDetail.extra_links.reports,
      github: props.furtherDetail.github,
      live_url: props.furtherDetail.live_url,
      organisation: props.furtherDetail.organisation,
    });

    function closeProjectDetail() {
      props.emitter.emit('project_card', null);
    }

    return {
      closeProjectDetail,
      detail,
    };
  },
};
</script>

<style lang="scss" scoped>
#further_detail {
  display: block;
  position: relative;
  top: 15%;
  width: 80%;
  padding: 20px;
  text-align: right;

  #close_btn {
    display: inline-block;
    position: relative;
    top: 15%;
    padding: 20px;
    border-radius: 50%;
    border: 4px dashed black;
    cursor: pointer;
    transition: .2s all;

    &:hover {
      border: 4px dashed white;
      color: white;
    }
  }

  #details {
    position: relative;
    margin: 10px;
    padding: 20px;
    color: rgb(0, 0, 0);
    border-radius: 5px;
    background-color: rgba(87, 87, 87, 0.233);
    text-align: left;

    section {
      margin: 20px;
      text-align: left;
    }

    #project_title {
      text-align: center;
    }
    #links_group {
      position: relative;
      margin: 10px;

      .detail {
        display: inline-block;
      }
    }

    .img_container {
      width: 12%;
      margin: 40px;
      cursor: pointer;

      img {
        width: 100%;
        transition: 0.2s filter;
      }

      &:hover {
        img {
          filter: invert(100%);
        }
      }
    }
  }
}
</style>
