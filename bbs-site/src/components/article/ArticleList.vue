<template>
  <div class="article-list">
    <div
      v-for="article in articles"
      :key="article.articleId"
      class="article-item"
    >
      <div class="article-item-main">
        <div class="article-info">
          <a href="" class="article-title">{{ article.title }}丢了</a>
          <div class="article-summary">
            {{ article.summary }}
          </div>
        </div>

        <div class="article-meta">
          <div class="article-meta-left">
            <span class="article-meta-item">
<!--              <nuxt-link-->
<!--                :to="'/user/' + article.user.id"-->
<!--                class="article-author"-->
<!--              >-->
<!--                <span>{{ article.user.nickname }}</span>-->
<!--              </nuxt-link>-->
              <time
                :datetime="
                  article.createTime "
              >发布于 {{ article.createTime }}</time
              >
            </span>
          </div>

          <div class="article-meta-right">
            <div v-if="article.tags && article.tags.length > 0">
<!--              <nuxt-link-->
<!--                v-for="tag in article.tags"-->
<!--                :key="tag.tagId"-->
<!--                class="article-tag"-->
<!--                :to="'/articles/' + tag.tagId"-->
<!--              >{{ tag.tagName }}</nuxt-link-->
              {{ tag.tagName }}
              >
            </div>
          </div>
        </div>
      </div>
      <div
        v-if="article.cover"
        v-lazy-container="{ selector: 'img' }"
        class="article-item-cover"
      >
        <img :data-src="article.cover.url" />
      </div>
    </div>
  </div>
</template>

<script>

export default {
  props: {
    articles: {
      type: Array,
      default() {
        return []
      },
      required: false,
    },
  },
}
</script>

<style scoped>

.article-list {
  margin: 0 !important;

  .article-item {
    padding: 12px 12px;
    transition: background 0.5s;
    /* border-radius: 3px; */
    background: #fff;
    line-height: 24px;
    border-bottom: #8c939d 1px solid;

    &:not(:last-child) {
      margin-bottom: 10px;
    }

    display: flex;
    .article-item-main {
      width: 100%;
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      .article-info {
        .article-title {
          font-size: 18px;
          line-height: 30px;
          font-weight: 500;
          color: #3e606f;
          overflow: hidden;
          text-overflow: ellipsis;
        }

        .article-summary {
          font-size: 14px;
          color: #3e606f;
          overflow: hidden;
          display: -webkit-box;
          -webkit-box-orient: vertical;
          -webkit-line-clamp: 3;
          text-align: justify;
          padding-top: 6px;
          word-break: break-all;
          text-overflow: ellipsis;
        }
      }

      .article-meta {
        display: flex;
      // justify-content: space-between;
        align-items: center;
        font-size: 13px;
        padding-top: 6px;

        .article-meta-left {
          .article-meta-item {
            padding: 0 6px 0 0;
            color: #91aa9d;

            .article-author {
              font-weight: bold;
              padding: 0 3px;
            }
          }
        }

        .article-meta-right {
          @media screen and (max-width: 768px) {
            & {
              display: none;
            }
          }

          margin-left: 10px;

          .article-tag {
            padding: 2px 8px;
            justify-content: center;
            align-items: center;
            border-radius: 12.5px;
            margin-right: 10px;
            background: #0ff123;
            border: 1px solid #8c939d;
            color: #3498db;
            font-size: 12px;

            &:hover {
              color: #dbdbdb;
              background: #ffffff;
            }
          }
        }
      }
    }

    .article-item-cover {
      display: flex;
      margin-left: 6px;
      img {
        min-width: 140px;
        min-height: 90px;
        width: 140px;
        height: 90px;
        object-fit: cover;

        @media screen and (max-width: 768px) {
          & {
            min-width: 110px;
            min-height: 80px;
            width: 110px;
            height: 80px;
          }
        }
      }
    }
  }
}

/deep/ .el-tabs__nav-scroll{
  padding-bottom: 10px;
}
</style>
